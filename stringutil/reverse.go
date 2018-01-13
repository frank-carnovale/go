package stringutil // utility functions for working with strings.

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	last := len(r) - 1
	if last <= 0 {
		return s
	}
	return string(r[last]) + Reverse(string(r[:last]))
}
