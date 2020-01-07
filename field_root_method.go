package typeman

import "reflect"

func (d *RootField) getFieldByName(name string) (o StructFieldLst) {
	return d.fieldNameMap[name]
}

func (d *RootField) getFieldByMark(tag string) (o *StructField) {
	return d.fieldMarkMap[tag]
}

func (d *RootField) getExtendFields() (lst StructFieldLst) {
	for _, v := range d.fields {
		if v.isAnonymous {
			lst = append(lst, v)
		}
	}
	return
}

func (d *RootField) GetRootFields() StructFieldLst {
	return d.rootFields
}

// get fields from root value
/// ***** should optimize, if complex nesting
func (d *RootField) GetRootValues(rootValue *reflect.Value) (result []reflect.Value) {
	for _, v := range d.rootFields {
		var value = *rootValue
		for _, _v := range v.dependLst {
			value = value.FieldByName(_v.name)
		}
		result = append(result, value)
	}
	return
}
