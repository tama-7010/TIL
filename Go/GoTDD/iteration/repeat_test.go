package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated, _ := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func TestRepeatTimes(t *testing.T) {
	_, repeatedTime := Repeat("a")
	expected := 5

	if repeatedTime != expected {
		t.Errorf("expected '%v' but got '%v'", expected, repeatedTime)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
