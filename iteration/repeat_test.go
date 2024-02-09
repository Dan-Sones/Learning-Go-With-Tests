package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	expected := "aaaaaa"

	t.Run("pass using my solution", func(t *testing.T) {
		repeated := Repeat("a", 6)

		assertOutputCorrect(t, repeated, expected)

	})

	t.Run("pass using inbuilt string method", func(t *testing.T) {
		repeated := strings.Repeat("a", 6)
		assertOutputCorrect(t, repeated, expected)
	})
}

func assertOutputCorrect(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	rep := Repeat("d", 5)
	fmt.Println(rep)
	// Output: ddddd
}
