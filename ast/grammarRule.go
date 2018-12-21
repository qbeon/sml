package ast

import (
	"regexp"
	"unicode"
)

var grammarRulePropertyName = regexp.MustCompile(`^[a-z][A-Za-z]{0,31}$`)
var grammarRuleTypeName = regexp.MustCompile(`^[A-Z][A-Za-z]{0,31}$`)
var grammarRuleModelName = regexp.MustCompile(`^[A-Z][A-Za-z]{0,31}$`)

func validateModelName(nameFragment *Fragment) error {
	if !grammarRuleModelName.MatchString(string(nameFragment.Raw)) {
		err := &ErrInvalidModelName{Fragment: nameFragment}
		if unicode.IsLower(nameFragment.Raw[0]) {
			err.Violation = "model names must begin with a capital letter"
		}
		return err
	}
	return nil
}

func validateTypeName(
	nameFragment *Fragment,
	graphNodeType GraphNodeType,
) error {
	if !grammarRuleTypeName.MatchString(string(nameFragment.Raw)) {
		err := &ErrInvalidTypeName{Fragment: nameFragment}
		if unicode.IsLower(nameFragment.Raw[0]) {
			err.Violation = "type names must begin with a capital letter"
		}
		return err
	}
	return nil
}

func validatePropertyName(
	nameFragment *Fragment,
	graphNodeType GraphNodeType,
) error {
	if !grammarRulePropertyName.MatchString(string(nameFragment.Raw)) {
		err := &ErrInvalidPropertyName{Fragment: nameFragment}
		if !unicode.IsLower(nameFragment.Raw[0]) {
			err.Violation = "property names must begin with a lower case letter"
		}
		return err
	}
	return nil
}
