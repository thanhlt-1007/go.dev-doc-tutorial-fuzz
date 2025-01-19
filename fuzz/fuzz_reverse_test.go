package main

import (
    "testing"
    "unicode/utf8"
)

func FuzzReverse(fuzz *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, testcase := range testcases {
        // Use f.Add to provide a seed corpus
        fuzz.Add(testcase)
    }

    fuzz.Fuzz(func(test *testing.T, orig string) {
        reverse := Reverse(orig)
        doubleRev := Reverse(reverse)
        if orig != doubleRev {
            test.Errorf("Before: %q, after: %q", orig, doubleRev)
        }

        if utf8.ValidString(orig) && !utf8.ValidString(reverse) {
            test.Errorf("Reverse produced invalid UTF-8 string %q", reverse)
        }
    })
}
