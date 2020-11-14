package impl

type Cat interface  {
	Echo(string) string
}
func Echo(message string) string {
	return message
}