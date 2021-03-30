package main

import "testing"

func TestNewConfiguration(t *testing.T) {
	c := NewConfiguration()

	if c.count <= 0 {
		t.Errorf("Invalid default count: %d", c.count)
	}

	if len(c.language) == 0 {
		t.Errorf("Invalid default language: %s", c.language)
	}

	if len(c.subject) == 0 {
		t.Errorf("Invalid default subject: %s", c.subject)
	}
}
