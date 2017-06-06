package curlit

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestDumpPostOneHeader(t *testing.T) {
	body := `{"key":"val"}`
	req, err := http.NewRequest(http.MethodPost, `http://www.google.com`, strings.NewReader(body))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	req.Header.Add(`Content-Type`, `application/json`)
	req.Header.Add(`Content-Length`, strconv.Itoa(len(body)))
	s, _ := Dump(req)
	if s != `curl -X POST \
http://www.google.com \
 -d '{"key":"val"}' \
 -H 'Content-Type:application/json' \
 -H 'Content-Length:13'` {
		t.Fail()
	}
}

func TestDumpPostMultiHeader(t *testing.T) {
	body := `{"key":"val"}`
	req, err := http.NewRequest(http.MethodPost, `http://www.google.com`, strings.NewReader(body))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	req.Header.Add(`Content-Type`, `application/json`)
	req.Header.Add(`Content-Type`, `gzip`)
	req.Header.Add(`Content-Length`, strconv.Itoa(len(body)))
	s, _ := Dump(req)

	if s != `curl -X POST \
http://www.google.com \
 -d '{"key":"val"}' \
 -H 'Content-Type:application/json,gzip' \
 -H 'Content-Length:13'` {
		t.Fail()
	}
}

func TestDumpGet(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, `http://www.google.com`, nil)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	req.Header.Add(`Content-Type`, `application/json`)
	s, _ := Dump(req)
	if s != `curl -X GET \
http://www.google.com \
 -H 'Content-Type:application/json'` {
		t.Fail()
	}
}
