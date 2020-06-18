package middleware

import (
	"github.com/namantalaycha/middletest/assert"
	"github.com/namantalaycha/middletest/middleware"
	"github.com/namantalaycha/middletest/mocks"
	"net/http"
	"testing"
)

func TestAuthNil(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	str := "ghrt"
	rr := &mocks.ResponseWriter{}
	req.Header.Add("name", str)
	handler := Auth(middleware.TestHandler)
	handler.ServeHTTP(rr, req)

	if str == "" {
		assert.Equalt(t, rr.StatusCode, http.StatusUnauthorized)
	} else {
		assert.Equalt(t, rr.StatusCode, http.StatusOK)
	}
}


func BenchmarkAuth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, "/books", nil)
		if err != nil {
			b.Fatal(err)
		}
		str := "dgzjfm"
		rr := &mocks.ResponseWriter{}
		req.Header.Add("name", str)

		handler := Auth(middleware.TestHandler)
		handler.ServeHTTP(rr, req)
		if str == "" {
			assert.Equalb(b, rr.StatusCode, http.StatusUnauthorized)
		} else {
			assert.Equalb(b, rr.StatusCode, http.StatusOK)
		}
		assert.Equalb(b, "naman", "naman")
	}
}
