package utils

import (
	"path/filepath"
	"fmt"
	"os"
	"regexp"
	"io/ioutil"
	"time"
	"strconv"
	"strings"
	"bytes"
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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func isFolder(path string) (bool) {
	if _, err := regexp.MatchString(".*/", path); err == nil {
		return true
	}
	return false
}

func (response *Response) setGeneral(method string, path *string, doc_root string) {
	folder := false
	if isFolder(*path) {
		fmt.Println("It's a folder!")
		*path += DEFAULT_FILE
		fmt.Println("So, I'll give you: " + *path)
		folder = true
	}

	file, err := ioutil.ReadFile(*path)

	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			if folder {
				fmt.Println("Forbidden!")
				response.status = Forbidden
			} else {
				fmt.Println("File not found!")
				response.status = NotFound
			}
		}
		return
	}

	if method == "GET" {
		fmt.Println("Adding file to body...")
		response.body = file
	}
	response.status = OK
}

func contentTypeFromPath(path string) string  {
	return Content_Types[strings.Trim(path, ".")]
}

func (response *Response) setHeaders(path string)  {
	response.headers = Headers{}
	fmt.Println("Setting headers:")
	fmt.Println("Date" + time.Now().String())
	response.headers.Add("Date", time.Now().String())
	fmt.Println("Server" + "nazarginx v0.1")
	response.headers.Add("Server", "nazarginx v0.1")
	fmt.Println("Content-Length" + strconv.Itoa(len(response.body)))
	response.headers.Add("Content-Length", strconv.Itoa(len(response.body)))
	fmt.Println("Content-Type" + contentTypeFromPath(path))
	response.headers.Add("Content-Type", contentTypeFromPath(path))
	fmt.Println("Connection" + "close")
	response.headers.Add("Connection", "close")
}

func (response *Response) Byte() []byte  {
	var result bytes.Buffer

	result.WriteString(response.protocol + " " + response.status + Separators["field"])
	result.WriteString(response.headers.String() + Separators["field"])
	result.WriteString(string(response.body))

	return result.Bytes()
}

func (response *Response) CreateResponse(method string, path string, doc_root string) {

	absPath, _ := filepath.Abs(".." + path)
	fmt.Println("Trying to find at: " + absPath)

	if _, err := exists(absPath); err == nil {
		fmt.Println("WOW, file exists")

		current_dir, _ := os.Getwd()
		path = current_dir + path

		fmt.Println("AZAZA " + path)
		response.setGeneral(method, &path, doc_root)
		response.setHeaders(absPath)
	} else {
		fmt.Println("OOOPS, file doesn't exists")
		response.status = NotFound
	}

	response.protocol = HttpProtocol
}

