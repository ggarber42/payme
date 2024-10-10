package logger

type ILogger interface{
	print(level Level, message string, properties map[string]string) (int, error)
	PrintInfo(message string, properties map[string]string)
	PrintError(err error, properties map[string]string)
	PrintFatal(err error, properties map[string]string)
}