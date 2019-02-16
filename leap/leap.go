// package implements the function which tells whether given year is a leap one.
package leap

// takes year int and returns boolean value, whether given year is a leap one. Following rules:
// leap year is on every year that is evenly divisible by 4
// except every year that is evenly divisible by 100
// unless the year is also evenly divisible by 400
func IsLeapYear(year int) bool {
	return year%4 == 0  && (year%100 != 0 || year%400 == 0)
}
