package integrationtests

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// --------------------------------------------------------------

type tSuite struct {
	suite.Suite
}

func TestEToESuite(t *testing.T) {
	suite.Run(t, new(tSuite))
}

func (suite *tSuite) TestHello() {
	response, err := http.Get("http://localhost:8080/")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Hello World!", string(body))
}

func (suite *tSuite) TestAdd() {
	response, err := http.Get("http://localhost:8080/add/3/2")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "5", string(body))
}

func (suite *tSuite) TestSub() {
	response, err := http.Get("http://localhost:8080/sub/3/2")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "1", string(body))
}

func (suite *tSuite) TestMult() {
	response, err := http.Get("http://localhost:8080/mult/3/2")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "6", string(body))
}

func (suite *tSuite) TestDiv() {
	response, err := http.Get("http://localhost:8080/div/3/2")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "1.5", string(body))
}

func (suite *tSuite) TestDivZero() {
	response, err := http.Get("http://localhost:8080/div/3/0")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "+Inf", string(body))
}

func (suite *tSuite) TestSqrt() {
	response, err := http.Get("http://localhost:8080/sqrt/4")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "2", string(body))
}

func (suite *tSuite) TestCbrt() {
	response, err := http.Get("http://localhost:8080/cbrt/27")
	if assert.Nil(suite.T(), err) {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "3", string(body))
}
