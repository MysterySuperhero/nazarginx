package utils

import (
	"strings"
	"net/url"
	"errors"
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Body     string
	Headers  Headers
}

func (r *Request) getGeneral(field string) (error) {
	field_strings := strings.Split(field, " ")
	if (len(field_strings) >= 3) {
		r.Method = field_strings[0]
		uri := strings.Split(field_strings[1], "?")
		r.Path, _ = url.QueryUnescape(uri[0])
		r.Protocol = field_strings[2]
		return nil
	} else {
		return errors.New("Bad request")
	}
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

func ParseRequest(req string) (*Request, error) {
	request := new(Request)

	request_strings := strings.Split(req, Separators["field"])

	err := request.getGeneral(request_strings[0])

	if err != nil {
		return nil, err
	}

	request.Headers = parseHeaders(request_strings[1:])

	return request, nil
}
