package main

import (
	"RegistrationToActivity/Record/dataBase"
	"RegistrationToActivity/api/auth"
	"log"
	"os"
)

var authClient auth.AuthClient
var app Application

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

const(
	authPort = ":50051"
	port = ":50053"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Db       *dataBase.RecordModel
}
