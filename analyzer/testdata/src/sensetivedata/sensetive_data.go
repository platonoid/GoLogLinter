package sensitivedata

import "testlog"

func test(password string) {
	var l testlog.Logger
	l.Info(password) // want "sensitive variable"
}
