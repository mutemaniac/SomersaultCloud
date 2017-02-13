package kong

import (
	"errors"
	"net/url"
	"service-cloud/utils/kong/models"
	"strconv"

	"github.com/astaxie/beego/httplib"
)

// AddACLPlugin2API Add Plugin to  a specific API
func AddACLPlugin2API(apiNameOrID string, whitelist []string, blacklist []string) (*models.ACLPlugin, error) {
	//curl -X POST http://kong:8001/apis/{api}/plugins  --data "name=acls"
	if len(apiNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	var retACLPlugin models.ACLPlugin
	var aclConfig models.ACLConfig
	aclConfig.Blacklist = blacklist
	aclConfig.Whitelist = whitelist
	err := AddPlugin("acl", apiNameOrID, "", &aclConfig, &retACLPlugin)
	if err != nil {
		return nil, err
	}
	return &retACLPlugin, nil
}

// GetACLPlugin Retrieve Plugin
// id required	The unique identifier of the plugin to retrieve
func GetACLPlugin(id string) (*models.ACLPlugin, error) {
	//Get /plugins/{id}
	if len(id) == 0 {
		return nil, errors.New("The unique identifier of the ACLPlugin can not be null")
	}
	var retACLPlugin models.ACLPlugin

	err := GetPlugin(id, &retACLPlugin)
	if err != nil {
		return nil, err
	}
	return &retACLPlugin, nil
}

// ListACLPlugin List All Plugins
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.ApiId -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
func ListACLPlugin(plugin models.Plugin, size int, offset string) (*models.ACLPluginList, error) {
	var retACLPluginList models.ACLPluginList
	err := ListPlugin(plugin, size, offset, &retACLPluginList)
	if err != nil {
		return nil, err
	}
	return &retACLPluginList, nil
}

// ListACLPluginPerAPI List All Plugins for specific api
// plugin.id -- optional	A filter on the list based on the id field.
// plugin.name -- optional	A filter on the list based on the name field.
// plugin.api_id -- optional	A filter on the list based on the api_id field.
// plugin.consumer_id -- optional	A filter on the list based on the consumer_id field.
// plugin.size -- optional, default is 100	A limit on the number of objects to be returned.
// plugin.offset -- optional	A cursor used for pagination. offset is an object identifier that defines a place in the list.
// apiNameOrId --
func ListACLPluginPerAPI(plugin models.Plugin, size int, offset string, apiNameOrID string) (*models.ACLPluginList, error) {
	//GET /apis/{api name or id}/plugins/
	if len(apiNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	var retACLPluginList models.ACLPluginList
	err := ListPluginPerAPI(plugin, size, offset, apiNameOrID, &retACLPluginList)
	if err != nil {
		return nil, err
	}
	return &retACLPluginList, nil
}

// UpdateACLPlugin update the ACL plugin
// apiNameOrId (required) The unique identifier or the name of the API for which to update the plugin configuration
// pluginID (required) The unique identifier of the plugin configuration to update on this API
// config The configuration properties for the Plugin which want to update.
func UpdateACLPlugin(apiNameOrID string, pluginID string, config *models.ACLConfig) (*models.ACLPlugin, error) {
	if len(apiNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the api can not be null")
	}
	if len(pluginID) == 0 {
		return nil, errors.New("The unique identifier or the name of the plugin can not be null")
	}
	var retACLPlugin models.ACLPlugin
	err := UpdatePlugin(apiNameOrID, pluginID, config, &retACLPlugin)
	if err != nil {
		return nil, err
	}

	return &retACLPlugin, nil
}

// DeleteACLPluginPerAPI Delete Plugin.
// apiNameOrId -- required	The unique identifier or the name of the API for which to delete the plugin configuration
// id -- required	The unique identifier of the plugin configuration to delete on this API
func DeleteACLPluginPerAPI(id string, apiNameOrID string) error {
	//DELETE /apis/{api name or id}/plugins/{id}
	if len(apiNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the API can not be null")
	}
	if len(id) == 0 {
		return errors.New("The unique identifier of the plugin can not be null")
	}

	return DeletePluginPerAPI(id, apiNameOrID)
}

// AssociateGroup  associate a group to a Consumer
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// group -- The arbitrary group name to associate to the consumer.
func AssociateGroup(consumerNameOrID string, group string) (*models.ACLGroup, error) {
	//POST curl -X POST http://kong:8001/consumers/{consumer}/acls \
	// --data "group=group1"
	if len(consumerNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(group) == 0 {
		return nil, errors.New("The the name of the group can not be null")
	}
	req := httplib.Post(kongAdminURL + `/consumers/` + consumerNameOrID + `/acls/`)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	if len(group) > 0 {
		req.Param("group", group)
	}

	var aclGroup models.ACLGroup
	err := req.ToJSON(&aclGroup)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &aclGroup, nil
}

//GetACLGroup Get acl group by id.
func GetACLGroup(consumerNameOrID string, aclGroupid string) (*models.ACLGroup, error) {
	// GET http://13.76.42.81:8001/consumers/guan/acls/a0fdba77-fc6d-4632-845c-cc1a623cf59d
	if len(consumerNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(aclGroupid) == 0 {
		return nil, errors.New("The unique identifier of the aclGroup can not be null")
	}
	req := httplib.Get(kongAdminURL + `/consumers/` + consumerNameOrID + `/acls/` + aclGroupid)
	var retACLGroup models.ACLGroup
	err := req.ToJSON(&retACLGroup)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retACLGroup, nil
}

// ListACLGroup List acl groups of one consumer.
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
func ListACLGroup(consumerNameOrID string, size int, offset string) (*models.ACLGroupList, error) {
	//Get http://kong:8001/consumers/{consumerNameOrId}/acls?size=1 -d '' HTTP/1.1 201 Created
	if len(consumerNameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the consumer can not be null")
	}
	//do get
	u, err := url.Parse(kongAdminURL + `/consumers/` + consumerNameOrID + `/acls`)
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
	var retACLGroupList models.ACLGroupList
	err = req.ToJSON(&retACLGroupList)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retACLGroupList, nil
}

// DeleteACLGroup Delete a acl group by id.
// consumerNameOrId -- The id or username property of the Consumer entity to associate the credentials to.
// aclGroupid -- acl group id
func DeleteACLGroup(consumerNameOrID string, aclGroupid string) error {
	//Delete http://kong:8001/consumers/{consumerNameOrId}/acls/{aclGroupid} -d '' HTTP/1.1 201 Created
	if len(consumerNameOrID) == 0 {
		return errors.New("The unique identifier or the name of the consumer can not be null")
	}
	if len(aclGroupid) == 0 {
		return errors.New("The unique identifier of the aclGroup can not be null")
	}
	req := httplib.Delete(kongAdminURL + `/consumers/` + consumerNameOrID + `/acls/` + aclGroupid)
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
