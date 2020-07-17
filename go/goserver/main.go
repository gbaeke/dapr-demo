package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"

	commonv1pb "github.com/dapr/go-sdk/dapr/proto/common/v1"
	"google.golang.org/grpc"
)

// server is our user app
type server struct {
}

var (
	logger = log.New(os.Stdout, "", 0)
)

func main() {
	// create listener
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterAppCallbackServer(s, &server{})

	// and start...
	if err = s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

// HelloFromGo is a simple demo method to invoke
func (s *server) HelloFromGo(messageID float64) string {
	return fmt.Sprintf("Message ID = %v", messageID)

}

// This method gets invoked when a remote service has called the app through Dapr
// The payload carries a Method to identify the method, a set of metadata properties and an optional payload
func (s *server) OnInvoke(ctx context.Context, in *commonv1pb.InvokeRequest) (*commonv1pb.InvokeResponse, error) {
	var response string

	switch in.Method {
	case "HelloFromGo":

		// in.Data.Value contains the data sent by the caller
		data := fmt.Sprintf("%s", in.Data.Value)

		// convert data to JSON
		var helloMessage map[string]interface{}
		if err := json.Unmarshal([]byte(data), &helloMessage); err != nil {
			logger.Println("Error converting JSON ", err)
			response = ""
		}

		// call HelloFromGo passing the message Id
		logger.Println(helloMessage)
		response = s.HelloFromGo(helloMessage["messageId"].(float64))
		logger.Println(response)
	}

	return &commonv1pb.InvokeResponse{
		ContentType: "text/plain; charset=UTF-8",
		Data:        &any.Any{Value: []byte(response)},
	}, nil
}

// Dapr will call this method to get the list of bindings the app will get invoked by. In this example, we are telling Dapr
// To invoke our app with a binding named storage
func (s *server) ListInputBindings(ctx context.Context, in *empty.Empty) (*pb.ListInputBindingsResponse, error) {
	return nil, nil
}

// Dapr will call this method to get the list of topics the app wants to subscribe to. In this example, we are telling Dapr
// To subscribe to a topic named TopicA
func (s *server) ListTopicSubscriptions(ctx context.Context, in *empty.Empty) (*pb.ListTopicSubscriptionsResponse, error) {
	return nil, nil
}

// This method gets invoked every time a new event is fired from a registerd binding. The message carries the binding name, a payload and optional metadata
func (s *server) OnBindingEvent(ctx context.Context, in *pb.BindingEventRequest) (*pb.BindingEventResponse, error) {
	return &pb.BindingEventResponse{}, nil
}

// This method is fired whenever a message has been published to a topic that has been subscribed. Dapr sends published messages in a CloudEvents 0.3 envelope.
func (s *server) OnTopicEvent(ctx context.Context, in *pb.TopicEventRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
