package main

import (
	"AuthService/db"
	pb "AuthService/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	db *db.DbUsers
}

func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	token, err := s.db.AddUser(req.Username, req.Password)
	if err != nil {
		log.Println(err)
		return &pb.SignUpResponse{Status: false, Token: "Couldn't create the account"}, nil
	}
	return &pb.SignUpResponse{Status: true, Token: token}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	token, err := s.db.CheckUser(req.Username, req.Password)
	if err != nil {
		log.Println("SignIn error:", err)
		return &pb.SignInResponse{Status: false, Token: ""}, nil
	}

	return &pb.SignInResponse{Status: true, Token: token}, nil
}

func main() {
	usersDB, err := db.Conn()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer usersDB.DB.Close()

	ls, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthService{
		db: usersDB,
	})

	log.Println("gRPC Auth server listening on :50051")
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
