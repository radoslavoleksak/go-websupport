package websupport

import (
	"net/url"
	"net/http"
	"encoding/json"
	"io"
	"bytes"
	"encoding/base64"
)

const (
	DefaultEndpoint = "https://rest.websupport.sk"

	mediaType = "application/json"
)

type Client struct {
    BaseURL   *url.URL
    UserAgent string
 
	httpClient *http.Client

	headers map[string]string

	Users UserService
	DNS	  DNSService
}

func NewClient(username, password string, httpClient *http.Client) (*Client, error) {
    if httpClient == nil {
    	httpClient = http.DefaultClient
	}

	c := &Client {httpClient: httpClient}

	var err error
	c.BaseURL, err = url.Parse(DefaultEndpoint)
	if err != nil {
		return nil, err
	}

	c.headers = map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password)),
	}

	c.Users = &UserServiceImpl{client: c}
	c.DNS = &DNSServiceImpl{client: c}

	return c, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
    rel := &url.URL{Path: path}
    u := c.BaseURL.ResolveReference(rel)

    var buf io.ReadWriter
    if body != nil {
        buf = new(bytes.Buffer)
        err := json.NewEncoder(buf).Encode(body)
        if err != nil {
            return nil, err  
        }
    }

    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }
    if body != nil {  
        req.Header.Set("Content-Type", mediaType) 
    }
    req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	for k, v := range c.headers {
		req.Header.Add(k, v)
	}

    return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//fmt.Println(req.URL.String())
	//fmt.Println("ResponseBody: ")
	//fmt.Println(buf.String())

	err = json.NewDecoder(resp.Body).Decode(v)
    return resp, err
}