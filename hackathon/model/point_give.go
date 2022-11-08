package model

type AllGiveResForHTTPGet struct {
	Id         string `json:"id"`
	FromUserId string `json:"fromUserId"`
	Points     int    `json:"points"`
	Message    string `json:"message"`
	ToUserId   string `json:"toUserId"`
}

type GiveReqHTTPPost struct {
	FromName string
	Points   int
	Message  string
	ToName   string
	//FromUserId string `json:"fromUserId"`
	//Points     int    `json:"points"`
	//Message    string `json:"message"`
	//ToUserId   string `json:"toUserId"`
}

type GiveReqHTTPPut struct {
	Id      string
	Points  int
	Message string
}
type GiveReqHTTPDelete struct {
	Id string //point取引のid

}
