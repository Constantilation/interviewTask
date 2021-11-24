package MyError

type MultiLogger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Sync() error
}

type Errors struct {
	Alias string
	Text  string
}

func (e *Errors) Error() string {
	return e.Alias
}