package specialsymbols

import "testlog"

func test() {
	var l testlog.Logger
	l.Info("Hello!!!") // want "only letters and numbers"
}
