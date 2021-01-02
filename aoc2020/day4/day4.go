package day4

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type passport struct {
	BirthYear      string `validate:"name=byr,presence"`
	IssueYear      string `validate:"name=iyr,presence"`
	ExpirationYear string `validate:"name=eyr,presence"`
	Height         string `validate:"name=hgt,presence"`
	HairColor      string `validate:"name=hcl,presence"`
	EyeColor       string `validate:"name=ecl,presence"`
	PassportID     string `validate:"name=pid,presence"`
	CountryID      string `validate:"name=cid"`
}

func (p passport) GetFieldForTagName(name string) (string, error) {
	switch name {
	case "byr":
		return p.BirthYear, nil
	case "iyr":
		return p.IssueYear, nil
	case "eyr":
		return p.ExpirationYear, nil
	case "hgt":
		return p.Height, nil
	case "hcl":
		return p.HairColor, nil
	case "ecl":
		return p.EyeColor, nil
	case "pid":
		return p.PassportID, nil
	case "cid":
		return p.CountryID, nil
	}
	return "", errors.New("Unknown tag name")
}

func (p passport) SetFieldForTagName(name string, value string) (passport, error) {
	found := false
	switch name {
	case "byr":
		p.BirthYear = value
		found = true
	case "iyr":
		p.IssueYear = value
		found = true
	case "eyr":
		p.ExpirationYear = value
		found = true
	case "hgt":
		p.Height = value
		found = true
	case "hcl":
		p.HairColor = value
		found = true
	case "ecl":
		p.EyeColor = value
		found = true
	case "pid":
		p.PassportID = value
		found = true
	case "cid":
		p.CountryID = value
		found = true
	}
	if !found {
		return p, errors.New("Unknown tag name")
	}
	return p, nil
}

// CountValidPassports takes a mapgrid (. for open, # for tree)
// and counts the number of trees that would be hit traversing
// from top to bottom using slope 3 right, 1 down
// assuming a horizontally tiled repeating map
func CountValidPassports(in string) int {
	fmt.Printf("\n\nInput:\n%+v\n", in)
	records := parseCollection(in)
	validCount := 0
	for _, record := range records {
		passportRecord, errs := createPassport(record)
		errs = append(errs, validatePassportRecord(passportRecord)...)
		if len(errs) == 0 {
			validCount++
		}
		fmt.Printf("Record: %+v, Passport valid? %t\n", passportRecord, len(errs) == 0)
	}
	return validCount
}

type validator struct {
	name     string
	presence bool
}

func buildValidator(tag string) validator {
	args := strings.Split(tag, ",")
	var v validator
	arg0vals := strings.Split(args[0], "=")
	v.name = arg0vals[1]
	v.presence = false
	if len(args) == 2 {
		v.presence = (args[1] == "presence")
	}
	return v
}

func (v validator) Validate(val passport) (bool, error) {
	fieldVal, _ := val.GetFieldForTagName(v.name)
	if v.presence && fieldVal == "" {
		return false, fmt.Errorf("Record is missing field '%s'", v.name)
	}
	return true, nil
}

func createPassport(record string) (passport, []error) {
	passportRecord := passport{}
	errs := []error{}

	fields := parseRecord(record)
	var err error

	for _, field := range fields {
		key, value := parseField(field)
		passportRecord, err = passportRecord.SetFieldForTagName(key, value)
		if err != nil {
			errs = append(errs,
				fmt.Errorf("Unable to set field with tag %s to %s", key, value))
		}
	}
	return passportRecord, errs
}

func validatePassportRecord(passportRecord passport) []error {
	errs := []error{}
	v := reflect.ValueOf(passportRecord)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("validate")
		validator := buildValidator(tag)
		valid, err := validator.Validate(passportRecord)
		if !valid && err != nil {
			errs = append(errs,
				fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}
	return errs
}

var recordBreakRE = regexp.MustCompile(`\n\n`)
var fieldBreakRE = regexp.MustCompile(`\s+`)

func parseCollection(data string) []string {
	return recordBreakRE.Split(data, -1)
}

func parseRecord(record string) []string {
	return fieldBreakRE.Split(record, -1)
}

func parseField(field string) (string, string) {
	parts := strings.Split(field, ":")
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
