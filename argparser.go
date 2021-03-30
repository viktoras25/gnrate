package main

import (
	"errors"
	"strconv"
)

var validSubjects = []string{
	"name", "names",
	"firstname", "firstnames",
	"address", "addresses",
	"email", "emails",
	"url", "urls",
	"phone", "phones",
	"company", "companies",
	"sentence", "sentences", "lorem",
}

func parse(arguments []string) (*Configuration, error) {
	c := NewConfiguration()

	if len(arguments) == 0 {
		return c, errors.New("empty arguments")
	}

	// Assign subject
	c.subject, arguments = cutLast(arguments)

	if !isInSlice(c.subject, validSubjects) {
		return c, errors.New("invalid subject")
	}

	if len(arguments) == 0 {
		return c, nil
	}

	// Check if count provided and assign if appropriate
	argument := ""
	argument, arguments = cutLast(arguments)

	if isNumeric(argument) {
		c.count, _ = strconv.Atoi(argument)
	} else {
		c.language = argument
	}

	if len(arguments) == 0 {
		return c, nil
	}

	// Parse last argument, should be a number
	argument = ""
	argument, _ = cutLast(arguments)

	if isNumeric(argument) {
		c.count, _ = strconv.Atoi(argument)
	}

	return c, nil
}

func cutLast(arguments []string) (string, []string) {
	if len(arguments) == 0 {
		return "", []string{}
	}

	lastElement := arguments[len(arguments)-1]
	remainingSlice := arguments[:len(arguments)-1]

	return lastElement, remainingSlice
}

func isInSlice(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)

	return err == nil
}
