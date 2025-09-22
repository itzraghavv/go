package main

func isValidPassword(password string) bool {
	hasDigit := false
	hasUpper := false

	if len(password) >= 5 && len(password) <= 12 {
		for i := 0; i < len(password); i++ {
			if password[i] >= 'A' && password[i] <= 'Z' {
				hasUpper = true
			}
			if password[i] >= '0' && password[i] <= '9' {
				hasDigit = true
			}
		}
		return hasDigit && hasUpper
	}
	return false

}
