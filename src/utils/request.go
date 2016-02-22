package utils

import (
	"strings"
)

type Request struct {
	method  	string
	path    	string
	protocol   	string
	body    	string
	headers 	Headers
}

func (r *Request) getMethodPathAndProtocol(field string)  {
	field_strings := strings.Split(field, " ")
	r.method = field_strings[0]
	r.path = field_strings[1]
	r.protocol = field_strings[2]
}

func parseHeaders(fields []string) (Headers) {
	headers := Headers{}
	for _, el := range fields {
		header := strings.Split(el, separators["header"])
		if len(header) == 2 {
			headers.Add(header[0], header[1])
		}
	}

	return headers
}

func ParseRequest(req string) (*Request) {
	request := new(Request)

	request_strings := strings.Split(req, separators["field"])

	request.getMethodPathAndProtocol(request_strings[0])

	request.headers = parseHeaders(request_strings[1:])

	return request
}
