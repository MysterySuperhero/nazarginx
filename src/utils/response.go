package utils

import (
	"fmt"
	"os"
	"regexp"
	"io/ioutil"
	"time"
	"strconv"
	"bytes"
	"strings"
)

type Response struct {
	status  	string
	protocol	string
	body		[]byte
	headers 	Headers
}

const (
	DEFAULT_FILE = "index.html"
)

func checkExistence(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func isFolder(path string) (bool) {
	if res, _ := regexp.MatchString(".*/$", path); res {
		return true
	}
	return false
}


// returns false if somebody wants to fool you
func checkPath(path string) bool  {
	if !strings.Contains(path, "../") {
		return true
	} else {
		return false
	}
}

func contentTypeFromPath(path string) string  {
	re := regexp.MustCompile(".*\\.")
	return Content_Types[re.Split(path, -1)[1]]
}

func (response *Response) setDefault()  {
	response.protocol = HttpProtocol
	response.headers = Headers{}
	response.headers.Add("Date", time.Now().String())
	response.headers.Add("Server", "nazarginx v0.1")
	response.headers.Add("Connection", "close")
}

func (response *Response) setGeneral(method string, path *string) {
	folder := false
	if isFolder(*path) {
		*path += DEFAULT_FILE
		folder = true
	}

	file, err := ioutil.ReadFile(*path)

	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			if folder {
				response.status = Forbidden
			} else {
				response.status = NotFound
			}
		}
		return
	}

	if method == "GET" {
		response.body = file
	}
	response.setSuccessHeaders(*path, file)
	response.status = OK
}

func (response *Response) setSuccessHeaders(path string, file []byte)  {
	response.headers.Add("Content-Length", strconv.Itoa(len(file)))
	response.headers.Add("Content-Type", contentTypeFromPath(path))
}

func (response *Response) Byte() []byte  {
	var result bytes.Buffer

	result.WriteString(response.protocol + " " + response.status + Separators["field"])
	result.WriteString(response.headers.String() + Separators["field"])
	result.WriteString(string(response.body))

	return result.Bytes()
}

func (response *Response) CreateResponse(method string, path string, doc_root string) {
	response.setDefault()

	if !Supported_Methods[method] {
		response.status = NotAllowed
		return
	}

	if !checkPath(path) {
		response.status = Forbidden
		return
	}

	current_dir, _ := os.Getwd()
	path = current_dir + path

	existence,err := checkExistence(path)
	if existence &&  err == nil {
		response.setGeneral(method, &path)
	} else {
		response.status = NotFound
	}

}

