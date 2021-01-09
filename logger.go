package hydra

type Logger interface {
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Println(args ...interface{})
}
