package main

import (
	"testing"
)

// func TestHello(t *testing.T) {
// 	got := HelloOld2()
// 	want := "Hello, World!"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestHello2(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestHello3(t *testing.T) {
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris")
// 		want := "Hello, Chris"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }

// func TestHello4(t *testing.T) {
// 	assertCorrectMessage := func(t testing.TB, got, want string) {
// 		t.Helper()
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	}
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris")
// 		want := "Hello, Chris"
// 		assertCorrectMessage(t, got, want)
// 	})
// 	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"
// 		assertCorrectMessage(t, got, want)
// 	})
// }

// func TestHello5(t *testing.T) {
// 	assertCorrectMessage := func(t testing.TB, got, want string) {
// 		t.Helper()
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	}
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris")
// 		want := "Hello, Chris"
// 		assertCorrectMessage(t, got, want)
// 	})
// 	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"
// 		assertCorrectMessage(t, got, want)
// 	})
// 	t.Run("say 'Hola, Chris!' when an Spanish is supplied as the language", func(t *testing.T) {
// 		got := Hello("Chris", "Spanish")
// 		want := "Hola, Chris"
// 		assertCorrectMessage(t, got, want)
// 	})
// }

func TestHello6(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	tests := []struct {
		language string
		name     string
		want     string
	}{
		{"English", "", "Hello, World"},
		{"English", "Chris", "Hello, Chris"},
		{"Spanish", "Chris", "Hola, Chris"},
		{"Italian", "Chris", "Ciao, Chris"},
		{"German", "Chris", "Hallo, Chris"},
		{"French", "Chris", "Bonjour, Chris"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.name, tt.language)
			want := tt.want
			assertCorrectMessage(t, got, want)
		})
	}
}
