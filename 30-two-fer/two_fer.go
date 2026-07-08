// Package twofer provides a function for sharing cookies between two people.
package twofer

// ShareWith returns a sentence about sharing a cookie with the given person.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
