package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"os"

	"github.com/mikan/force-client-go/config"
	"github.com/mikan/force-client-go/force"
	"github.com/mikan/force-client-go/sobject"
)

var f = flag.String("f", "config-dev.json", "load specified credential file")

func main() {
	flag.Parse()
	params, _ := config.Load(*f)

	// Login
	client, _ := force.NewClient(params.Instance, params.Env(), params.Ver, log.New(os.Stdout, "", log.LstdFlags))
	if err := client.Login(context.Background(), params.Cred()); err != nil {
		log.Fatal(err)
	}

	// Query
	fmt.Println("##### Query")
	var set sobject.ContactSet
	client.Query(context.Background(), "SELECT Id,FirstName,LastName FROM Contact", &set)
	for _, c := range set.Records {
		fmt.Printf("%s: %s %s\n", c.Id, c.FirstName, c.LastName)
	}

	// Create
	fmt.Println("##### Create")
	contact := sobject.Contact{FirstName: "Test", LastName: "User"}
	id, err := client.Create(context.Background(), sobject.ContactObjectName, &contact)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created the Contact with id %s\n", id)

	// Read
	fmt.Println("##### Read")
	var readResult sobject.Contact
	err = client.Read(context.Background(), sobject.ContactObjectName, id, &readResult)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", readResult)

	// Update
	fmt.Println("##### Update")
	update := sobject.Contact{FirstName: "Test2"}
	err = client.Update(context.Background(), sobject.ContactObjectName, id, &update)
	if err != nil {
		log.Fatal(err)
	}

	// Delete
	fmt.Println("##### Delete")
	err = client.Delete(context.Background(), sobject.ContactObjectName, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted the Contact with id %s\n", id)
}
