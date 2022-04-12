#GODEBUG='gctrace=1' CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -ldflags "-s -w" ebox.go
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go
