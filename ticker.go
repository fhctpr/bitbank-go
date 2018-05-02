package bitbank

import (
	"context"
	"net/http"
	"strings"
)

type TickerService struct {
	client *Client
}

type TickerResponce struct {
	Success int `json:"success"`
	Data    struct {
		Sell      string `json:"sell"`
		Buy       string `json:"buy"`
		High      string `json:"high"`
		Low       string `json:"low"`
		Last      string `json:"last"`
		Vol       string `json:"vol"`
		Timestamp int    `json:"timestamp"`
	} `json:"data"`
}

func (t *TickerService) Get(ctx context.Context, pair string) (*TickerResponce, error) {
	pair = strings.ToLower(pair)
	req, err := t.client.newRequest(ctx, "GET", pair+"/ticker")
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var ticker TickerResponce
	if err := decodeBody(res, &ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}
