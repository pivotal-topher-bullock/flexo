package flexo

type Formatter interface {
	Format(<-chan Token) []string
}

type FormatterConfig struct {
	LinkPrefix string `long:"link-prefix" description:"Prefix for all links in the formatted output"`
}
