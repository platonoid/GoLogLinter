package slog

import "log/slog"

func bad() {
	var l *slog.Logger
	l.Info("hello world") // want "uppercase letter"
	l.Warn("warning!!!")  // want "only letters and numbers"
}

func good() {
	var l *slog.Logger
	l.Info("Hello world")
	l.Warn("Be careful 123")
}
