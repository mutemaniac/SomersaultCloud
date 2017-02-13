package kong

import (
	"service-cloud/utils/kong/models"
	"testing"
)

func TestAddApi(t *testing.T) {
	//test add api
	var api models.API
	api.Name = "test"
	api.PreserveHost = true
	api.RequestHost = "reqhost"
	api.RequestPath = "/test"
	api.StripRequestPath = false
	api.UpstreamURL = "http://www.baidu.com/"

	retAPI, err := AddAPI(&api)
	if err != nil {
		t.Error("AddAPI Error, ", err)
	}
	if retAPI.Name != api.Name {
		t.Errorf("AddAPI error: %s.", retAPI.Name)
		t.Error("AddAPI Error, ", *retAPI)
	}
	// test ListAPIs
	apilist, err := ListAPIs(5, "")
	if err != nil {
		t.Error("ListAPIs Error, ", err)
	}
	if apilist.Total < 1 {
		t.Error("ListAPIs Error, ", *apilist)
	}
	//test GetAPI
	retAPI, err = GetAPI(api.Name)
	if err != nil {
		t.Error("GetAPI Error, ", err)
	}
	if retAPI.Name != api.Name {
		t.Error("GetAPI Error, ", *retAPI)
	}
	//test UpdateAPI
	api.Name = "test1"
	retAPI, err = UpdateAPI(retAPI.Name, &api)
	if err != nil {
		t.Error("UpdateAPI Error, ", err)
	}
	if retAPI.Name != "test1" {
		t.Error("UpdateAPI Error, ", *retAPI)
	}

	//test DeleteAPI
	err = DeleteAPI(retAPI.ID)
	if err != nil {
		t.Error("DeleteAPI Error, ", err)
	}
	retAPI, err = GetAPI(retAPI.ID)
	if err == nil {
		t.Error("DeleteAPI Error. process failure!")
	}
}

// func TestClearAll(t *testing.T) {
// 	clearPlugins(t, "")
// 	clearConsumer(t, "")
// 	clearApis(t, "")
// }

// func clearPlugins(t *testing.T, offset string) {
// 	var plugin models.Plugin
// 	var pluginList models.PluginList
// 	ListPlugin(plugin, 1000, offset, pluginList)
// 	for _, p := range pluginList.Data {
// 		DeletePluginPerApi(p.ID, p.ApiId)
// 	}
// 	if pluginList.Total > len(pluginList.Data) {
// 		clearPlugins(t, pluginList.Offset)
// 	}
// }
// func clearConsumer(t *testing.T, offset string) {
// 	consumerList, _ := ListConsumers(1000, offset)
// 	for _, c := range consumerList.Data {
// 		DeleteConsumer(c.ID)
// 	}
// 	if consumerList.Total > len(consumerList.Data) {
// 		clearConsumer(t, consumerList.Offset)
// 	}
// }
// func clearApis(t *testing.T, offset string) {
// 	apiList, _ := ListAPIs(1000, offset)
// 	for _, a := range apiList.Data {
// 		DeleteAPI(a.ID)
// 	}
// 	if apiList.Total > len(apiList.Data) {
// 		clearApis(t, apiList.Offset)
// 	}
// }
