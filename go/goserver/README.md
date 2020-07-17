# Go Server

Start the Go server with:

dapr run --app-id goserver --protocol grpc --app-port 5001 go run main.go

# Go client

Start the Go client with:

dapr run --app-id goclient go run main.go

# Python client

Start the Python client with:

dapr run --app-id pyclient python3 app.py