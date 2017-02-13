package models

// PluginConfigs is the collection of all plugins. The key is the name of that plugin.
// The value is the name of that plugin config type.
var PluginConfigs map[string]string = make(map[string]string)

//Plugin kong plugin model
type Plugin struct {
	ID         string      `json:"id,omitempty" url:"id,omitempty"`
	ApiId      string      `json:"api_id,omitempty" url:"api_id,omitempty"`
	ConsumerId string      `json:"consumer_id,omitempty" url:"consumer_id,omitempty"`
	Name       string      `json:"name,omitempty" url:"name,omitempty"`
	Config     interface{} `json:"config,omitempty" url:"-"`
	Enabled    bool        `json:"enabled,omitempty" url:"-"`
	CreatedAt  KongTime    `json:"created_at,omitempty" url:"-"`
}

//PluginList find all apis by page
type PluginList struct {
	Total  int      `json:"total,omitempty"` // total count of apis
	Data   []Plugin `json:"data,omitempty"`  // apis
	Next   string   `json:"next,omitempty"`  // next page url
	Offset string   `json:"offset,omitempty"`
}
