package atype

import "context"

type Cipher string // varchar(13)

func (x Cipher) Decode(ctx context.Context, desense bool, decoder func(context.Context, Cipher, bool) (string, bool)) string {
	tel, _ := decoder(ctx, x, desense)
	return tel
}
