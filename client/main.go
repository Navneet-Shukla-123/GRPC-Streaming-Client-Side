package main

import (
	"client-streaming/file"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Println("Error in connecting to grpc server ", err)
		return
	}

	client := file.NewMyStreamingServiceClient(conn)

	stream, err := client.SendData(context.Background())
	if err != nil {
		log.Println("Error in creating the client ", err)
		return
	}

	for i := 1; i <= 10; i++ {
		req := &file.RequestBody{
			X: int32(i),
		}
		if err := stream.Send(req); err != nil {
			log.Println("Error sending request:", err)
			return
		}
	}
	// Close the stream
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Error receiving response: ", err)
		return
	}

	log.Println("Received response is ", resp)

}
