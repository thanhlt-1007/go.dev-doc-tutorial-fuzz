package main

import (
    "testing"
)

type Testcase struct {
    in, want string
}

func TestRever(t *testing.T) {
    testcases := []Testcase {
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }

    for _, testcase := range testcases {
        reverse := Reverse(testcase.in)
        if reverse != testcase.want {
            t.Errorf("Reverse: %q, want %q", reverse, testcase.want)
        }
    }
}
