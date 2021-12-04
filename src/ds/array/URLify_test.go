package array

import (
	"fmt"
	"testing"
)

// input: "duc khue 01", 11
// output: "duc%20khue%2001"
func TestURLify01(t *testing.T) {

	want := "duc%20khue%2001"
	got := URLify01("duc khue 01", 11)
	fmt.Println(got)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}

}

func TestURLify02(t *testing.T) {

	want := "duc%20khue%2001"
	got := URLify02([]rune("duc khue 01"), 11)
	fmt.Println(got)
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}

}

func BenchmarkURLify02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URLify02([]rune("            "), 11)
	}
}

func BenchmarkURLify01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URLify01("           ", 11)
	}
}
