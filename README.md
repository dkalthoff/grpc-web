# gRPC-Web
This repo contains the files needed to make gRPC requests from the browser.

## Implementation
1. How to generate the go pb files from the .proto

`~/go/src/grpc-web/server$ protoc calculator.proto --go_out=plugins=grpc:../server/calculatorpb/`
`~/go/src/grpc-web/server$ protoc customers.proto --go_out=plugins=grpc:../server/customerspb/`

2. Create the go.mod

`~/go/src/grpc-web/server$ go mod init github.com/kaysush/grpc-calculator`

3. Compile the calculator.proto file for JavaScript

`~/go/src/grpc-web/protos$ protoc calculator.proto --js_out=import_style=commonjs,binary:../client --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../client`

4. Install npm and package JavaScript code into a minified file

`~/go/src/grpc-web/client$ npm install
~/go/src/grpc-web/client$ npx webpack client.js`

## Running

1. Start the Webserver:

`~/go/src/grpc-web/client$ python3 -m http.server 8081 &`

2. Start the gRPC server:

`~/go/src/grpc-web/server$ go run server.go`

3. Start the envoy proxy:
    - Creating the docker image:
    
    `docker build -t grpcproxy/envoy -f ./envoy.grpcproxy.Dockerfile .`
    
    - Run the docker container
    
    `docker run -d --network=host grpcproxy/envoy:latest`
    
    - View all the docker containers
    
    `docker ps`
