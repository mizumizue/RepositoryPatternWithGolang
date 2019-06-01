package main

import (
	"context"
	"fmt"
	"github.com/trewanek/repositoryPattern/repository"
	"log"
)

func main() {
	var rep repository.UserRepository // Interface

	ctx := context.Background()
	//rep = repository.NewUserDatastoreRepository(ctx) // Impl by Datastore
	rep = repository.NewUserInMemoryRepository() // Impl by InMemory

	// Use
	ul, err := rep.List(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, u := range ul {
		fmt.Printf("%#v\n", u)
	}
}
