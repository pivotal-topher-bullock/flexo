package flexo

type Formatter interface {
	Format(<-chan Token) []string
}

type FormatterConfig struct {
	LinkPrefix string
}
