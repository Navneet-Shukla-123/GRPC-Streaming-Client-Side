package main

import (
	"client-streaming/file"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type clientStreaming struct {
	file.UnimplementedMyStreamingServiceServer
}

// SendData function will receive the stream of data from request and sum it
func (s clientStreaming) SendData(stream file.MyStreamingService_SendDataServer) error {
	var sum int32 = 0
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				resp := &file.ResponseBody{
					X: sum,
				}
				log.Println("Sum of total value is ", sum)
				return stream.SendAndClose(resp)
			}
			return err

		}

		log.Println("Received value is ", req)
		sum += req.GetX()

	}
	
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Println("Error in starting the TCP server ", err)
		return
	}

	grpcServer := grpc.NewServer()

	clientStream := &clientStreaming{}

	file.RegisterMyStreamingServiceServer(grpcServer, clientStream)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Println("Error in starting the Server")
		return
	}

}
