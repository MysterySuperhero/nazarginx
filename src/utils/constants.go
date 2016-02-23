package utils

var (
	Separators = map[string]string {
		"header": ": ",
		"field": "\r\n",
		"azaza": "\r\n\r\n",
	}

	 Content_Types = map[string]string {
		 "css": "text/css",
		 "gif": "image/gif",
		 "html": "text/html",
		 "jpeg": "image/jpeg",
		 "jpg": "image/jpeg",
		 "js": "text/javascript",
		 "json": "application/json",
		 "txt": "application/text",
		 "png": "image/png",
		 "swf": "application/x-shockwave-flash",
	}

	Supported_Methods = map[string]bool {
		"GET": true,
		"HEAD": true,
		"CONNECT": false,
		"DELETE": false,
		"OPTIONS": false,
		"PATCH": false,
		"POST": false,
		"PUT": false,
		"TRACE": false,
	}

)

const (
	OK = "200 OK"
	Forbidden = "403 Forbidden"
	NotFound = "404 Not Found"
	NotAllowed = "405 Method Not Allowed"
	HttpProtocol = "HTTP/1.1"
)

