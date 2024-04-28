package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	// "github.com/mmcferren/go-micro/internal/implementation/auth"
	pb "github.com/mmcferren/go-micro/proto"
	"google.golang.org/grpc"
)

const (
	dbDriver = "mysql"
	dbUser = "root"
	dbPassword = "Admin123"
	dbName = "money_movement"
)

var db * sql.DB

func main() {
	var err error

	// Open a database connection
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)
	
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Printf("error closing db: %s", err)
		}
	}()

	// Ping the database to ensure connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// gRPC server setup
	// grpcServer := grpc.NewServer()
	// pb.RegisterMoneyMovementServiceServer(grpcServer,)

	// // Listen and serve
	// listener, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen on port 9000: %v", err)
	// }

	// log.Printf("server listening at %v", listener.Addr())	
	// if err := grpcServer.Serve(listener); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}