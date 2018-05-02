package bitbank

type AssetsService struct {
	client *Client
}

type AssetInfo struct {
	Success int `json:"success"`
	Data    struct {
		Assets []struct {
			Asset           string `json:"asset"`
			AmountPrecision int    `json:"amount_precision"`
			OnhandAmount    string `json:"onhand_amount"`
			LockedAmount    string `json:"locked_amount"`
			FreeAmount      string `json:"free_amount"`
			WithdrawalFee   string `json:"withdrawal_fee"`
		} `json:"assets"`
	} `json:"data"`
}

/*
func (a *AssetsService) GetInfo(ctx context.Context) (AssetInfo, error) {
	req, err := a.client.newPrivateRequest(ctx, "GET", "/user/assets", nil)

	return AssetInfo{}, nil
}
*/
