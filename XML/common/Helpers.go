package common

import "regexp"

func CheckRegexSQL(input string) bool {
	matched1, _ := regexp.MatchString(`[\s\W\n\r\t]`, input)   //a whitespace character: [\t\n\f\r ]
	matched2, _ := regexp.MatchString(`[\'\"\#\$\%\_]`, input) //a non-whitespace character: [^\t\n\f\r ]
	matched3, _ := regexp.MatchString(`\\`, input)             //a whitespace character: [\t\n\f\r ]
	if matched1 || matched2 || matched3 {
		return true
	} else {
		return false
	}
}

func CheckForSQLInjection(input []string) bool {
	matched := false
	for _, in := range input {
		m := CheckRegexSQL(in)
		if m {
			matched = true
		}
	}
	return matched
}
