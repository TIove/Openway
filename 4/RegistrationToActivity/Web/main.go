package main

import (
	"RegistrationToActivity/api/auth"
	"RegistrationToActivity/api/record"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	port := flag.String(
		"address",
		":8001",
		"Port of this web service")
	flag.Parse()

	app := Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	app.InfoLog.Printf("Server started on %s", "localhost" + *port)

	go upRecordClient()
	go upAuthClient()

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func upRecordClient() {
	conn, err := grpc.Dial("localhost" + recordPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	recordClient = record.NewRecordClient(conn)
}

func upAuthClient() {
	conn, err := grpc.Dial("localhost" + authPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	authClient = auth.NewAuthClient(conn)
}
