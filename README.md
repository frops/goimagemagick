# golang_imagemagick


### build
```bash
GOARCH=amd64 CGO_ENABLED=1 CGO_CFLAGS_ALLOW='-Xpreprocessor' go build -o build/app 
```

### run
```bash
build/app
```
_-or-_
```bash
go run main.go
```