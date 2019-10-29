Implementation
1. How to generate the go pb file form the .proto
~/go/src/grpc-web/server$ protoc calculator.proto --go_out=plugins=grpc:../server/calculatorpb/
2. Create the go.mod
~/go/src/grpc-web/server$ go mod init github.com/kaysush/grpc-calculator
3. Compile the calculator.proto file for JavaScript
~/go/src/grpc-web/protos$ protoc calculator.proto --js_out=import_style=commonjs,binary:../client --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../client
4. Package JavaScript code into a minified file
npx webpack client.js

Running

1. Start the Webserver: 
    ~/go/src/grpc-web/client$ python3 -m http.server 8081 &

2. Start the gRPC server: 
    ~/go/src/grpc-web/server$ go run server.go

3. Start the envoy proxy: [TBD]
    creating the docker image: 
    ~/go/src/grpc-web/envoy$ docker build -t calculator/envoy -f ./envoy.Dockerfile .