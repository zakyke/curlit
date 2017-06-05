package curlit

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

//Dump create curl command string from http.Request
func Dump(req *http.Request) (string, error) {
	curl := bytes.NewBuffer([]byte(`curl `))
	headers := bytes.NewBuffer([]byte{})

	if len(req.Header) > 0 {
		headers = bytes.NewBuffer([]byte{})
		for k, v := range req.Header {
			headers.WriteString(` -H "`)
			headers.WriteString(k)
			headers.WriteRune(':')
			for i := range v {
				headers.WriteString(v[i])
				if i+1 < len(v) {
					headers.WriteRune(',')
				}
			}

			headers.WriteString(`" `)

		}
	}
	io.Copy(curl, headers)

	if req.Method == http.MethodPost {
		curl.WriteString(` -d "`)
		w := bytes.NewBuffer([]byte{})
		io.Copy(w, req.Body)
		body := strings.Replace(w.String(), `"`, `/"`, -1)
		curl.WriteString(body)
		curl.WriteString(`" `)
	}
	curl.WriteString(req.URL.String())

	return curl.String(), nil
}
