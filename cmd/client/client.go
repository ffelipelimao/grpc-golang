package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ffelipelimai/grpc-test/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "123",
		Name:  "Limao",
		Email: "limao@limao.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "123",
		Name:  "Limao",
		Email: "limao@limao.com",
	}
	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "123",
			Name:  "Limao",
			Email: "limao@limao.com",
		},
		{
			Id:    "1234",
			Name:  "Limao2",
			Email: "limao2@limao.com",
		},
		{
			Id:    "12345",
			Name:  "Limao3",
			Email: "limao3@limao.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Could not create the req: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not receive the response: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Could not create the req: %v", err)
	}
	reqs := []*pb.User{
		{
			Id:    "123",
			Name:  "Limao",
			Email: "limao@limao.com",
		},
		{
			Id:    "1234",
			Name:  "Limao2",
			Email: "limao2@limao.com",
		},
		{
			Id:    "12345",
			Name:  "Limao3",
			Email: "limao3@limao.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Could not receive the message: %v", err)
				break
			}
			fmt.Println("Status: ", res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
