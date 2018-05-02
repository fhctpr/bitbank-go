package bitbank

import (
	"context"
	"net/http"
	"strings"
)

type DepthService struct {
	client *Client
}

type DepthResponce struct {
	Success int `json:"success"`
	Data    struct {
		Asks      [][]string `json:"asks"`
		Bids      [][]string `json:"bids"`
		Timestamp int64      `json:"timestamp"`
	} `json:"data"`
}

func (d *DepthService) Get(ctx context.Context, pair string) (*DepthResponce, error) {
	pair = strings.ToLower(pair)
	req, err := d.client.newRequest(ctx, "GET", pair+"/depth")
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var depth DepthResponce
	if err := decodeBody(res, &depth); err != nil {
		return nil, err
	}

	return &depth, nil
}
