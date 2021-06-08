package main

import (
	"RegistrationToActivity/api/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

const(
	port = ":50051"
	adminLogin = "login"
	adminPassword = "password"
)

type server struct {
	auth.UnimplementedAuthServer
}

func main() {
	upAuthServer()
}

func upAuthServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	auth.RegisterAuthServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}