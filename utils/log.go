package utils

import "hiveview"

//日志打印
func LogPrint(logLevel string, msg interface{}) {
	switch {
	case logLevel == "warn":
		hiveview.CONFIG.Logger.Warn(msg)
	case logLevel == "err":
		hiveview.CONFIG.Logger.Error(msg)
	case logLevel == "fatal":
		hiveview.CONFIG.Logger.Fatal(msg)
	default:
		hiveview.CONFIG.Logger.Info(msg)
	}
}
