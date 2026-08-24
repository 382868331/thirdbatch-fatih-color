package color

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixColor001SourceContract(t *testing.T) {
    source, err := os.ReadFile("color.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if c == nil || c2 == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
