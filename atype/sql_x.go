package atype

import "context"

type TelX string // varchar(13)

func (x TelX) Decode(ctx context.Context, desense bool, decoder func(context.Context, TelX, bool) (string, bool)) string {
	tel, _ := decoder(ctx, x, desense)
	return tel
}
