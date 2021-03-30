package main

import "testing"

func TestParseEmptyArgumentsList(t *testing.T) {
	_, err := parse([]string{})

	if err == nil {
		t.Error("Expected error on empty slice")
	}
}

func TestParseSubjectOnly(t *testing.T) {
	c, err := parse([]string{"email"})

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if c.subject != "email" {
		t.Errorf("Failed parsing subject")
	}

	_, err = parse([]string{"invalid"})

	if err == nil {
		t.Error("Expected error on invalid subject")
	}
}

func TestParseTwoArguments(t *testing.T) {
	// Parse subject and count
	c, err := parse([]string{"10", "emails"})

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if c.subject != "emails" {
		t.Errorf("Failed parsing subject")
	}

	if c.count != 10 {
		t.Errorf("Failed parsing count")
	}

	// Parse subject and language
	c, err = parse([]string{"german", "companies"})

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if c.subject != "companies" {
		t.Errorf("Failed parsing subject")
	}

	if c.language != "german" {
		t.Errorf("Failed parsing count")
	}
}

func TestParseThreeArguments(t *testing.T) {
	// Parse subject and count
	c, err := parse([]string{"5", "russian", "addresses"})

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if c.subject != "addresses" {
		t.Errorf("Failed parsing subject")
	}

	if c.count != 5 {
		t.Errorf("Failed parsing count")
	}

	if c.language != "russian" {
		t.Errorf("Failed parsing language")
	}
}
