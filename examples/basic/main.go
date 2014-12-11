package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ttacon/instagram"
	"github.com/ttacon/pretty"
)

var (
	tok    = flag.String("tok", "", "access token to use")
	userID = flag.String("u", "", "id of user to retrieve")
)

func main() {
	flag.Parse()

	c := instagram.NewClient(http.DefaultClient, *tok)
	user, err := c.User(*userID)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		pretty.Println(user)
	}
}
