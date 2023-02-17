package json

import (
	"io"

	"github.com/timandy/go-auth0/codec/api"
)

var Api api.JsonApi

func Marshal(v any) ([]byte, error) {
	return Api.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return Api.Unmarshal(data, v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return Api.MarshalIndent(v, prefix, indent)
}

func NewEncoder(writer io.Writer) api.JsonEncoder {
	return Api.NewEncoder(writer)
}

func NewDecoder(reader io.Reader) api.JsonDecoder {
	return Api.NewDecoder(reader)
}
