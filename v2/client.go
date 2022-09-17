package deeplapi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

const APIVersion string = "v2"

var (
	domain = "api.deepl.com"
	mux    = new(sync.RWMutex)
)

func Domain() string {
	mux.RLock()
	defer mux.RUnlock()
	return domain
}

// SetDomain
//
// - 'api.deepl.com' for pro account
// - 'api-free.deepl.com' for free account
func SetDomain(url_ string) {
	mux.Lock()
	defer mux.Unlock()
	domain = url_
}

func BaseURL() string {
	return "https://" + Domain()
}

// =================================================
// API Client
// =================================================

type Client struct {
	*http.Client
	AuthKey *string
}

func New(authKey string) *Client {
	return &Client{
		Client:  http.DefaultClient,
		AuthKey: &authKey,
	}
}

func (c *Client) doRequest(ctx context.Context, method, _url string, body io.ReadSeeker, param url.Values) ([]byte, error) {
	if c.AuthKey == nil {
		return nil, fmt.Errorf("AuthKey is nil")
	}
	var result []byte
	if param == nil {
		param = url.Values{}
	}
	req, err := http.NewRequestWithContext(ctx, method, _url, body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Accept", "application/json,plain/text")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+*c.AuthKey)
	req.URL.RawQuery = param.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do-request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK { // TODO(micheam): Handle errors
		b := new(bytes.Buffer)
		log.Printf("doRequest: %s: %s\n", resp.Status, b.String())
		return nil, errors.New(resp.Status)
	}
	result, err = ioutil.ReadAll(resp.Body) // TODO(micheam): use io.ReadAll() instead
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}
	return result, nil
}
