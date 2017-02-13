package models

func init() {
	PluginConfigs["acl"] = "ACLConfig"
}

// ACLPlugin -- Restrict access to an API by whitelisting or blacklisting consumers using arbitrary ACL group names.
// This plugin requires an authentication plugin to have been already enabled on the API.
type ACLPlugin struct {
	ID         string    `json:"id,omitempty" url:"id,omitempty"`                   //: "4d924084-1adb-40a5-c042-63b19db421d1",
	ApiId      string    `json:"api_id,omitempty" url:"api_id,omitempty"`           //: "5fd1z584-1adb-40a5-c042-63b19db49x21",
	ConsumerId string    `json:"consumer_id,omitempty" url:"consumer_id,omitempty"` //: "a3dX2dh2-1adb-40a5-c042-63b19dbx83hF4",
	Name       string    `json:"name,omitempty" url:"name,omitempty"`               //: "rate-limiting",
	Config     ACLConfig `json:"config,omitempty" url:"-"`                          //:///
	Enabled    bool      `json:"enabled,omitempty" url:"-"`                         //: true,
	CreatedAt  KongTime  `json:"created_at,omitempty" url:"-"`                      //: 1422386534
}

// ACLConfig Configuration of ACL plugin
type ACLConfig struct {
	Whitelist []string `json:"whitelist,omitempty" url:"config.whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty" url:"config.blacklist,omitempty"`
}

type ACLGroup struct {
	Group      string   `json:"group,omitempty"`       //: "1234567890",
	ConsumerId string   `json:"consumer_id,omitempty"` //: "5878ca3a-13a8-4cee-8ac9-de2cdb588381",
	CreatedAt  KongTime `json:"created_at,omitempty"`  //: 1484278640000,
	ID         string   `json:"id,omitempty"`          //: "88ca1d4a-cc29-4856-98e2-36c08afbf698"
}

//ACLPluginList find all apis by page
type ACLPluginList struct {
	Total  int         `json:"total,omitempty"` // total count of apis
	Data   []ACLPlugin `json:"data,omitempty"`  // apis
	Next   string      `json:"next,omitempty"`  // next page url
	Offset string      `json:"offset,omitempty"`
}

type ACLGroupList struct {
	Total  int        `json:"total,omitempty"` // total count of apis
	Data   []ACLGroup `json:"data,omitempty"`  // apis
	Next   string     `json:"next,omitempty"`  // next page url
	Offset string     `json:"offset,omitempty"`
}
