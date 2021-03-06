package pagination

import (
	safeerr "github.com/leaq-ru/lib-safeerr"
	"github.com/leaq-ru/proto/codegen/go/opts"
)

func ApplyDefaultLimit(req opter) (limit uint32, err error) {
	limit = 20
	if req.GetOpts() != nil {
		if req.GetOpts().GetLimit() > 100 {
			err = safeerr.LimitOutOfRange
			return
		}

		if req.GetOpts().GetLimit() != 0 {
			limit = req.GetOpts().GetLimit()
		}
	}
	return
}

type opter interface {
	GetOpts() *opts.SkipLimit
}
