package contracts

// Severity provides constants for the severity level of a traced statement.
type Severity int32

const (
	// Verbose represents a verbose message, typically for debugging
	Verbose Severity = 0

	// Information represents an informational message.
	Information Severity = 1

	// Warning represents a message of interest that is not critical to the application's ability to function
	Warning Severity = 2

	// Error represents a message that an exceptional condition has occured but has not caused the application to fail
	Error Severity = 3

	// Critical represents a message that an exceptional condition has occured which has caused the application to fail.
	Critical Severity = 4
)

// ToString converts the Severity to a readable string
func (s Severity) ToString() string {
	switch int(s) {
	case 0:
		return "Verbose"
	case 1:
		return "Information"
	case 2:
		return "Warning"
	case 3:
		return "Error"
	case 4:
		return "Critical"
	default:
		return "<unknown>"
	}
}
