package action

import "io"

// Context makes easier to pass args and flags around
type Context struct {
	Args  []string
	Flags map[string]string
}

// Action ...
type Action struct {
	Output io.Writer
}

// Execute is a type func that takes Context as parameter and returns an error
type Execute func(Context) error

// WriteString takes string and prints it to the Output
func (a Action) WriteString(s string) (n int, err error) {
	n, err = a.Output.Write([]byte(s))
	return
}
