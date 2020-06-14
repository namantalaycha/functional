package test

import (
	"encoding/json"
	"fmt"
	"github.com/namantalaycha/middletest/jsonschema"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type Routes struct {
	RouteNodes []Route `json:"routes"`
}
type Route struct {
	Method  string   `json:"method"`
	Path    string   `json:"path"`
	Headers []Header `json:"headers"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestApp(t *testing.T) {

	schema := "file:///home/naman/go/src/github.com/namantalaycha/functional/test/schema.json"
	input := "file:///home/naman/go/src/github.com/namantalaycha/functional/test/input.json"

	if !jsonschema.Validate(input, schema) {
		t.FailNow()
	}

	jsonFile, err := os.Open("input.json")
	if err != nil {
		fmt.Println(err)
	}
	file, _ := ioutil.ReadAll(jsonFile)
	var routes Routes
	if err := json.Unmarshal(file, &routes); err != nil {
		fmt.Println(err)
	}

	for _, route := range routes.RouteNodes {
		req, err := http.NewRequest(route.Method, "http://localhost:3000"+route.Path, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, header := range route.Headers {
			req.Header.Add(header.Key, header.Value)
		}
		res, err := http.DefaultClient.Do(req)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.EqualValues(t, res.StatusCode, http.StatusOK, route.Path)
	}
}

func BenchmarkApp(b *testing.B) {
	for i := 0; i < b.N; i++ {

		schema := "file:///home/naman/go/src/github.com/namantalaycha/functional/test/schema.json"
		input := "file:///home/naman/go/src/github.com/namantalaycha/functional/test/input.json"

		if !jsonschema.Validate(input, schema) {
			b.FailNow()
		}

		jsonFile, err := os.Open("input.json")
		if err != nil {
			fmt.Println(err)
		}
		file, _ := ioutil.ReadAll(jsonFile)
		var routes Routes
		if err := json.Unmarshal(file, &routes); err != nil {
			fmt.Println(err)
		}

		for _, route := range routes.RouteNodes {
			req, err := http.NewRequest(route.Method, "http://localhost:3000"+route.Path, nil)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, header := range route.Headers {
				req.Header.Add(header.Key, header.Value)
			}
			res, err := http.DefaultClient.Do(req)
			assert.Nil(b, err)
			assert.NotNil(b, res)
			assert.EqualValues(b, res.StatusCode, http.StatusOK, route.Path)

		}
	}
}
