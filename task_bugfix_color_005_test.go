package color

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixColor005SourceContract(t *testing.T) {
    source, err := os.ReadFile("color.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if os.Stdout == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if false && os.Stdout == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
