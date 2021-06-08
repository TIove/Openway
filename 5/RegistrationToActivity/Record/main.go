package main

import (
	"RegistrationToActivity/Record/dataBase"
	"RegistrationToActivity/api/auth"
	"RegistrationToActivity/api/record"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	record.UnimplementedRecordServer
}

func main() {
	dbConnection := flag.String(
		"dbConnection",
		"root:root@/?parseTime=true",
		"Database connection string")
	flag.Parse()

	db, err := openDb("mysql", *dbConnection)

	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	defer db.Close()

	app = Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Db:       &dataBase.RecordModel{Db: db},
	}

	go upAuthClient()
	upServer()
}

func upAuthClient() {
	conn, err := grpc.Dial("localhost" + authPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	authClient = auth.NewAuthClient(conn)
}

func upServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	record.RegisterRecordServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func openDb(driverName string, dbCredentials string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dbCredentials)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	createDb(db, "Records")

	return db, err
}

func createDb(db *sql.DB, dbName string) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}
