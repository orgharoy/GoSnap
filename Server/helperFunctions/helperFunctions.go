package helperfunctions

import "regexp"

func IsValidPassword(password string) bool {
	// Check for at least 8 characters
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	uppercaseRegex := regexp.MustCompile("[A-Z]")
	if !uppercaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one lowercase letter
	lowercaseRegex := regexp.MustCompile("[a-z]")
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	// Check for at least one special character (you can modify the character class)
	specialCharRegex := regexp.MustCompile("[!@#$%^&*()_+]")
	return specialCharRegex.MatchString(password)
}
