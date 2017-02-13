package models

//API is api model of kong
type API struct {
	ID               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`               //"Mockbin",
	RequestHost      string   `json:"request_host,omitempty"`       //"mockbin.com",
	RequestPath      string   `json:"request_path,omitempty"`       //"/someservice",
	StripRequestPath bool     `json:"strip_request_path,omitempty"` //false,
	PreserveHost     bool     `json:"preserve_host,omitempty"`      //false,
	UpstreamURL      string   `json:"upstream_url,omitempty"`       //"https://mockbin.com"
	CreateAt         KongTime `json:"created_at,omitempty"`
}

//APIList find all apis by page
type APIList struct {
	Total  int    `json:"total,omitempty"` // total count of apis
	Data   []API  `json:"data,omitempty"`  // apis
	Next   string `json:"next,omitempty"`  // next page url
	Offset string `json:"offset,omitempty"`
}
