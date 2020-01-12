package typeman

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type Config struct {
	isAsync bool
	IsValid func(i interface{}) bool
	Doc     string
	sign    map[interface{}]interface{}
}

type RootField struct {
	path string
	StructField
	isAsyncType     bool // is format type from String
	interfaceFields map[string]*StructField
	FieldMarkMap    map[string]*StructField
	FieldNameMap    map[string]StructFieldLst
}
type RootFieldLst []*RootField

func newRootField(i interface{}, conf Config) (_RootField *RootField, err error) {

	// append RootField.Fields
	_RootFieldSourceT := reflect.TypeOf(i)
	isAsyncType := _RootFieldSourceT.Kind()
	RootFieldSourceT := _RootFieldSourceT.Elem()
	_RootField = &StructField{
		name:       RootFieldSourceT.Name(),
		sourceType: &RootFieldSourceT,
	}
	Fields := newFieldLst(_RootField, RootFieldSourceT)
	_RootField.Fields = *Fields
isAsyncType:
	isAsyncType
	_RootField.FieldMarkMap = makeFieldLstMarkMap(_RootField)
	_RootField.FieldNameMap = makeFieldLstNameMap(_RootField)
	_RootField.RootFields = makeRootFieldNameMap(_RootField)
	_RootField.interfaceFields = makeInterfaceFields(_RootField)
	return
}

func newFieldLst(d *RootField, RootFieldSourceT reflect.Type) *FieldLst {
	var lst FieldLst
	if RootFieldSourceT.Kind() == reflect.Struct {
		cont := RootFieldSourceT.NumField()
		for i := 0; i < cont; i++ {
			Field := RootFieldSourceT.Field(i)
			// addFieldsLock.Add(1)
			// go newFieldLst(d, &lst, &Field, nil)
			newField(d, &lst, &Field, nil)
		}
		// check Fields Name, Can't both same name in one Col
		// RootField.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("RootField type is "+RootFieldSourceT.Kind().String()+"!,Type should be Struct"))
	}
	// addFieldsLock.Wait()
	return &lst
}

func makeFieldLstMarkMap(d *RootField) (m map[string]*Field) {
	_d := d.Fields
	for _, v := range _d {
		tagPtr := v.odmTag.mark
		if tagPtr != "" {
			m[tagPtr] = v
		}
	}
	return m
}

func makeFieldLstNameMap(d *RootField) map[string]FieldLst {
	_d := d.Fields
	var _map = make(map[string]FieldLst)
	for _, v := range _d {
		name := v.Name()
		// new m[name]
		if _, ok := _map[name]; !ok {
			var temp FieldLst
			_map[name] = temp
		}
		_map[name] = append(_map[name], v)
	}
	return _map
}

func makeRootFieldNameMap(d *RootField) (lst []*Field) {
	_d := d.Fields
	for _, v := range _d {
		if v.logicParent == nil && v.isAnonymous() == false {
			lst = append(lst, v)
		}
	}
	return
}

func makeInterfaceFields(d *RootField) (lst map[string]*Field) {
	for _, v := range d.Fields {
		if _, ok := lst[v.name]; !ok {
			if v.Kind() == Interface {
				lst[v.name] = v
			}
		}

	}
	return
}
