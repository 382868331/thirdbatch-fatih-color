package color

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisColor008SourceContract(t *testing.T) {
    source, err := os.ReadFile("color.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "defer c.unset()") {
        t.Fatalf("expected source contract is missing")
    }
}
