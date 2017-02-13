package models

// The Consumer object represents a consumer - or a user - of an API.
type Consumer struct {
	Username string   `json:"username,omitempty"`   //: "guan",
	CustomID string   `json:"custom_id,omitempty"`  //: "abc123",
	CreateAt KongTime `json:"created_at,omitempty"` //: 1484277177000,
	ID       string   `json:"id,omitempty"`         //: "5878ca3a-13a8-4cee-8ac9-de2cdb588381"
}

//ConsumerList find all Consumers by page
type ConsumerList struct {
	Total  int        `json:"total,omitempty"` // total count of apis
	Data   []Consumer `json:"data,omitempty"`  // apis
	Next   string     `json:"next,omitempty"`  // next page url
	Offset string     `json:"offset,omitempty"`
}
