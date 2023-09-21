package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q instead", expected, repeated)
	}
}

// testing gives us access to b.N, when this is executed,
// it runs b.N times and measures how long it takes.
// To run the benchmarks do go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
