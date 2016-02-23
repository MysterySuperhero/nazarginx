package utils

import (
	"strings"
	"net/url"
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Body     string
	Headers  Headers
}

func (r *Request) getGeneral(field string)  {
	field_strings := strings.Split(field, " ")
	r.Method = field_strings[0]
	uri := strings.Split(field_strings[1], "?")
	r.Path, _ = url.QueryUnescape(uri[0])
	r.Protocol = field_strings[2]
}

func parseHeaders(fields []string) (Headers) {
	headers := Headers{}
	for _, el := range fields {
		header := strings.Split(el, Separators["header"])
		if len(header) == 2 {
			headers.Add(header[0], header[1])
		}
	}

	return headers
}

func ParseRequest(req string) (*Request) {
	request := new(Request)

	request_strings := strings.Split(req, Separators["field"])

	request.getGeneral(request_strings[0])

	request.Headers = parseHeaders(request_strings[1:])

	return request
}
