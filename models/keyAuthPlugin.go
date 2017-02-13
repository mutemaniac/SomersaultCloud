package models

func init() {
	PluginConfigs["key-auth"] = "KeyAuthConfig"
}

type KeyAuthPlugin struct {
	ID         string        `json:"id,omitempty" url:"id,omitempty"`                   //: "4d924084-1adb-40a5-c042-63b19db421d1",
	ApiId      string        `json:"api_id,omitempty" url:"api_id,omitempty"`           //: "5fd1z584-1adb-40a5-c042-63b19db49x21",
	ConsumerId string        `json:"consumer_id,omitempty" url:"consumer_id,omitempty"` //: "a3dX2dh2-1adb-40a5-c042-63b19dbx83hF4",
	Name       string        `json:"name,omitempty" url:"name,omitempty"`               //: "rate-limiting",
	Config     KeyAuthConfig `json:"config,omitempty" url:"-"`                          //:///
	Enabled    bool          `json:"enabled,omitempty" url:"-"`                         //: true,
	CreatedAt  KongTime      `json:"created_at,omitempty" url:"-"`                      //: 1422386534
}

type KeyAuthConfig struct {
	// Describes an array of comma separated parameter names where the plugin will look for a key.
	// The client must send the authentication key in one of those key names, and the plugin will
	// try to read the credential from a header or the querystring parameter with the same name.
	KeyNames []string `json:"key_names,omitempty" url:"config.key_names,omitempty"`

	// An optional boolean value telling the plugin to hide the credential to the upstream API server.
	// It will be removed by Kong before proxying the request.
	HideCredentials bool `json:"hide_credentials,omitempty" url:"config.hide_credentials,omitempty"`
}

type ApiKey struct {
	Key        string   `json:"key,omitempty"`         //: "1234567890",
	ConsumerId string   `json:"consumer_id,omitempty"` //: "5878ca3a-13a8-4cee-8ac9-de2cdb588381",
	CreatedAt  KongTime `json:"created_at,omitempty"`  //: 1484278640000,
	ID         string   `json:"id,omitempty"`          //: "88ca1d4a-cc29-4856-98e2-36c08afbf698"
}

//KeyAuthPluginList find all apis by page
type KeyAuthPluginList struct {
	Total  int             `json:"total,omitempty"` // total count of apis
	Data   []KeyAuthPlugin `json:"data,omitempty"`  // apis
	Next   string          `json:"next,omitempty"`  // next page url
	Offset string          `json:"offset,omitempty"`
}

type ApiKeyList struct {
	Total  int      `json:"total,omitempty"` // total count of apis
	Data   []ApiKey `json:"data,omitempty"`  // apis
	Next   string   `json:"next,omitempty"`  // next page url
	Offset string   `json:"offset,omitempty"`
}
