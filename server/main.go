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
	err := s.db.AddUser(req.Username, req.Email, req.Password)
	if err != nil {
		log.Println(err)
		return &pb.SignUpResponse{Status: false, Token: ""}, nil
	}
	return &pb.SignUpResponse{Status: true, Token: "asdasdasd"}, nil
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
