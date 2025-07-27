package el

import (
	"github.com/pbinitiative/feel"
)

type StringFormatArgs struct {
	Input string            `json:"input"`
	Args  map[string]string `json:"args"`
}

func format(kwargs map[string]any) (any, error) {
	// formatArgs := StringFormatArgs{}
	return nil, nil
}

func NewStringFormatFunction() *feel.NativeFun {
	return feel.NewNativeFunc(format).Required("input", "args")
}
