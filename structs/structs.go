package structs

type ExchangeResult struct {
	Id          string `json:"exchange.id"`
	VPToken     string `json:"exchange.variables.results.default.vpToken"`
	TimeCreated string `json:"exchange.updatedAt"`
}

type Exchange struct {
	Id   string `json:"id"`
	Link string `json:"OID4VP"`
	Qr   string `json:"QR"`
}
