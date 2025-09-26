package grpc

import (
	pb "AuthService/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func SignUpGRPC(email, username string, password []byte) (bool, string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return false, "Couldn't connect to gRPC server"
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SignUp(ctx, &pb.SignUpRequest{
		Email:    email,
		Username: username,
		Password: string(password),
	})
	if err != nil {
		log.Println(err)
		return false, "Couldn't request 'SignUp' method"
	}

	return resp.Status, resp.Token
}

func SignInGRPC(username, password string) (bool, string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return false, "Couldn't connecto to gRPC server"
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.SignIn(ctx, &pb.SignInRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println(err)
		return false, "Couldn't request 'SignIn' method"
	}

	return resp.Status, resp.Token
}
