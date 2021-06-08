package main

import (
	"RegistrationToActivity/api/auth"
	"RegistrationToActivity/api/record"
	"log"
	"os"
)

const (
	recordPort = ":50053"
	authPort = ":50051"
)

var recordClient record.RecordClient
var authClient auth.AuthClient

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}


