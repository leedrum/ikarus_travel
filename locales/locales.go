package locales

import (
	"context"
	"embed"

	"github.com/invopop/ctxi18n"
	"github.com/invopop/ctxi18n/i18n"
)

//go:embed en

var Content embed.FS

func Translate(ctx context.Context, key string, args ...any) string {
	if args == nil {
		return ctxi18n.Get(i18n.Code(ctx.Value("lang").(string))).T(key)
	}

	return ctxi18n.Get(i18n.Code(ctx.Value("lang").(string))).T(key, args)
}
