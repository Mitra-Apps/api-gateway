# API-GATEWAY

## Configure visual studio code to use WSL
https://code.visualstudio.com/docs/remote/wsl

## How to run
run the apps using command : 
go mod tidy
go mod vendor
go run main.go

## Generate pb file from proto file
### Install buf
https://buf.build/docs/installation
If failed, run : brew install buf

### generate protobuf
Run : buf generate

## Swagger
To generate api documentation files : swag init -g main.go --output docs/
After running the app, the documentation can be seen : http://localhost:{port}/swagger/index.html