package bitbank

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CandlestickService struct {
	client *Client
}

type CandlestickResponce struct {
	Success int
	Data    struct {
		Candlestick struct {
			Type  string
			Ohlcv []Ohlcv
		}
		Timestamp int64
	}
}
type Ohlcv struct {
	Open      string
	Low       string
	High      string
	Close     string
	Volume    string
	Timestamp int64
}
type CandlestickDefaultResponce struct {
	Success int `json:"success"`
	Data    struct {
		Candlestick []struct {
			Type  string        `json:"type"`
			Ohlcv []interface{} `json:"ohlcv"`
		} `json:"candlestick"`
		Timestamp int64 `json:"timestamp"`
	} `json:"data"`
}

func (c *CandlestickService) GetByDate(ctx context.Context, pair, candleType, dateStr string) (*CandlestickResponce, error) {
	err := isValidCandleType(candleType)
	if err != nil {
		return nil, err
	}
	pair = strings.ToLower(pair)
	req, err := c.client.newRequest(ctx, "GET", pair+"/candlestick/"+candleType+"/"+dateStr)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var cd CandlestickDefaultResponce
	if err := decodeBody(res, &cd); err != nil {
		return nil, err
	}
	candle, err := parseCandlestick(cd)
	if err != nil {
		return nil, err
	}

	return &candle, nil
}

func parseCandlestick(cd CandlestickDefaultResponce) (CandlestickResponce, error) {
	var candle CandlestickResponce
	candle.Success = cd.Success
	candle.Data.Timestamp = cd.Data.Timestamp
	candle.Data.Candlestick.Type = cd.Data.Candlestick[0].Type
	candle.Data.Candlestick.Ohlcv = make([]Ohlcv, len(cd.Data.Candlestick[0].Ohlcv))
	for i := range cd.Data.Candlestick[0].Ohlcv {
		candle.Data.Candlestick.Ohlcv[i].Open = cd.Data.Candlestick[0].Ohlcv[i].([]interface{})[0].(string)
		candle.Data.Candlestick.Ohlcv[i].High = cd.Data.Candlestick[0].Ohlcv[i].([]interface{})[1].(string)
		candle.Data.Candlestick.Ohlcv[i].Low = cd.Data.Candlestick[0].Ohlcv[i].([]interface{})[2].(string)
		candle.Data.Candlestick.Ohlcv[i].Close = cd.Data.Candlestick[0].Ohlcv[i].([]interface{})[3].(string)
		candle.Data.Candlestick.Ohlcv[i].Volume = cd.Data.Candlestick[0].Ohlcv[i].([]interface{})[4].(string)
		tstr := strconv.FormatFloat(cd.Data.Candlestick[0].Ohlcv[i].([]interface{})[5].(float64), 'f', -1, 64)
		t, _ := strconv.ParseInt(tstr, 10, 64)
		candle.Data.Candlestick.Ohlcv[i].Timestamp = t
	}

	return candle, nil
}

func isValidCandleType(candleType string) error {
	var c []string = []string{"1min", "5min", "15min", "30min", "1hour", "4hour", "8hour", "12hour", "1day", "1week"}
	for i := range c {
		if candleType == c[i] {
			return nil
		}
	}

	return fmt.Errorf("Unable to find the candle type. %v", candleType)
}
