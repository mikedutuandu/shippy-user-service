package main

import (
	"fmt"
	"log"

	pb "github.com/mikedutuandu/shippy-user-service/proto/auth"
	"github.com/micro/go-micro"

)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}


	srv := micro.NewService(
		micro.Name("shippy.auth"),
	)

	srv.Init()



	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
