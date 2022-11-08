package model

type UserResForHTTPGet struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AffiliationId string `json:"affiliationId"`
	Points        int    `json:"points"`
}
type AllUserResForHTTPGet struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AffiliationId string `json:"affiliationId"`
	Points        int    `json:"points"`
}
type MemberUserResForHTTPGet struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AffiliationId string `json:"affiliationId"`
	Points        int    `json:"points"`
}
type UserReqHTTPPost struct {
	Name        string
	Affiliation string
}

type UserIdReqPost struct {
	Id string `json:"id"`
}

type UserReqHTTPUpdate struct {
	Name string
}
