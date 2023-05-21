package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace," + " " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	createBorder:= strings.Repeat("*", numStarsPerLine)
    return createBorder + "\n" + welcomeMsg + "\n" + createBorder
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// Remove all stars from the text
    cleanedMsg := strings.ReplaceAll(oldMsg, "*", "")
    
    // Remove leading and trailing whitespaces
    cleanedMsg = strings.TrimSpace(cleanedMsg)
    
    return cleanedMsg
}
