package stdlib

// Code generated by 'goexports text/tabwriter'. DO NOT EDIT.

import (
	"reflect"
	"text/tabwriter"
)

func init() {
	Value["text/tabwriter"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AlignRight":          reflect.ValueOf(tabwriter.AlignRight),
		"Debug":               reflect.ValueOf(tabwriter.Debug),
		"DiscardEmptyColumns": reflect.ValueOf(tabwriter.DiscardEmptyColumns),
		"Escape":              reflect.ValueOf(tabwriter.Escape),
		"FilterHTML":          reflect.ValueOf(tabwriter.FilterHTML),
		"NewWriter":           reflect.ValueOf(tabwriter.NewWriter),
		"StripEscape":         reflect.ValueOf(tabwriter.StripEscape),
		"TabIndent":           reflect.ValueOf(tabwriter.TabIndent),

		// type definitions
		"Writer": reflect.ValueOf((*tabwriter.Writer)(nil)),
	}
}
