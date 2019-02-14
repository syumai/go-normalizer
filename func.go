package normalizer

import (
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

type NormalizerFunc func(string) string

type FuncMap map[string]NormalizerFunc

func NewFuncMap() FuncMap {
	return map[string]NormalizerFunc{
		"lower":   strings.ToLower,
		"upper":   strings.ToUpper,
		"capital": strings.Title,
		"title":   strings.Title,
		"snake":   strcase.ToSnake,
		"camel":   strcase.ToLowerCamel,
		"pascal":  strcase.ToCamel,
		"kebab":   strcase.ToKebab,
		"widen":   widen,
		"narrow":  narrow,
	}
}

func widen(s string) string {
	return norm.NFC.String(width.Widen.String(s))
}

func narrow(s string) string {
	return width.Narrow.String(s)
}

func (m FuncMap) Get(name string) (NormalizerFunc, bool) {
	f, ok := m[name]
	if !ok {
		return nil, false
	}
	return f, true
}

func (m FuncMap) Set(name string, f NormalizerFunc) {
	m[name] = f
}
