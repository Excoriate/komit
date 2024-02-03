package ui

import (
	"strings"

	"github.com/pterm/pterm"
)

// displayMessage is a generalized function to display messages using pterm.
// messageType can be "error", "info", or "success" to determine the message style.
func displayMessage(messageType, prefix, message string, err error) {
	// Set the prefix style based on the messageType
	if prefix != "" {
		switch messageType {
		case "error":
			pterm.Error.Prefix = pterm.Prefix{
				Text:  strings.ToUpper(prefix),
				Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed),
			}
		case "info":
			pterm.Info.Prefix = pterm.Prefix{
				Text:  strings.ToUpper(prefix),
				Style: pterm.NewStyle(pterm.BgBlue, pterm.FgWhite),
			}
		case "success":
			pterm.Success.Prefix = pterm.Prefix{
				Text:  strings.ToUpper(prefix),
				Style: pterm.NewStyle(pterm.BgGreen, pterm.FgWhite),
			}

		case "fatal":
			pterm.Fatal.Prefix = pterm.Prefix{
				Text:  strings.ToUpper(prefix),
				Style: pterm.NewStyle(pterm.BgRed, pterm.FgWhite),
			}
		}
	}

	var errMsg string
	if err != nil {
		if message != "" {
			errMsg = message + ": " + err.Error()
		} else {
			errMsg = err.Error()
		}
	} else {
		errMsg = message
	}

	// Display the message based on the messageType
	switch messageType {
	case "error":
		pterm.Error.Println(errMsg)
	case "info":
		pterm.Info.Println(errMsg)
	case "success":
		pterm.Success.Println(errMsg)
	case "fatal":
		pterm.Fatal.Println(errMsg)
	}
}

// Error displays an error message, optionally prefixed and including an error.
func Error(prefix, message string, err error) {
	displayMessage("error", prefix, message, err)
}

func Fatal(prefix, message string, err error) {
	displayMessage("fatal", prefix, message, err)
}

// Info displays an informational message, optionally prefixed.
func Info(prefix, message string) {
	displayMessage("info", prefix, message, nil)
}

// Success displays a success message, optionally prefixed.
func Success(prefix, message string) {
	displayMessage("success", prefix, message, nil)
}
