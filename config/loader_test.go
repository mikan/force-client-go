package config

import "testing"

func TestLoadWithSampleFile(t *testing.T) {
	params, err := Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(params.Instance) == 0 {
		t.Fatal("Instance not found")
	}
	if !params.Prod {
		t.Fatal("Prod is false or missing")
	}
	if len(params.Ver) == 0 {
		t.Fatal("Ver not found")
	}
	if len(params.ClientID) == 0 {
		t.Fatal("ClientID not found")
	}
	if len(params.ClientSecret) == 0 {
		t.Fatal("ClientSecret not found")
	}
	if len(params.Username) == 0 {
		t.Fatal("Username not found")
	}
	if len(params.Password) == 0 {
		t.Fatal("Password not found")
	}
	if len(params.APIToken) == 0 {
		t.Fatal("APIToken not found")
	}
	cred := params.Cred()
	if len(cred.ClientID) == 0 {
		t.Fatal("ClientID not found")
	}
	if len(cred.ClientSecret) == 0 {
		t.Fatal("ClientSecret not found")
	}
	if len(cred.Username) == 0 {
		t.Fatal("Username not found")
	}
	if len(cred.Password) == 0 {
		t.Fatal("Password not found")
	}
	if len(cred.APIToken) == 0 {
		t.Fatal("APIToken not found")
	}
}

func TestLoadWithNoSuchFile(t *testing.T) {
	_, err := Load("not-found.json")
	if err == nil {
		t.Fatal("missing expected error")
	}
}

func TestLoadWithIllegalFile(t *testing.T) {
	_, err := Load("../README.md")
	if err == nil {
		t.Fatal("missing expected error")
	}
}
