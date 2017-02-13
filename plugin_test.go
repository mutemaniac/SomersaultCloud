package kong

import (
	"service-cloud/utils/kong/models"
	"testing"
)

func TestAddPlugin(t *testing.T) {
	//Creat an api for using first.
	var api models.API
	api.Name = "testPlugin"
	api.PreserveHost = true
	api.RequestHost = "baidu.com"
	api.RequestPath = "/testPlugin"
	api.StripRequestPath = false
	api.UpstreamURL = "http://www.baidu.com/"
	retAPI, _ := AddAPI(&api)
	//Creat a consumer for using next.
	var consumer models.Consumer
	consumer.Username = "testkeyauth"
	retConsumer, _ := AddConsumer(&consumer)
	defer func() {
		//test DeleteConsumer
		err := DeleteConsumer(retConsumer.ID)
		if err != nil {
			t.Error("DeleteConsumer Error, ", err)
		}
		//DeleteAPI
		err = DeleteAPI(retAPI.ID)
		if err != nil {
			t.Error("DeleteAPI Error, ", err)
		}
	}()

	//----------------------------key-auth test below-----------------------------------

	//test add plugin
	//var plugin models.Plugin
	apiNameOrID := retAPI.ID
	//consumerId := retConsumer.ID

	//var ret models.KeyAuthConfig
	var keyconfig models.KeyAuthConfig
	var retKeyAuthPlugin models.KeyAuthPlugin

	keyconfig.KeyNames = []string{"apikey", "test-key"}
	err := AddPlugin("key-auth", apiNameOrID, "", keyconfig, &retKeyAuthPlugin)
	if err != nil {
		t.Error("AddPlugin Error, ", err)
	}
	if len(retKeyAuthPlugin.ID) < 1 {
		t.Error("AddPlugin2Api Error, ", retKeyAuthPlugin)
	}

	//test ListPlugin
	var plugin models.Plugin
	var pluginList models.KeyAuthPluginList
	err = ListPlugin(plugin, 5, "", &pluginList)
	if err != nil {
		t.Error("ListPlugin Error, ", err)
	}
	if pluginList.Total < 1 {
		t.Error("ListPlugin Error, ", pluginList)
	}
	//test GetPlugin
	err = GetPlugin(retKeyAuthPlugin.ID, &retKeyAuthPlugin)
	if err != nil {
		t.Error("GetPlugin Error, ", err)
	}
	if len(retKeyAuthPlugin.ID) < 1 {
		t.Error("GetPlugin Error, ", retKeyAuthPlugin)
	}

	//test update
	var uKeyconfig models.KeyAuthConfig
	uKeyconfig.KeyNames = []string{"key-test", "key-test2"}
	var uRetKeyAuthPlugin models.KeyAuthPlugin
	err = UpdatePlugin(apiNameOrID, retKeyAuthPlugin.ID, &uKeyconfig, &uRetKeyAuthPlugin)
	if err != nil {
		t.Error("UpdatePlugin Error, ", err)
	}
	if len(uRetKeyAuthPlugin.Config.KeyNames) < 1 ||
		uRetKeyAuthPlugin.Config.KeyNames[0] != "key-test" {
		t.Error("UpdatePlugin failure, u_retKeyAuthPlugin=", uRetKeyAuthPlugin)
	}

	//test DeletePlugin
	err = DeletePluginPerAPI(retKeyAuthPlugin.ID, apiNameOrID)
	if err != nil {
		t.Error("DeletePlugin Error, ", err)
	}
	var retPlugin models.KeyAuthPlugin
	err = GetPlugin(retKeyAuthPlugin.ID, retPlugin)
	if err == nil || len(retPlugin.ID) > 0 {
		t.Error("DeletePlugin Error. process failure!", retPlugin)
	}

}
