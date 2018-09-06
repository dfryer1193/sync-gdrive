package sync

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Credentials which stores google ids
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

var (
	cred  Credentials
	conf  *oauth2.Config
	state string
	ctx   context.Context
)

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func init() {
	var c Credentials
	file, err := ioutil.ReadFile("./client_id.json")
	if err != nil {
		fmt.Printf("File errpr: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &c)

	ctx = context.Background()

	conf = &oauth2.Config{
		ClientID:     c.Cid,
		ClientSecret: c.Csecret,
		RedirectURL:  "http://localhost:9090/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/drive",
		},
		Endpoint: google.Endpoint,
	}
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}

func authHandler() {

}
