package tempconv

func cToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
