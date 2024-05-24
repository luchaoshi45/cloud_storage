package configurator

type Config interface {
	parse() error
	GetAttr(attr string) string
}
