package tester

import (
	"net/http/httptest"
	"testing"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
	"strings"
	"encoding/json"
	"github.com/cescoferraro/tools/logger"
	"github.com/cescoferraro/api/iot/types"
	"reflect"
)

//Server is the global server var
var TestServer *httptest.Server

//TableTest is an object
type TableTest struct {
	Method       string
	Path         string
	Jwt          string
	Body         interface{}
	BodyContains string
	Status       int
	Name         string
	Description  string
}

func (test TableTest) Spin(t *testing.T) string {
	return string(test.innnerSpin(t))
}

func (test TableTest) Device(t *testing.T) types.Device {
	actualBody:= test.innnerSpin(t)
	var receivedev types.Device
	assert.NoError(t, json.Unmarshal(actualBody,receivedev))
	return receivedev
}

func(test TableTest) innnerSpin(t *testing.T) []byte{
	logger:= logger.New(strings.ToUpper(test.Name))
	logger.Print("Name: "+test.Name)
	logger.Print("Description: "+test.Description)
	url := TestServer.URL + test.Path
	b, err := json.Marshal(test.Body)
	assert.NoError(t, err)
	r, err := http.NewRequest(test.Method, url, strings.NewReader(string(b)))
	assert.NoError(t, err)
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", test.Jwt))
	response, err := http.DefaultClient.Do(r)
	assert.NoError(t, err)
	actualBody, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(actualBody), test.BodyContains, "body")
	assert.Equal(t, test.Status, response.StatusCode, "status code")
	return actualBody
}



func (test TableTest) DoubleSpin(t *testing.T) interface{} {
	actualBody := test.innnerSpin(t)
	thetype := reflect.TypeOf(test.Body)
	receivedev := reflect.New(thetype)
	err := json.Unmarshal(actualBody, receivedev.Interface())
	assert.NoError(t, err)
	return receivedev.Interface()
}
