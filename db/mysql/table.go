package mysql

// Table 抽象表格
type Table interface {
	SetAttrs(attrs map[string]interface{}) error
	GetAttr(attr string) (interface{}, error)
}

type FieldNotFoundError struct {
	Field string
}

func (e *FieldNotFoundError) Error() string {
	return "field not found: " + e.Field
}

type FieldNotSettableError struct {
	Field string
}

func (e *FieldNotSettableError) Error() string {
	return "field not settable: " + e.Field
}

type FieldTypeMismatchError struct {
	Field    string
	Expected string
	Actual   string
}

func (e *FieldTypeMismatchError) Error() string {
	return "field type mismatch for " + e.Field + ": expected " + e.Expected + ", got " + e.Actual
}

type FieldNotGettableError struct {
	Field string
}

func (e *FieldNotGettableError) Error() string {
	return "field not gettable: " + e.Field
}
