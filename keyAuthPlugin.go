package SomersaultCloud

import (
	"errors"
	"net/url"
	"SomersaultCloud/models"
	"strconv"

	"github.com/astaxie/beego/httplib"
)

// AddKeyAuthPlugin2Api Add Plugin to  a specific API
func AddKeyAuthPlugin2Api(apiNameOrID string) (*models.KeyAuthPlugin, error) {
	//curl -X POST http://kong:8001/apis/{api}/plugins  --data "name=key-auth"
	if len(apiNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	var retKeyAuthPlugin models.KeyAuthPlugin
	err := AddPlugin("key-auth", apiNameOrID, "", nil, &retKeyAuthPlugin)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPlugin, nil
}

// GetKeyAuthPlugin Retrieve Plugin
// id required	The unique identifier of the plugin to retrieve
func GetKeyAuthPlugin(id string) (*models.KeyAuthPlugin, error) {
	//Get /plugins/{id}
	if len(id) == 0 {
		return nil, errors.New("The unique identifier of the KeyAuthPlugin can not be null")
	}
	var retKeyAuthPlugin models.KeyAuthPlugin
	err := GetPlugin(id, &retKeyAuthPlugin)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPlugin, nil
}

// ListKeyAuthPlugin List All Plugins
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.ApiId -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
func ListKeyAuthPlugin(plugin models.KeyAuthPlugin, size int, offset string) (*models.KeyAuthPluginList, error) {
	//GET /plugins/
	var retKeyAuthPluginList models.KeyAuthPluginList
	err := ListPlugin(plugin, size, offset, &retKeyAuthPluginList)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPluginList, nil
}

// ListKeyAuthPluginPerApi List All Plugins for specific api
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.api_id -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
// apiNameOrId --
func ListKeyAuthPluginPerApi(plugin models.KeyAuthPlugin, size int, offset string, apiNameOrID string) (*models.KeyAuthPluginList, error) {
	//GET /apis/{api name or id}/plugins/
	if len(apiNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	var retKeyAuthPluginList models.KeyAuthPluginList
	err := ListPluginPerAPI(plugin, size, offset, apiNameOrID, retKeyAuthPluginList)
	if err != nil {
		return nil, err
	}
	return &retKeyAuthPluginList, nil
}

//UpdateKeyAuthPlugin Update the key auth plugin.
func UpdateKeyAuthPlugin(apiNameOrID string, pluginID string, config *models.KeyAuthConfig) (*models.KeyAuthPlugin, error) {
	if len(apiNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	if len(pluginID) == 0 {
		return nil, errors.New("The unique identifier or the name of the plugin can not be null")
	}
	var retKeyAuthPlugin models.KeyAuthPlugin
	err := UpdatePlugin(apiNameOrID, pluginID, config, &retKeyAuthPlugin)
	if err != nil {
		return nil, err
	}

	return &retKeyAuthPlugin, nil
}

// DeleteKeyAuthPluginPerApi Delete Plugin
// apiNameOrId -- required	The unique identifier or the name of the API for which to delete the plugin configuration
// id -- required	The unique identifier of the plugin configuration to delete on this API
func DeleteKeyAuthPluginPerApi(id string, apiNameOrID string) error {
	//DELETE /apis/{api name or id}/plugins/{id}
	if len(apiNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the API can not be null")
	}
	if len(id) == 0 {
		return errors.New("The unique identifier of the plugin can not be null")
	}
	return DeletePluginPerAPI(id, apiNameOrID)
}

// CreateAPIKey Create an API Key
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// apikey -- (optional)You can optionally set your own unique key to authenticate the client. If missing, the plugin will generate one.
func CreateAPIKey(consumerNameOrID string, apikey string) (*models.ApiKey, error) {
	//POST http://kong:8001/consumers/{consumerNameOrId}/key-auth -d '' HTTP/1.1 201 Created
	if len(consumerNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	req := httplib.Post(kongAdminURL + `/consumers/` + consumerNameOrID + `/key-auth/`)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	if len(apikey) > 0 {
		req.Param("key", apikey)
	}

	var retApikey models.ApiKey
	err := req.ToJSON(&retApikey)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retApikey, nil
}

// GetAPIKey Get apikey by consumer & apikey id.
func GetAPIKey(consumerNameOrID string, apikeyid string) (*models.ApiKey, error) {
	// GET http://13.76.42.81:8001/consumers/guan/key-auth/a0fdba77-fc6d-4632-845c-cc1a623cf59d
	if len(consumerNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(apikeyid) == 0 {
		return nil, errors.New("The unique identifier of the apikey can not be null")
	}
	req := httplib.Get(kongAdminURL + `/consumers/` + consumerNameOrID + `/key-auth/` + apikeyid)
	var retApikey models.ApiKey
	err := req.ToJSON(&retApikey)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retApikey, nil
}

// ListAPIKey List api keys of one users.
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
func ListAPIKey(consumerNameOrID string, size int, offset string) (*models.ApiKeyList, error) {
	//Get http://kong:8001/consumers/{consumerNameOrId}/key-auth?size=1 -d '' HTTP/1.1 201 Created
	if len(consumerNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	//do get
	u, err := url.Parse(kongAdminURL + `/consumers/` + consumerNameOrID + `/key-auth`)
	if err != nil {
		return nil, err
	}
	urlValues := u.Query()
	if size > 0 {
		urlValues.Add("size", strconv.Itoa(size))
	}
	if len(offset) > 0 {
		urlValues.Add("offset", offset)
	}
	u.RawQuery = urlValues.Encode()
	req := httplib.Get(u.String())
	//to json
	var retApikeyList models.ApiKeyList
	err = req.ToJSON(&retApikeyList)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retApikeyList, nil
}

// DeleteAPIKey Delete an API Key by id.
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// apikeyid -- api_key id
func DeleteAPIKey(consumerNameOrID string, apikeyid string) error {
	//Delete http://kong:8001/consumers/{consumerNameOrId}/key-auth/{apikeyid} -d '' HTTP/1.1 201 Created
	if len(consumerNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(apikeyid) == 0 {
		return errors.New("The unique identifier of the apikey can not be null")
	}
	req := httplib.Delete(kongAdminURL + `/consumers/` + consumerNameOrID + `/key-auth/` + apikeyid)
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
