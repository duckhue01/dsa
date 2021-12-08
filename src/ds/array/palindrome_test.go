package array

import "testing"

func TestPalindrone(t *testing.T) {

	t.Run("test case #1", func(t *testing.T) {
		want := true
		got := palindrome("tact coa")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})

	t.Run("test case #2", func(t *testing.T) {
		want := false
		got := palindrome("abc")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #3", func(t *testing.T) {
		want := true
		got := palindrome("a")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #4", func(t *testing.T) {
		want := true
		got := palindrome("abbbb")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #5", func(t *testing.T) {
		want := true
		got := palindrome(" aa")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #6", func(t *testing.T) {
		want := true
		got := palindrome("bb ")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})

}



func TestPalindrone01(t *testing.T) {

	t.Run("test case #1", func(t *testing.T) {
		want := true
		got := palindrome01("tact coa")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})

	t.Run("test case #2", func(t *testing.T) {
		want := false
		got := palindrome01("abc")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #3", func(t *testing.T) {
		want := true
		got := palindrome01("a")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #4", func(t *testing.T) {
		want := true
		got := palindrome01("abbbb")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #5", func(t *testing.T) {
		want := true
		got := palindrome01(" aa")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #6", func(t *testing.T) {
		want := true
		got := palindrome01("bb ")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})
	t.Run("test case #6", func(t *testing.T) {
		want := true
		got := palindrome01("bb ")
		if want != got {
			t.Fatalf("want %v but got %v", want, got)

		}

	})

}