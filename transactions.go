package bitbank

import (
	"context"
	"net/http"
	"strings"
)

type TransactionsService struct {
	client *Client
}

type TransactionsResponce struct {
	Success int `json:"success"`
	Data    struct {
		Transactions []struct {
			TransactionID int    `json:"transaction_id"`
			Side          string `json:"side"`
			Price         string `json:"price"`
			Amount        string `json:"amount"`
			ExecutedAt    int    `json:"executed_at"`
		} `json:"transactions"`
	} `json:"data"`
}

func (t *TransactionsService) Get(ctx context.Context, pair string) (*TransactionsResponce, error) {
	pair = strings.ToLower(pair)
	req, err := t.client.newRequest(ctx, "GET", pair+"/transactions")
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var transactions TransactionsResponce
	if err := decodeBody(res, &transactions); err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (t *TransactionsService) GetByDate(ctx context.Context, pair, date string) (*TransactionsResponce, error) {
	pair = strings.ToLower(pair)
	req, err := t.client.newRequest(ctx, "GET", pair+"/transactions/"+date)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var transactions TransactionsResponce
	if err := decodeBody(res, &transactions); err != nil {
		return nil, err
	}

	return &transactions, nil
}
