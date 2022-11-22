package model

type AffiliationResForHTTPGet struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}
type AffiliationReqForHTTPPost struct {
	Name string
}
