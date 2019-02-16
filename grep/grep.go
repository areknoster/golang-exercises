// Package grep does awesome stuff.
package grep

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Search searches and its all it does
func Search(pattern string, flags, files []string) (result []string) {
	var withNumber, caseInsensitive, wholeLine, returnFile, invert bool
	var matches []string
	for _, v := range flags {

		switch v[1] {
		case 'n':
			withNumber = true
		case 'i':
			caseInsensitive = true
		case 'l':
			returnFile = true
		case 'x':
			wholeLine = true
		case 'v':
			invert = true

		}
	}

	for _, v := range files {
		f, err := ioutil.ReadFile(v)
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(f), "\n")
		matches = findLines(lines, pattern, withNumber, caseInsensitive, wholeLine, returnFile, invert)
		if len(matches) == 0 {
			continue
		}
		if returnFile {
			matches[0] = v
		} else if len(files) > 1 {
			for i := 0; i < len(matches); i++ {
				matches[i] = v + ":" + matches[i]
			}
		}

		result = append(result, matches...)
	}

	fmt.Printf("%v\n", len(result))
	if len(result) == 0 {
		return []string{}
	}
	return result
}

func findLines(lines []string, pattern string, withNumber, caseInsensitive, wholeLine, returnFile, invert bool) (matches []string) {
	var vCheck string
	if caseInsensitive {
		pattern = strings.ToLower(pattern)
	}
	for i, v := range lines {
		if caseInsensitive {
			vCheck = strings.ToLower(v)
		} else {
			vCheck = v
		}
		if checkMatch(vCheck, pattern, wholeLine) == invert {
			continue
		}
		if returnFile {
			return append(matches, "")
		}
		if withNumber {
			v = fmt.Sprintf("%v:%v", i+1, v)
		}
		if v != "" {
			matches = append(matches, v)
		}
	}
	return matches
}

func checkMatch(v, pattern string, wholeLine bool) bool {
	if wholeLine {
		return v == pattern
	}
	return strings.Contains(v, pattern)
}
