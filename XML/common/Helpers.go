package common

import "regexp"

func CheckRegexSQL(input string) bool {
	matched1, _ := regexp.MatchString(`[\s\n\r\t]`, input)
	matched2, _ := regexp.MatchString(`[\'\"\#\$\%\_\|]`, input)
	matched3, _ := regexp.MatchString(`\\`, input)
	matched4, _ := regexp.MatchString(`\>\<\&`, input)
	if matched1 || matched2 || matched3 || matched4 {
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
