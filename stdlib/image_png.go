package stdlib

// Code generated by 'goexports image/png'. DO NOT EDIT.

import (
	"image/png"
	"reflect"
)

func init() {
	Value["image/png"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BestCompression":    reflect.ValueOf(png.BestCompression),
		"BestSpeed":          reflect.ValueOf(png.BestSpeed),
		"Decode":             reflect.ValueOf(png.Decode),
		"DecodeConfig":       reflect.ValueOf(png.DecodeConfig),
		"DefaultCompression": reflect.ValueOf(png.DefaultCompression),
		"Encode":             reflect.ValueOf(png.Encode),
		"NoCompression":      reflect.ValueOf(png.NoCompression),

		// type definitions
		"CompressionLevel":  reflect.ValueOf((*png.CompressionLevel)(nil)),
		"Encoder":           reflect.ValueOf((*png.Encoder)(nil)),
		"EncoderBuffer":     reflect.ValueOf((*png.EncoderBuffer)(nil)),
		"EncoderBufferPool": reflect.ValueOf((*png.EncoderBufferPool)(nil)),
		"FormatError":       reflect.ValueOf((*png.FormatError)(nil)),
		"UnsupportedError":  reflect.ValueOf((*png.UnsupportedError)(nil)),
	}
}
