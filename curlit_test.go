package curlit

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestDumpPost(t *testing.T) {
	body := `{"key":"val"}`
	req, err := http.NewRequest(http.MethodPost, `http://www.google.com`, strings.NewReader(body))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	req.Header.Add(`Content-Type`, `application/json`)
	req.Header.Add(`Content-Length`, strconv.Itoa(len(body)))
	s, _ := Dump(req)
	if s != `curl  -H "Content-Type:application/json"  -H "Content-Length:13"  -d "{/"key/":/"val/"}" http://www.google.com` {
		t.Fail()
	}
	t.Log(s)
}

func TestDumpGet(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, `http://www.google.com`, nil)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	req.Header.Add(`Content-Type`, `application/json`)
	s, _ := Dump(req)
	if s != `curl  -H "Content-Type:application/json" http://www.google.com` {
		t.Fail()
	}
	t.Log(s)
}
