package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"service-cloud/utils/kong/models"
	"strconv"

	"github.com/astaxie/beego/httplib"
)

// AddConsumer Create Consumer
// consumer -- You must send either username or custom_id in the consumer.
func AddConsumer(consumer *models.Consumer) (*models.Consumer, error) {
	// POST /consumers/
	req := httplib.Post(kongAdminURL + `/consumers/`)
	if len(consumer.Username) > 0 {
		req.Param("username", consumer.Username)
	}
	if len(consumer.CustomID) > 0 {
		req.Param("custom_id", consumer.CustomID)
	}

	var retConsumer models.Consumer
	err := req.ToJSON(&retConsumer)

	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retConsumer, nil
}

//GetConsumer Retrieve Consumer
//nameOrID -- The unique identifier or the username of the consumer to retrieve
func GetConsumer(nameOrID string) (*models.Consumer, error) {
	//GET /consumers/{username or id}
	if len(nameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the Consumer can not be null")
	}
	req := httplib.Get(kongAdminURL + `/consumers/` + nameOrID)

	var retConsumer models.Consumer
	err := req.ToJSON(&retConsumer)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retConsumer, nil
}

//ListConsumers List Consumers
func ListConsumers(size int, offset string) (*models.ConsumerList, error) {
	// GET /consumers/
	u, err := url.Parse(kongAdminURL + `/consumers/`)
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
	var retConsumerList models.ConsumerList
	err = req.ToJSON(&retConsumerList)
	if err != nil {
		return nil, err
	}
	resp, _ := req.Response()
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		retStr, _ := req.String()
		return nil, errors.New(retStr)
	}
	return &retConsumerList, nil
}

// UpdateConsumer Update Consumer
// usernameOrID -- (required)The unique identifier or the username of the consumer to update
// consumer -- new information
func UpdateConsumer(usernameOrID string, consumer *models.Consumer) (*models.Consumer, error) {
	//PATCH /consumers/{name or id}
	if len(usernameOrID) == 0 {
		return nil, errors.New("The unique identifier or the name of the Consumer can not be null")
	}
	jsonStr, err := json.Marshal(consumer)
	url := kongAdminURL + `/consumers/` + usernameOrID
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var retConsumer models.Consumer
	err = json.NewDecoder(resp.Body).Decode(&retConsumer)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return &retConsumer, errors.New("UpdateAPI error, " + resp.Status)
	}

	return &retConsumer, nil
}

//DeleteConsumer Delete Consumer
func DeleteConsumer(nameOrID string) error {
	//DELETE /consumers/{username or id}
	if len(nameOrID) == 0 {
		return errors.New("The unique identifier or the name of the Consumer can not be null")
	}
	req := httplib.Delete(kongAdminURL + `/consumers/` + nameOrID)

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
