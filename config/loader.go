package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/mikan/force-client-go/force"
)

type Params struct {
	force.Credential
	Instance string `json:"instance"`
	Prod     bool   `json:"production"`
	Ver      string `json:"version"`
}

func (p *Params) Cred() *force.Credential {
	return &force.Credential{
		ClientID:     p.ClientID,
		ClientSecret: p.ClientSecret,
		Username:     p.Username,
		Password:     p.Password,
		APIToken:     p.APIToken}
}

func (p *Params) Env() force.Env {
	if p.Prod {
		return force.Production
	} else {
		return force.Sandbox
	}
}

// LoadCredential loads force.com login parameter data from specified JSON file.
func Load(filename string) (*Params, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var params Params
	if err := json.Unmarshal(file, &params); err != nil {
		log.Printf("failed to parse %s: %s", file, string(file))
		return nil, err
	}
	return &params, nil
}
