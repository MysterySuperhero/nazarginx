package utils

type Headers map[string]string

func (headers Headers) Add(key string, value string)  {
	headers[key] = value
}

func (headers Headers) Remove(key string) {
	delete(headers, key)
}

func (headers Headers) Get(key string) (string) {
	return headers[key]
}

func (headers Headers) String() string  {
	result  := ""

	for key, value := range headers {
		result += key + ": " + value + Separators["field"]
	}

	return result
}