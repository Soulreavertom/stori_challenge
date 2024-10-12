build:
	GOOS=linux GOARCH=arm64  go build -o bootstrap ./cmd/lambda/main.go

package_lambda: build
	zip lambda-handler.zip bootstrap