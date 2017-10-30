package dto

type PlayerCountResponse struct {
	Response PlayerCountResponseData `json:"response"`
}

type PlayerCountResponseData struct {
	PlayerCount int `json:"player_count"`
	Result int `json:"result"`
}

type AssetPriceResponse struct {
	Result AssetPriceResponseData `json:"result"`
}

type AssetPriceResponseData struct {
	Success bool `json:"success"`
	Assets []AssetPricePricesListing
}

type AssetPricePricesListing struct {
	Prices CurrencyPrices `json:"prices"`
	OriginalPrices CurrencyPrices `json:"original_prices"`
}

type CurrencyPrices struct {
	USD int
	GBP int
	EUR int
	JPY int
	KRW int
	CAD int
	AUD int
	CNY int
	HKD int
}
