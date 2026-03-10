package capital

import "testlog"

func test() {
	var l testlog.Logger
	l.Info("hello world") // want "uppercase letter"
}
