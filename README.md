# dockerfile

---

`dockerfile` is a CLI utility for quickly structured extraction of information contained in dockerfile

## Example Usage

```bash
# Find all the BaseImages used in the Dockerfile
dockerfile inspect --path ./Dockerfile -f "{{.From}}"
# Find the BaseImage used in all Dockerfiles in the current directory
dockerfile inspect --path . -f "{{.From}}"
```
