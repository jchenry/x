package log

// None provides a logger that doesnt log anything
type None struct{}

func (n None) Fatal(v ...any)                 {}
func (n None) Fatalf(format string, v ...any) {}
func (n None) Fatalln(v ...any)               {}
func (n None) Panic(v ...any)                 {}
func (n None) Panicf(format string, v ...any) {}
func (n None) Panicln(v ...any)               {}
func (n None) Print(v ...any)                 {}
func (n None) Printf(format string, v ...any) {}
func (n None) Println(v ...any)               {}
