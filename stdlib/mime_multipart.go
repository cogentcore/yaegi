package stdlib

// Code generated by 'goexports mime/multipart'. DO NOT EDIT.

import (
	"mime/multipart"
	"reflect"
)

func init() {
	Value["mime/multipart"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrMessageTooLarge": reflect.ValueOf(&multipart.ErrMessageTooLarge).Elem(),
		"NewReader":          reflect.ValueOf(multipart.NewReader),
		"NewWriter":          reflect.ValueOf(multipart.NewWriter),

		// type definitions
		"File":       reflect.ValueOf((*multipart.File)(nil)),
		"FileHeader": reflect.ValueOf((*multipart.FileHeader)(nil)),
		"Form":       reflect.ValueOf((*multipart.Form)(nil)),
		"Part":       reflect.ValueOf((*multipart.Part)(nil)),
		"Reader":     reflect.ValueOf((*multipart.Reader)(nil)),
		"Writer":     reflect.ValueOf((*multipart.Writer)(nil)),
	}
}
