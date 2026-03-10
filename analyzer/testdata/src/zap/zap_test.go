package zap

import "go.uber.org/zap"

func bad() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("hello world")    // want "uppercase letter"
	sugar.Warn("Предупреждение") // want "only English letters"
	sugar.Error("password")      // want "sensitive variable"
}

func good() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Info("Hello world")
	sugar.Warn("Be careful")
}
