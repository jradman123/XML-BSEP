package common

import (
	"fmt"
	"regexp"
)

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

func BadTexts(input []string) bool {
	matched := false
	for _, in := range input {
		m := BadText(in)
		if m {
			matched = true
		}
	}
	return matched
}

func BadUsername(input string) bool {
	//justString, _ := regexp.MatchString(`^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$`, input)
	justString, _ := regexp.MatchString(`[^a-zA-Z0-9-_.]`, input)
	fmt.Println(justString)
	return justString //uklonjnen !
}

func BadName(input string) bool {
	justString, _ := regexp.MatchString(`[^a-zA-Z]`, input)
	return justString //vrati true ako je name los, ako ima ista sto nije slovo
}

func BadNumber(input string) bool {
	justNum, _ := regexp.MatchString(`[^0-9]`, input)
	return justNum //vrati true ako je username los, ako ima ista sto nije slovo
}

func BadJWTToken(input string) bool {
	justJWT, _ := regexp.MatchString(`[^a-zA-Z0-9\.]`, input)
	return justJWT
}

func BadDate(input string) bool {
	justDate, _ := regexp.MatchString(`^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2}(?:\.\d*)?)((-(\d{2}):(\d{2})|Z)?)$`, input)
	return !justDate
}

func BadText(input string) bool {
	badTxt, _ := regexp.MatchString(`([^a-zA-Z0-9 \?\!\.\,]+)`, input)
	return badTxt
}

func BadEmail(input string) bool {
	justMail, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input)
	return !justMail
}

//TODO:FIX THIS
func BadPassword(input string) bool {
	goodPass, _ := regexp.MatchString(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$`, input)
	fmt.Println(goodPass)
	return false
}

func BadId(input string) bool {
	justID, _ := regexp.MatchString(`[^a-zA-Z0-9]`, input)
	return justID
}

func BadImagePath(input string) bool {
	justPath, _ := regexp.MatchString(`^([\/]{1}[a-z0-9.]+)+(\/?){1}$|^([\/]{1})$`, input)
	return !justPath
}
func BadPaths(input []string) bool {
	matched := false
	for _, in := range input {
		m := BadImagePath(in)
		if m {
			matched = true
		}
	}
	return matched
}
