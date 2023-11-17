// Little wrapper library for making HTTP request and marhsalling data into structs
package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client
}

func NewHTTPClient(timeout time.Duration) *httpClient {
	return &httpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *httpClient) Get(ctx context.Context, url string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (h *httpClient) Post(ctx context.Context, url string, payload, response interface{}) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (h *httpClient) NewTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return ctx, cancel
}
