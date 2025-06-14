package util

func LogError(err error) {
	if err != nil {
		println("Error:", err.Error())
	}
}

func LogInfo(message string, data ...interface{}) {
	if message != "" {
		println("Info:", message)
	}
}

func LogDebug(message string) {
	if message != "" {
		println("Debug:", message)
	}
}
