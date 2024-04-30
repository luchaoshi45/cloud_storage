package configurator

type Config interface {
	Parse(location string) (Config, error)
	GetAttr(attr string) string
}

func NewMysqlConfig() Config {
	return new(MysqlConfig)
}
