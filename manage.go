package typeman

type Config struct {
	IsValid func(i interface{}) bool
}
type TypeManager struct {
	typeLst []*RootField
}

func (t *TypeManager) Register(i interface{}, ...conf Config) (err error) {
	newRootField(i)
	return
}

func (t *TypeManager) RegisterAsync(s string, ...conf Config) (err error) {
	rootField, err := newRootField(i)
	if err != nil {
		return err
	}
	t.typeLst = append(t.typeLst, rootField)
	return nil
}