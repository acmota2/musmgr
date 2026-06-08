package controller

// TextOrNull returns a pointer to the string if it's not empty, otherwise nil.
func TextOrNull(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
