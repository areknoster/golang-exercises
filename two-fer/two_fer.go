//  Package twofer implements a single function, which turns a name into a 2-fer sentence.
package twofer

// ShareWith takes a string "name" as argument and returns a string following the rule: "One for [name], one for me".
// If empty string is given, puts "you" instead of name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
