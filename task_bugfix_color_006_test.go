package color

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixColor006SourceContract(t *testing.T) {
    source, err := os.ReadFile("color.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if os.Stderr == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
