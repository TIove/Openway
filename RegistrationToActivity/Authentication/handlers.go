package main

import (
	"RegistrationToActivity/api/auth"
	"context"
	"errors"
)

func (s *server) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	if in.Login == adminLogin && in.Password == adminPassword {
		return &auth.LoginResponse{Token: "tokensampleforadmin"}, nil //TODO add JWT
	} else {
		return nil, nil
	}
}

func (s *server) IsAdmin(ctx context.Context, in *auth.IsAdminRequest) (*auth.IsAdminResponse, error) {
	if in.Token == "tokensampleforadmin" {
		return &auth.IsAdminResponse{IsAdmin: true}, nil
	} else {
		return &auth.IsAdminResponse{IsAdmin: false}, errors.New("forbidden")
	}
}
