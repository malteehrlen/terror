package terror

import (
"testing"
)

func TestTryFinally(t *testing.T) {
    tb := &TryBlock[string]{}
    tb.Try = func() string {return "dont want this"}
    tb.Finally = func(_ string) string {
        return "want this"
    }

    a, err := tb.Run()
    if err != nil {
        t.Fatalf("uncaught error")
    }
    if a != "want this" {
        t.Fatalf("wrong output")
    }
}

func TestTryCatchNoThrow(t *testing.T) {
    tb := &TryBlock[string]{}
    tb.Try = func() string {return "want this"}
    tb.Catch = func(c Catchable) string {
        _, ok := c.(string)
        if ok {
            return "dont want this"
        }
        return "uncaught error"
    }
    a, err := tb.Run()
    if err != nil {
        t.Fatalf("uncaught error")
    }
    if a != "want this" {
        t.Fatalf("wrong output")
    }
}

func TestTryCatchThrow(t *testing.T) {
    tb := &TryBlock[string]{}
    tb.Try = func() string {
        Throw("aaargh")
        return "dont want this"
    }
    tb.Catch = func(c Catchable) string {
        _, ok := c.(string)
        if ok {
            return "want this"
        }
        return "uncaught error"
    }
    a, err := tb.Run()
    if err != nil {
        t.Fatalf("uncaught error")
    }
    if a != "want this" {
        t.Fatalf("wrong output")
    }
}
