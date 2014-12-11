package instagram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://api.instagram.com/v1"

type Client interface {
	User(string) (*User, error)
}

type client struct {
	tok string
	h   *http.Client
}

func NewClient(c *http.Client, tok string) Client {
	return &client{
		tok: tok,
		h:   c,
	}
}

func (c *client) User(userid string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s?access_token=%s", BASE_URL, userid, c.tok), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.h.Do(req)
	if err != nil {
		return nil, err
	}

	dbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data UserResponse
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dbytes, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, err
}
