## Praktikum 2 

Only Google Chrome browser is supported! There are 4 banks, 
one of them is not reporting to the stock exchange in order to save its bank reserve.
I will later request money from them or send him Money with RPC. 

## To test

use the command 'docker compose build', then the command: 'docker compose up'. 
To access a bank in the browser, enter `localhost:6543` to 'localhost:6546' 
in Google Chrome.

If you enter `localhost:6543`, you will see in Docker which bank received the HTTP request.

## for Praktikum 3 

gRPC Create a .proto file than Compile it with protoc

for example: at first make sure that protoc installed `protoc --version`

`protoc --go_out=chat --go_opt=paths=source_relative --go-grpc_out=chat --go-grpc_opt=paths=source_relative chat.proto`

## for Praktikum 4 MQTT

Note: make sure that the line `go StartMQTT() ` in bank/main.go is not commented out.

How the content of the received message is transferred? -> 'responseChannel'
How to create a channel? -> Type 'chan'
