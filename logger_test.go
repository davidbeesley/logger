package logger

import "testing"

func TestLogger(t *testing.T) {
	Level = DEBUG
	Debug("Debug example")
	Info("Info example")
	Warning("Warning example")
	Error("Error example")
	/*
		cases := []struct {
			in, want string
		}{
			{"Hello, world", "dlrow ,olleH"},
			{"Hello, 世界", "界世 ,olleH"},
			{"", ""},
		}
		for _, c := range cases {
			got := Reverse(c.in)
			if got != c.want {
				t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	*/
}
