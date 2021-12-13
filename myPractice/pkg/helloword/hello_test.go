package helloword

import (
	"testing"
)

func TestHelloWord(t *testing.T) {
	if Helloword() != "Hello word!" {
		t.Error("Expected println and return \"Hello word!\"")
	}
	if Addition(2, 3) != 5 {
		t.Error("Expected 2 + 3 = 5")
	}
}
