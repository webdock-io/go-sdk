package gosdk

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Webdock struct {
	Token         string
	Authorization string
	BASE_URL      string
	client        *http.Client
}

func (w *Webdock) GetFormatedToken() string {
	return fmt.Sprintf("Bearer %s", w.Token)

}

func (w *Webdock) GetFormatedURL(p string) string {
	apiURL, err := url.JoinPath(w.BASE_URL, p)
	if err != nil {
		log.Fatal(errors.New("failed to build API URL"))
	}
	return apiURL
}

type WebdockOptions struct {
	TOKEN string
}

func New(opts WebdockOptions) Webdock {
	return Webdock{
		Token:         opts.TOKEN,
		Authorization: "Authorization",
		BASE_URL:      "api.webdock.io",
		client:        &http.Client{},
	}
}
