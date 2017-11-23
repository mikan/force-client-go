package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/mikan/force-client-go/config"
	"github.com/mikan/force-client-go/force"
)

var f = flag.String("f", "config.json", "load specified credential & parameter file")
var q = flag.String("q", "", "run SOQL (e.g. \"SELECT Id FROM Contact\")")
var c = flag.String("c", "", "create (also specify \"j\" option)")
var r = flag.String("r", "", "read (also specify \"i\" option)")
var u = flag.String("u", "", "update (also specify \"i\" and \"j\" option)")
var d = flag.String("d", "", "delete (also specify \"i\" option)")
var j = flag.String("j", "", "JSON data for create/update")
var i = flag.String("i", "", "SObject ID for read/update/delete (e.g. 0037F00000Hc2GyQAJ)")

// op defines object operators
type op int

const (
	queryOp op = iota
	createOp
	readOp
	updateOp
	deleteOp
)

type request struct {
	op   op
	name string
	id   string
	data string
}

var logger = log.New(os.Stdout, "", log.LstdFlags)

// main demonstrates CRUD operations for each standard SObjects.
func main() {
	flag.Parse()
	p := load(*f)
	exec(p.Cred(), client(p.Prod, p.Ver), detect(*c, *r, *u, *d, *q, *i, *j))
}

func load(file string) *config.Params {
	if len(file) == 0 {
		logger.Println("Please specify credential file with \"-f\" option (\"-h\" for help).")
		return nil
	}
	cred, err := config.Load(file)
	if err != nil {
		logger.Printf("Failed to load credential(%s): %v", file, err)
		return nil
	}
	logger.Printf("%s loaded.", file)
	return cred
}

func detect(create, read, update, del, query, id, data string) *request {
	if len(query) > 0 {
		return &request{queryOp, "", "", query}
	}
	if len(create) > 0 {
		return &request{createOp, create, "", data}
	}
	if len(read) > 0 {
		return &request{readOp, read, id, ""}
	}
	if len(update) > 0 {
		return &request{updateOp, update, id, data}
	}
	if len(del) > 0 {
		return &request{deleteOp, del, id, ""}
	}
	logger.Println("unknown operation (\"-h\" for help).")
	return nil
}

func client(production bool, version string) *force.Client {
	env := force.Sandbox
	if production {
		env = force.Production
	}
	client, err := force.NewClient(env, version, logger)
	if err != nil {
		logger.Fatalf("Failed to create new client: %v", err)
	}
	return client
}

func exec(cred *force.Credential, client *force.Client, request *request) {
	if cred == nil || client == nil || request == nil {
		logger.Fatal("precondition check failed.")
	}
	ctx := context.Background()
	err := client.Login(ctx, cred)
	if err != nil {
		logger.Fatalf("Failed to login: %v", err)
	}
	switch request.op {
	case queryOp:
		var res interface{}
		next, err := client.Query(ctx, request.data, &res)
		if err != nil {
			client.Logger.Printf("failed to execute query: %v", err)
			logger.Fatal(err)
		}
		prettyPrint("RESULT: ", res)
		if len(next) > 0 {
			logger.Printf("Next resource found: %s", next)
		}
	case createOp:
		id, err := client.Create(ctx, request.name, request.data)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Printf("Created %s\n", id)
	case readOp:
		var res interface{}
		if err := client.Read(ctx, request.name, request.id, &res); err != nil {
			logger.Fatal(err)
		}
		prettyPrint("RESULT: ", res)
	case updateOp:
		err := client.Update(ctx, request.name, request.id, request.data)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Printf("Update successfly: %s\n", request.id)
	case deleteOp:
		err := client.Delete(ctx, request.name, request.id)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Printf("Delete successfly: %s\n", request.id)
	default:
		// Oops!
	}
}

func prettyPrint(prefix string, v interface{}) {
	indent, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		logger.Printf("%s\n%s\n", prefix, string(indent)) // pretty print
	} else {
		logger.Printf("%s%v\n", prefix, v) // raw print
	}
}
