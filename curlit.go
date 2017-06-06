/*Package curlit dump request as curl command

Usually use in errors cases where the request is dump to some persistance (GCS, S3) and later execute again

Example: See tests
*/
package curlit

import (
	"bytes"
	"io"
	"net/http"
)

const newLine = " \\\n"

//Dump create curl command string from http.Request
func Dump(req *http.Request) (string, error) {
	curl := bytes.NewBuffer([]byte(`curl -X `))
	curl.WriteString(req.Method)
	curl.WriteString(newLine)
	curl.WriteString(req.URL.String())
	curl.WriteString(newLine)

	if req.Method == http.MethodPost {
		w := bytes.NewBuffer(nil)
		io.Copy(w, req.Body)

		if w.Len() > 0 {
			curl.WriteString(` -d '`)
			w.WriteTo(curl)
			curl.WriteString(`'`)
			curl.WriteString(newLine)
		}
	}

	headers := bytes.NewBuffer([]byte{})
	if len(req.Header) > 0 {
		headers = bytes.NewBuffer([]byte{})
		for k, v := range req.Header {
			headers.WriteString(` -H '`)
			headers.WriteString(k)
			headers.WriteRune(':')
			for i := range v {
				headers.WriteString(v[i])
				if i+1 < len(v) {
					headers.WriteRune(',')
				}
			}

			headers.WriteString(`'`)
			headers.WriteString(newLine)
		}
	}
	headers.WriteTo(curl)
	b := curl.Bytes()
	ii := len(b) - 1
	for ; ii > 0; ii-- {
		if b[ii] == 92 {
			break
		}
	}
	s := string(b[:ii-1])
	return s, nil
}
