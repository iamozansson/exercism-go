// Pakage leap provides function that is controlling if the year is leap or not.
package leap

// Function leap controls if the year can be divided to 4 even or not.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
