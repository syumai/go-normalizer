package normalizer

import (
	"fmt"
	"reflect"
	"strings"
)

type Normalizer interface {
	Normalize(interface{}) error
}

type normalizer struct {
	funcMap FuncMap
}

func New() Normalizer {
	return &normalizer{
		funcMap: NewFuncMap(),
	}
}

func (n *normalizer) Normalize(i interface{}) error {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("normalizer can use only pointer type. can't use: %s", t.Kind().String())
	}

	v := reflect.ValueOf(i).Elem()
	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		if valueField.Kind() != reflect.String {
			continue
		}

		typeField := v.Type().Field(i)
		tags := strings.Split(typeField.Tag.Get("normalize"), ",")
		for _, tag := range tags {
			f, ok := n.funcMap.Get(tag)
			if !ok {
				return fmt.Errorf("no such normalizer func: %s", tag)
			}
			valueField.SetString(f(valueField.String()))
		}
	}
	return nil
}
