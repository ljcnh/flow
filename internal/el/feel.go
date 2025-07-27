package el

import (
	el "github.com/ljcnh/flow/internal/el/function"
	"github.com/pbinitiative/feel"
)

func RegisterCustomFunction() {
	prelude := feel.GetPrelude()
	if _, ok := prelude.Resolve("format"); ok {
		prelude.Bind("format", el.NewStringFormatFunction())
	}
}
