package greeting

import "testing"

func TestForWithName(t *testing.T) {
	g := For("Priya")
	if g.Name != "Priya" {
		t.Errorf("Name = %q, want %q", g.Name, "Priya")
	}
	if g.Message != "Hello, Priya!" {
		t.Errorf("Message = %q, want %q", g.Message, "Hello, Priya!")
	}
}

func TestForWithoutName(t *testing.T) {
	g := For("")
	if g.Name != "" {
		t.Errorf("Name = %q, want empty", g.Name)
	}
	if g.Message != "Hello, World!" {
		t.Errorf("Message = %q, want %q", g.Message, "Hello, World!")
	}
}
