package gobgcolor

func Black(text string) string {
	result := "\x1b[37m\x1b[40m" + text + "\x1b[0m\n"
	return result
}

func Red(text string) string {
	result := "\x1b[37m\x1b[41m" + text + "\x1b[0m\n"
	return result
}

func Green(text string) string {
	result := "\x1b[37m\x1b[42m" + text + "\x1b[0m\n"
	return result
}

func Yellow(text string) string {
	result := "\x1b[37m\x1b[43m" + text + "\x1b[0m\n"
	return result
}

func Blue(text string) string {
	result := "\x1b[37m\x1b[44m" + text + "\x1b[0m\n"
	return result
}

func Magenta(text string) string {
	result := "\x1b[37m\x1b[45m" + text + "\x1b[0m\n"
	return result
}

func Cyan(text string) string {
	result := "\x1b[30m\x1b[46m" + text + "\x1b[0m\n"
	return result
}

func White(text string) string {
	result := "\x1b[30m\x1b[47m" + text + "\x1b[0m\n"
	return result
}
