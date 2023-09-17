package atype

import "context"

type TelX string // varchar(13)

func (x TelX) Decode(ctx context.Context, desense bool, decoder func(context.Context, string, bool) (string, bool)) string {
	tel, _ := decoder(ctx, string(x), desense)
	return tel
}
