package model

type AllTakeResForHTTPGet struct {
	Id         string `json:"id"`
	FromUserId string `json:"fromUserId"`
	Points     int    `json:"points"`
	Message    string `json:"message"`
	ToUserId   string `json:"toUserId"`
}
