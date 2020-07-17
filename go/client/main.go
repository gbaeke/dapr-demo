package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	dapr "github.com/dapr/go-sdk/client"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

func main() {
	// just for this demo
	ctx := context.Background()
	data, _ := json.Marshal(map[string]float64{"messageId": 1000})

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		logger.Panic(err)
	}
	defer client.Close()

	// invoke a method called EchoMethod on another dapr enabled service
	resp, err := client.InvokeServiceWithContent(ctx, "goserver", "HelloFromGo",
		"application/json; charset=UTF-8", data)
	if err != nil {
		logger.Panic(err)
	}
	logger.Printf("service method invoked, response: %s", string(resp))

}
