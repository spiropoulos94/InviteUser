package utils

import "strings"

// Function to extract domain from email
func ExtractDomain(email string) string {
	atIndex := strings.LastIndex(email, "@")
	if atIndex == -1 || atIndex == len(email)-1 {
		return "" // invalid email
	}
	return email[atIndex+1:]
}
