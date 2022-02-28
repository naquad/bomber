.PHONY: clean

all: bomber

bomber: main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -s'
	upx -9 $@

clean:
	rm -f bomber
