package exception

func or(message1 string, message2 string) string {
	if message1 != "" {
		return message1
	}
	return message2
}
