# Go Echo

A websocket server that does nothing but reply back with the message you send it.

## Running

To install dependencies, first install and run `dep`

`dep ensure`

Once setup, launch with:

`go run main.go`

To change the port:

`PORT=8888 go run main.go`

## Options

- `PORT`: set the port of the server (default: 8080

## Using

Once the server is running, you can connect to it locally with `ws://localhost:8080`
