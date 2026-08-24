package color

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixColor009SourceContract(t *testing.T) {
    source, err := os.ReadFile("color.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return os.Getenv(\"NO_COLOR\") != \"\"") {
        t.Fatalf("expected source contract is missing")
    }
}
