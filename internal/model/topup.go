package model

type TopupSchema struct {
	NotificationID string `json:"NotificationID"`
	EventBeginTime string `json:"EventBeginTime"`
	MSISDN         string `json:"MSISDN"`
	EventType      string `json:"EventType"`
	Parameters     struct {
		RequestID             string `json:"RequestID"`
		BalanceType           string `json:"BalanceType"`
		Amount                string `json:"Amount"`
		NewBalance            string `json:"NewBalance"`
		Validity              string `json:"Validity"`
		NewValidity           string `json:"NewValidity"`
		SubscriberPricePlanID string `json:"SubscriberPricePlanID"`
		AccountType           string `json:"AccountType"`
		IMSI                  string `json:"IMSI"`
		Language              string `json:"Language"`
		ChainID               string `json:"ChainID"`
	} `json:"parameters"`
}
