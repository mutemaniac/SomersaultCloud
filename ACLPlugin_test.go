package SomersaultCloud

import (
	"SomersaultCloud/models"
	"testing"
)

func TestAddACLPlugin2Api(t *testing.T) {
	//Creat an api for using first.
	var api models.API
	api.Name = "testkeyauth"
	api.PreserveHost = true
	api.RequestHost = "reqhost"
	api.RequestPath = "/test"
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
	var plugin models.Plugin
	apiNameOrID := retAPI.ID
	apiName := retAPI.Name
	consumerID := retConsumer.ID

	retACLPlugin, err := AddACLPlugin2API(apiNameOrID, []string{apiName}, nil)
	if err != nil {
		t.Error("AddACLPlugin Error, ", err)
	}
	if len(retACLPlugin.ID) < 1 {
		t.Error("AddACLPlugin2Api Error, ", *retACLPlugin)
	}

	// test ListACLPlugin
	pluginlist, err := ListACLPlugin(plugin, 5, "")
	if err != nil {
		t.Error("ListACLPlugin Error, ", err)
	}
	if pluginlist.Total < 1 {
		t.Error("ListACLPlugin Error, ", *pluginlist)
	}
	//test GetACLPlugin
	retACLPlugin, err = GetACLPlugin(retACLPlugin.ID)
	if err != nil {
		t.Error("GetACLPlugin Error, ", err)
	}
	if len(retACLPlugin.ID) < 1 {
		t.Error("GetACLPlugin Error, ", *retACLPlugin)
	}

	retACLGroup, err := AssociateGroup(consumerID, apiName)
	if err != nil {
		t.Error("CreateACLGroup Error, ", consumerID, err)
	}
	if len(retACLGroup.ID) < 1 {
		t.Error("CreateACLGroup Error, ", *retACLPlugin)
	}
	// test ListACLPlugin
	retACLGrouplist, err := ListACLGroup(consumerID, 5, "")
	if err != nil {
		t.Error("ListACLGroup Error, ", err)
	}
	if retACLGrouplist.Total < 1 {
		t.Error("ListACLGroup Error, ", *retACLGrouplist)
	}
	//test GetACLPlugin
	retACLGroup, err = GetACLGroup(consumerID, retACLGroup.ID)
	if err != nil {
		t.Error("GetACLGroup Error, ", err)
	}
	if len(retACLGroup.ID) < 1 {
		t.Error("GetACLGroup Error, ", *retACLGroup)
	}

	//test DeleteACLGroup
	err = DeleteACLGroup(consumerID, retACLGroup.ID)
	if err != nil {
		t.Error("DeleteACLGroup Error, ", err)
	}
	retACLGroup, err = GetACLGroup(consumerID, retACLGroup.ID)
	if err == nil {
		t.Error("DeleteACLGroup Error. process failure!")
	}

	//test DeleteACLPlugin
	err = DeleteACLPluginPerAPI(retACLPlugin.ID, apiNameOrID)
	if err != nil {
		t.Error("DeleteACLPlugin Error, ", err)
	}
	retACLPlugin, err = GetACLPlugin(retACLPlugin.ID)
	if err == nil {
		t.Error("DeleteACLPlugin Error. process failure!")
	}

}
