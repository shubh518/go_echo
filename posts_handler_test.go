package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

// here we need to pass e the EndTOend suite by using the function  of end to endsuite t *testing.T and run the tests

func TestEndToEndSuite(t *testing.T) {

	suite.Run(t, new(EndToEndSuite))

}

// here write and run basic Test

func (s *EndToEndSuite) PostHandler() {

	// here we created a client and to test if server works
	//you have to create a client that will test the server is w9orking properly or not

	c := http.Client{}
	r, _ := c.Get("http://localhost:1234/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}
func (s *EndToEndSuite) TestPostNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1234/posts/142345453")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status": "ok","data:[]}`, string(b))

}
