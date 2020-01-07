package typema

import (
	"fmt"
	"reflect"
	"strings"
)

// this vars could be user modify, so use var not const
var (
	tagName          = "xodm"
	tagItemSeparator = ":"  // Ex: xodm:"name:value"
	tagSeparator     = " "  // Ex: xodm:"name1:value1 name2:value2"
	tagMark          = "@"  // Ex: xodm:"@mark"
	tagFunc          = "()" // Ex: xodm:"default:now()"
	tagDefault       = "default"
)

// `xodm:"@sign"`
type odmTag struct {
	sourceTag string
	mark      string // find docField quick by custom string
	notNull   bool
	// if value not null use default value, defalut value get from string, so must change type and cover in reflect.Value
	defaultValue func() interface{}
	lst          map[string]string
}

func newTag(s string, Field *Field) *odmTag {
	_o := &odmTag{}
	_o.lst = make(map[string]string)
	_s := strings.TrimSpace(s)
	_o.sourceTag = _s
	lst := strings.Split(_s, tagSeparator)
	for _, v := range lst {
		if v == "" {
			continue
		}
		FieldLst := strings.Split(v, tagItemSeparator)
		FieldLstLen := len(FieldLst)
		var name string
		var value string
		// tagMark value like @xxx, so format it diffrent with other
		if FieldLstLen == 1 {
			var _str = FieldLst[0]
			switch {
			// notnull
			case _str == "notnull":
				_o.notNull = true
				// @xxx
			case strings.Index(_str, tagMark) == 0:
				_o.mark = string([]rune(FieldLst[0])[1:])
			}
			continue
		}
		name = strings.TrimSpace(FieldLst[0])
		value = strings.TrimSpace(FieldLst[1])
		if name == tagDefault {
			var FuncIndex = strings.LastIndexAny(value, tagFunc)
			var isFunc = FuncIndex == len(value)-len(tagFunc)
			if isFunc {
				funcName := value[:len(value)-len(tagFunc)]
				_o.defaultValue = customTypeBox.defaultFuncMap[funcName]
			} else {
				_o.defaultValue = func() interface{} {
					newValue := reflect.New(Field.sourceType)
					_ = Field.Parse([]byte(value), newValue)
					return newValue.Interface()
				}
			}
			continue
		}
		_o.lst[name] = value
	}
	return _o
}

func (t *odmTag) NotNull() bool {
	return t.notNull
}

func (t *odmTag) Lst() map[string]string {
	fmt.Print(t.lst)
	return t.lst
}

func (t *odmTag) DefaultValue() interface{} {
	return t.defaultValue()
}

func (t *odmTag) HasDefault() bool {
	return t.defaultValue != nil
}
