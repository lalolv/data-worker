version=v0.1
LFILE=./release/$version/data-worker

echo "Start build amd64 version for Linux ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -v -o $LFILE-linux main.go

echo "Start build amd64 version for macOS ..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -v -o $LFILE-mac main.go

echo "Start build amd64 version for Windows ..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -v -o $LFILE-win.exe main.go
