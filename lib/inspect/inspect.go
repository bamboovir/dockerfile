package inspect

import (
	"fmt"
	"os"
	"strings"

	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	_ "github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// File defination
func File(fs afero.Fs, path string, formater string) ([]string, error) {
	errmsg := fmt.Sprintf("inspect path: [%s] with formater: [%s] failed", path, formater)
	f, err := fs.Open(path)

	if err != nil {
		log.Errorln(errmsg)
		return []string{}, errors.Wrap(err, errmsg)
	}

	defer f.Close()

	ast, err := parser.Parse(f)

	log.Infof("Dockerfile AST: [%s]\n", ast.AST.Dump())

	if err != nil {
		log.Errorln(errmsg)
		ast.PrintWarnings(log.StandardLogger().Out)
		return []string{}, errors.Wrap(err, errmsg)
	}
	
	stages, metaArgs, err := instructions.Parse(ast.AST)

	log.Infof("Stages : [%+v]", stages)
	log.Infof("MetaArgs : [%+v]", metaArgs)

	if err != nil {
		log.Errorln(errmsg)
		return []string{}, errors.Wrap(err, errmsg)
	}

	rst := make([]string, 0)

	for _, s := range stages {
		rst = append(rst, s.BaseName)
	}

	return rst, nil
}

// Folder defination
func Folder(fs afero.Fs, path string, formater string) ([]string, error) {
	errmsg := fmt.Sprintf("inspect path: [%s] with formater: [%s] failed", path, formater)

	rst := make([]string, 0)

	err := afero.Walk(fs, path, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if info.Mode().IsRegular() && strings.HasSuffix(path, "Dockerfile") {
			log.Infof("process [%s]", path)
			s, err := File(fs, path, formater)

			if err != nil {
				return err
			}

			rst = append(rst, s...)
		}

		return nil
	})

	if err != nil {
		return []string{}, errors.Wrap(err, errmsg)
	}

	return rst, nil
}

// Inspect defination
func Inspect(fs afero.Fs, path string, formater string) ([]string, error) {
	errmsg := fmt.Sprintf("inspect path: [%s] with formater: [%s] failed", path, formater)
	isDir, err := afero.IsDir(fs, path)
	if err != nil {
		return []string{}, errors.Wrap(err, errmsg)
	}

	if isDir {
		return Folder(fs, path, formater)
	}

	return File(fs, path, formater)
}
