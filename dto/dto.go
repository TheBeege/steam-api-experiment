package dto

type PlayerCountResponse struct {
	Response PlayerCountResponseData `json:"response"`
}

type PlayerCountResponseData struct {
	PlayerCount int `json:"player_count"`
	Result int `json:"result"`
}
