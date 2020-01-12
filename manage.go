package typeman

import "errors"

type TypeManager struct {
	typeLst []*RootField
	mapType map[string]*RootField
}

func New() (t *TypeManager) {
	t.mapType = make(map[string]*RootField)
	return
}
func (t *TypeManager) Register(i interface{}, c ...Config) (err error) {
	var conf Config
	if len(c) >= 1 {
		conf = c[0]
	}

	// check Async Type
	str, ok := i.(string)
	if ok {
		conf.isAsync = true
		i = str
	}
	rootField, err := newRootField(i, conf)
	if err != nil {
		return err
	}
	sameNameField := t.GetTypeByName(rootField.Name())
	if sameNameField != nil {
		return errors.New("There has same name field alive, field name : '" + sameNameField.Name() + "'")
	}
	t.typeLst = append(t.typeLst, rootField)
	t.mapType[sameNameField.Name()] = sameNameField
	return
}

func (t *TypeManager) GetTypeByName(s string) *RootField {
	return t.mapType[s]
}
