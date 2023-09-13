

build :
	GOOS=linux GOARCH=arm go build -ldflags="-s -w" .
