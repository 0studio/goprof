package goprof

var logError func(v ...interface{})
var logInfo func(v ...interface{})

func Init(_logError func(...interface{}), _logInfo func(v ...interface{})) {
	logError = _logError
	logInfo = _logInfo
}
