package typeman

type TypeManager struct {
	typeLst []*RootField
}

func (t *TypeManager) Register(i interface{}, c ...Config) (err error) {
	var conf Config
	if len(c) >= 1 {
		conf = c[0]
	}
	newRootField(i, conf)
	return
}

func (t *TypeManager) RegisterAsync(s string, conf ...Config) (err error) {
	rootField, err := newRootField(i)
	if err != nil {
		return err
	}
	t.typeLst = append(t.typeLst, rootField)
	return nil
}
