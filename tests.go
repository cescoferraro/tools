package tools

import (
	"net/http/httptest"
	"testing"

	"github.com/fatih/color"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
	"strings"
	"encoding/json"
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
	logger:= Logger{Title:strings.ToUpper(test.Name),Color:color.FgHiBlue}
	logger("Name: "+test.Name)
	logger("Description: "+test.Description)
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

	return string(actualBody)
}
