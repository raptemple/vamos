package main

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestOptQuote(t *testing.T) {
    want := "If a program is too slow, it must have a loop."
    if got := OptQuote(); got != want {
        t.Errorf("OptQuote() = %q, want %q", got, want)
    }
}

func TestGlassQuote(t *testing.T) {
    want := "I can eat glass and it doesn't hurt me."
    if got := GlassQuote(); got != want {
        t.Errorf("OptQuote() = %q, want %q", got, want)
    }
}

func TestGoQuote(t *testing.T) {
    want := "Don't communicate by sharing memory, share memory by communicating."
    if got := GoQuote(); got != want {
        t.Errorf("OptQuote() = %q, want %q", got, want)
    }
}
