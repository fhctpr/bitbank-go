package bitbank

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"
	//	"encoding/base64"
	// "io/ioutil"
	//	"time"
)

const (
	PublicBaseURL  = "https://public.bitbank.cc"
	PrivateBaseURL = "https://api.bitbank.cc"
)

var userAgent = fmt.Sprintf("bitbankgo/%s (%s)",
	"0.0.1", runtime.Version())

type Client struct {
	Url
	HTTPClient *http.Client

	ApiKey    string
	ApiSecret string

	Ticker       *TickerService
	Depth        *DepthService
	Transactions *TransactionsService
	Candlestick  *CandlestickService

	Assets *AssetsService
}

type Url struct {
	Public  *url.URL
	Private *url.URL
}

func New() *Client {
	public, _ := url.Parse(PublicBaseURL)
	private, _ := url.Parse(PrivateBaseURL)

	c := &Client{
		Url: Url{
			Public:  public,
			Private: private,
		},
	}
	c.Ticker = &TickerService{client: c}
	c.Depth = &DepthService{client: c}
	c.Transactions = &TransactionsService{client: c}
	c.Candlestick = &CandlestickService{client: c}

	c.Assets = &AssetsService{client: c}

	return c
}

func (c *Client) Auth(key, secret string) *Client {
	c.ApiKey = key
	c.ApiSecret = secret

	return c
}

func (c *Client) newRequest(ctx context.Context, method, spath string) (*http.Request, error) {
	ref, err := url.Parse(spath)
	if err != nil {
		return nil, err
	}

	u := c.Url.Public.ResolveReference(ref)
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	return req, nil
}

func (c *Client) newPrivateRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	req, err := c.newRequest(ctx, method, spath)
	if err != nil {
		return nil, err
	}
	nonce := getNonce()
	signature := signMessage(c.ApiSecret, nonce, spath)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("ACCESS-KEY", c.ApiKey)
	req.Header.Add("ACCESS-NONCE", nonce)
	req.Header.Add("ACCESS-SIGNATURE", signature)

	return req, nil
}

func getNonce() string {
	nonce := uint64(time.Now().UnixNano())
	return strconv.FormatUint(nonce, 10)
}

func signMessage(m ...string) string {
	h := hmac.New(sha256.New, []byte(m[0]))
	h.Write([]byte(m[1]))

	return ""
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
