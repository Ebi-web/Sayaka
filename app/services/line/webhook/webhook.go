package webhook

type Request struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

type Event struct {
	Type    string `json:"type"`
	Message struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"message"`
	WebhookEventID  string `json:"webhookEventId"`
	DeliveryContext struct {
		IsRedelivery bool `json:"isRedelivery"`
	} `json:"deliveryContext"`
	Timestamp int64 `json:"timestamp"`
	Source    struct {
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"source"`
	ReplyToken string `json:"replyToken"`
	Mode       string `json:"mode"`
}
