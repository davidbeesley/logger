package logger

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestLogger(t *testing.T) {
	Level = DEBUG

	var count = 20
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			Debugf("Thread example: %d", i)
			Debug("Thread exmample: ", i)
		}(i)
	}
	Debug("Debug example")
	Info("Info example")
	Warning("Warning example")
	Error("Error example")
	Print("Print and ")
	Println("Println")
	wg.Wait()
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
