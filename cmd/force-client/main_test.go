package main

import (
	"flag"
	"os"
	"testing"

	"github.com/mikan/force-client-go/force"
)

func TestMain(m *testing.M) {
	flag.Parse()
	code := m.Run() // begin testing
	os.Exit(code)   // end testing
}

func TestCredActual(t *testing.T) {
	cred := load("../../config.json")
	if cred == nil {
		t.Fatal("unexpected nil result")
	}
}

func TestCredEmpty(t *testing.T) {
	cred := load("")
	if cred != nil {
		t.Fatal("unexpected actual result")
	}
}

func TestCredIllegal(t *testing.T) {
	cred := load("../../README.md")
	if cred != nil {
		t.Fatal("unexpected actual result")
	}
}

func TestDetectForCreate(t *testing.T) {
	req := detect("Contact", "", "", "", "", "", "{}")
	if req.op != createOp {
		t.Fatal("mssmatched op")
	}
	if len(req.data) == 0 {
		t.Fatal("no data captured")
	}
}

func TestDetectForRead(t *testing.T) {
	req := detect("", "Contact", "", "", "", "000000000000000000", "")
	if req.op != readOp {
		t.Fatal("mssmatched op")
	}
	if len(req.id) == 0 {
		t.Fatal("no id captured")
	}
}

func TestDetectForUpdate(t *testing.T) {
	req := detect("", "", "Contact", "", "", "000000000000000000", "{}")
	if req.op != updateOp {
		t.Fatal("mssmatched op")
	}
	if len(req.id) == 0 {
		t.Fatal("no id captured")
	}
	if len(req.data) == 0 {
		t.Fatal("no data captured")
	}
}

func TestDetectForDelete(t *testing.T) {
	req := detect("", "", "", "Contact", "", "000000000000000000", "")
	if req.op != deleteOp {
		t.Fatal("mssmatched op")
	}
	if len(req.id) == 0 {
		t.Fatal("no id captured")
	}
}

func TestDetectForQuery(t *testing.T) {
	req := detect("", "", "", "", "SELECT Id FROM Contact", "", "")
	if req.op != queryOp {
		t.Fatal("mssmatched op")
	}
}

func TestDetectForNothing(t *testing.T) {
	req := detect("", "", "", "", "", "", "")
	if req != nil {
		t.Fatalf("expected nil, actual %v", req)
	}
}

func TestClientForProd(t *testing.T) {
	prod := client("localhost", true, "41.0")
	if prod.Instance.Host != "localhost" {
		t.Fatalf("expected %s, actual %s", "localhost", prod.Instance.Host)
	}
	if prod.Env != force.Production {
		t.Fatalf("expected %d, actual%d", force.Production, prod.Env)
	}
}

func TestClientForSand(t *testing.T) {
	sand := client("localhost", false, "20.0")
	if sand.Instance.Host != "localhost" {
		t.Fatalf("expected %s, actual %s", "localhost", sand.Instance.Host)
	}
	if sand.Env != force.Sandbox {
		t.Fatalf("expected %d, actual%d", force.Sandbox, sand.Env)
	}
}
