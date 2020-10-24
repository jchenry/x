package log

// Logger is a logging interface with only the essentials that a function that needs to log should care about. Compatible with standard Go logger.
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// None provides a logger that doesnt log anything
type None struct{}

func (n None) Print(v ...interface{})                 {}
func (n None) Printf(format string, v ...interface{}) {}
func (n None) Println(v ...interface{})               {}
