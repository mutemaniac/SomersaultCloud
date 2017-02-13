package SomersaultCloud

import (
	"errors"
	"net/url"
	"reflect"
	"SomersaultCloud/models"
	"strconv"

	"strings"

	"github.com/astaxie/beego/httplib"
	urlQuery "github.com/google/go-querystring/query"
)

// AddPlugin Add Plugin to  a specific API
// pluginName required
// config dependce on the type of plugin
// consumerID -- 不是每个插件都支持consumer
// For every API and Consumer. Don't set api and consumer_id.
// For every API and a specific Consumer. Only set consumer_id.
// For every Consumer and a specific API. Only set api.
// For a specific Consumer and API. Set both api and consumer_id.
func AddPlugin(pluginName string, apiNameOrID string, consumerID string, config interface{}, retPlugin interface{}) error {
	//curl -X POST http://kong:8001/apis/{api}/plugins  --data "name={plugin name}"
	//log.Println("Enter AddPlugin: ", pluginName, apiNameOrID, consumerID)
	//validate paramters
	if len(pluginName) == 0 {
		return errors.New("The the name(type) of the plugin can not be nil")
	}
	configType, ok := models.PluginConfigs[pluginName]
	if config != nil {
		if !ok || !strings.HasSuffix(reflect.TypeOf(config).String(), configType) {
			return errors.New("The name of the plugin and the config type must be consistent")
		}
	}

	//Post http
	var url string
	if len(apiNameOrID) == 0 {
		url = kongAdminURL + `/plugins`
	} else {
		url = kongAdminURL + `/apis/` + apiNameOrID + `/plugins`
	}

	req := httplib.Post(url)
	req.Param("name", pluginName)
	if len(consumerID) > 0 {
		req.Param("consumer_id", consumerID)
	}
	if config != nil {
		configs, _ := urlQuery.Values(config)
		if configs != nil {
			for key, val := range configs {
				req.Param(key, strings.Join(val, ","))
			}
		}
	}
	//var retPlugin interface{} //models.Plugin
	err := req.ToJSON(&retPlugin)
	if err != nil {
		return err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return errors.New(retStr)
	}
	return nil
}

// GetPlugin Retrieve Plugin
// id required	The unique identifier of the plugin to retrieve
func GetPlugin(id string, retPlugin interface{}) error {
	//Get /plugins/{id}
	if len(id) == 0 {
		return errors.New("The unique identifier of the Plugin can not be null")
	}
	req := httplib.Get(kongAdminURL + `/plugins/` + id)
	//var retPlugin models.Plugin
	err := req.ToJSON(&retPlugin)
	if err != nil {
		return err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return errors.New(retStr)
	}
	return nil
}

// ListPlugin List All Plugins
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.ApiId -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
func ListPlugin(plugin interface{}, size int, offset string, retPluginList interface{}) error {
	//GET /plugins/

	//do get
	u, err := url.Parse(kongAdminURL + `/plugins/`)
	if err != nil {
		return err
	}
	urlValues, _ := urlQuery.Values(plugin)
	if size > 0 {
		urlValues.Add("size", strconv.Itoa(size))
	}
	if len(offset) > 0 {
		urlValues.Add("offset", offset)
	}

	u.RawQuery = urlValues.Encode()
	req := httplib.Get(u.String())

	//to json
	//var retPluginList models.PluginList
	err = req.ToJSON(&retPluginList)
	if err != nil {
		return err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return errors.New(retStr)
	}
	return nil
}

// ListPluginPerAPI List All Plugins for specific api
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.api_id -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
// apiNameOrId --
func ListPluginPerAPI(plugin interface{}, size int, offset string, apiNameOrID string, retPluginList interface{}) error {
	//GET /apis/{api name or id}/plugins/
	if len(apiNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the api can not be null")
	}
	//Do get
	u, err := url.Parse(kongAdminURL + `/apis/` + apiNameOrID + `/plugins/`)
	if err != nil {
		return err
	}
	urlValues, _ := urlQuery.Values(plugin)
	if size > 0 {
		urlValues.Add("size", strconv.Itoa(size))
	}
	if len(offset) > 0 {
		urlValues.Add("offset", offset)
	}

	u.RawQuery = urlValues.Encode()
	req := httplib.Get(u.String())

	//to json
	// var retPluginList models.PluginList
	err = req.ToJSON(&retPluginList)
	if err != nil {
		return err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return errors.New(retStr)
	}
	return nil
}

// UpdatePlugin Update the plugin identify by id.
// apiNameOrId (required) The unique identifier or the name of the API for which to update the plugin configuration.
// pluginID (required) The unique identifier of the plugin configuration to update on this API.
// config The configuration properties for the Plugin which want to update.
// retPlugin The updated plugin information.
func UpdatePlugin(apiNameOrID string, pluginID string, config interface{}, retPlugin interface{}) error {
	//PATCH /apis/{api name or id}/plugins/{id}
	if len(apiNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the api can not be null")
	}
	if len(pluginID) == 0 {
		return errors.New("The unique identifier or the name of the plugin can not be null")
	}

	url := kongAdminURL + `/apis/` + apiNameOrID + `/plugins/` + pluginID
	req := httplib.NewBeegoRequest(url, "PATCH")
	if config != nil {
		configs, _ := urlQuery.Values(config)
		if configs != nil {
			for key, val := range configs {
				req.Param(key, strings.Join(val, ","))
			}
		}
	}
	//var retPlugin interface{} //models.Plugin
	err := req.ToJSON(&retPlugin)
	if err != nil {
		return err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return errors.New(retStr)
	}
	return nil
}

// DeletePluginPerAPI Delete Plugin
// apiNameOrId -- required	The unique identifier or the name of the API for which to delete the plugin configuration
// id -- required	The unique identifier of the plugin configuration to delete on this API
func DeletePluginPerAPI(id string, apiNameOrID string) error {
	//DELETE /apis/{api name or id}/plugins/{id}
	if len(apiNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the API can not be null")
	}
	if len(id) == 0 {
		return errors.New("The unique identifier of the plugin can not be null")
	}

	req := httplib.Delete(kongAdminURL + `/apis/` + apiNameOrID + `/plugins/` + id)
	_, err := req.Response()
	if err != nil {
		return err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return errors.New(retStr)
	}
	return nil
}
