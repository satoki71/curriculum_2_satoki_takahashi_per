package model

type UserResForHTTPGet struct {
	UserId        string `json:"userId"`
	Name          string `json:"name"`
	AffiliationId string `json:"affiliationId"`
	Points        int    `json:"points"`
}
type AllUserResForHTTPGet struct {
	UserId        string `json:"userId"`
	Name          string `json:"name"`
	AffiliationId string `json:"affiliationId"`
	Points        int    `json:"points"`
}
type MemberUserResForHTTPGet struct {
	UserId        string `json:"userId"`
	Name          string `json:"name"`
	AffiliationId string `json:"affiliationId"`
	Points        int    `json:"points"`
}
type UserReqHTTPPost struct {
	Name        string
	Affiliation string
}

type UserIdReqPost struct {
	UserId string `json:"userId"`
}

type UserReqHTTPUpdate struct {
	Name string
}
