package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	result := Hello("World")

	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestHelloEmptyName(t *testing.T) {
	expected := "Hello, World!"
	result := Hello("")

	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestHelloSophie(t *testing.T) {
	expected := "Hello, Sophie!"
	result := Hello("Sophie")

	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}
