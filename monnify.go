package monnify

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// defaultHTTPTimeout is the default timeout on the http client
	defaultHTTPTimeout = 15 * time.Second

	// defaultBaseURL is the base endpoint for all requests
	defaultBaseURL = "https://sandbox.monnify.com"
)

// Metadata is a key-value helper
type Metadata map[string]interface{}

type Client struct {
	client      *http.Client
	baseURL     string
	accessToken string
	apiKey      string
	secretKey   string
}

func New(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if IsStringEmpty(c.baseURL) {
		c.baseURL = defaultBaseURL
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base url provided")
	}

	c.baseURL = u.String()

	if c.client == nil {
		c.client = &http.Client{
			Timeout: defaultHTTPTimeout,
		}
	}

	return c, nil
}

type Response struct {
	ResponseMessage string          `json:"responseMessage"`
	ResponseCode    string          `json:"responseCode"`
	ResponseBody    json.RawMessage `json:"responseBody,omitempty"`
}

func (c *Client) doRequest(method, url string, body, v interface{}) (*http.Response, error) {
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseURL, url), buf)
	if err != nil {
		return nil, err
	}

	err = c.generateAccessToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var resp Response

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("request failed. monnify responded with status %d and message: %s",
			res.StatusCode,
			resp.ResponseMessage,
		)
	}

	if resp.ResponseBody != nil && v != nil {
		if err := json.Unmarshal(resp.ResponseBody, &v); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// generateAccessToken generates a bearer token to access Monnify API
func (c *Client) generateAccessToken() error {
	var r struct {
		ResponseMessage string `json:"responseMessage"`
		ResponseCode    string `json:"responseCode"`
		ResponseBody    struct {
			AccessToken string `json:"accessToken"`
			ExpiresIn   int    `json:"expiresIn"`
		} `json:"responseBody,omitempty"`
	}

	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.apiKey, c.secretKey)))

	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s%s", c.baseURL, "/api/v1/auth/login"), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicToken))

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}

	if resp.StatusCode > http.StatusCreated {
		return fmt.Errorf("failed to get access token. monnify responded with status %d and message: %s",
			resp.StatusCode,
			r.ResponseMessage,
		)
	}

	c.accessToken = r.ResponseBody.AccessToken

	return nil
}

type Option func(c *Client)

func WithHTTPClient(cl *http.Client) Option {
	return func(c *Client) {
		c.client = cl
	}
}

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithAPIKey(s string) Option {
	return func(c *Client) {
		c.apiKey = s
	}
}

func WithSecretKey(s string) Option {
	return func(c *Client) {
		c.secretKey = s
	}
}

// IsStringEmpty checks is string is empty
func IsStringEmpty(s string) bool { return len(strings.TrimSpace(s)) == 0 }
