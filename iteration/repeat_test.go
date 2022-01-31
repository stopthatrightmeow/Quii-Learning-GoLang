package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 10, "")
	expected := "aaaaaaaaaa"

	if repeat != expected {
		t.Errorf("Sent %q, expected %q", repeat, expected)
	}
}

func TestUpperRepeat(t *testing.T) {
	upperrepeat := Repeat("a", 10, "Up")
	expected := "AAAAAAAAAA"
	if upperrepeat != expected {
		t.Errorf("Sent %q, expected %q", upperrepeat, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10, "")
	}
}

func BenchmarkUpperRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10, "Up")
	}
}
