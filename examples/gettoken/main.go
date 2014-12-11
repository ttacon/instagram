package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ttacon/pretty"
)
import "code.google.com/p/goauth2/oauth"

var (
	clientId     = flag.String("cid", "", "id of your client")
	clientSecret = flag.String("csec", "", "your client secret")

	oauthConfig *oauth.Config
)

func main() {
	flag.Parse()

	oauthConfig = &oauth.Config{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		Scope:        "basic", // so even though instagram says scope is optional, it isn't...that's annoying
		RedirectURL:  "http://localhost:7371/consume_oauth",
		AuthURL:      "https://api.instagram.com/oauth/authorize/",
		TokenURL:     "https://api.instagram.com/oauth/access_token",
	}

	http.HandleFunc("/", landing)
	http.HandleFunc("/consume_oauth", consume)
	http.ListenAndServe(":7371", nil)
}

// from code.google.com/p/goauth2/oauth example:
// (visible at http://godoc.org/code.google.com/p/goauth2/oauth)

// A landing page redirects to the OAuth provider to get the auth code.
func landing(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, oauthConfig.AuthCodeURL("foo"), http.StatusFound)
}

// The user will be redirected back to this handler, that takes the
// "code" query parameter and Exchanges it for an access token.
func consume(w http.ResponseWriter, r *http.Request) {
	t := &oauth.Transport{Config: oauthConfig}
	tok, err := t.Exchange(r.FormValue("code"))
	// The Transport now has a valid Token. Create an *http.Client
	// with which we can make authenticated API requests.
	pretty.Println(t)
	fmt.Println("err: ", err, "\ntok: ")
	pretty.Println(tok)
	// ...
	// btw, r.FormValue("state") == "foo"
}
