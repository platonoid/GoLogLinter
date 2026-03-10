package english

import "testlog"

func test() {
	var l testlog.Logger
	l.Info("Ошибка") // want "only English letters"
}
