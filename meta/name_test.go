package meta

import (
	"testing"
)

func TestIsExportedName(t *testing.T) {
	caseList := []struct {
		Name       string
		IsExported bool
	}{
		{"", false},
		{"a", false},
		{"A", true},
		{"_x9", false},
		{"ThisVariableIsExported", true},
		{"english", false},
		{"English", true},
		{"ελληνικά", false},
		{"Ελληνικά", true},
		{"русские", false},
		{"Русские", true},
		{"中文", false},
		{"日本語", false},
		{"にほんご", false},
		{"ニホンゴ", false},
	}

	for _, kase := range caseList {
		got := IsExportedName(kase.Name)
		if got != kase.IsExported {
			t.Errorf("IsExportedName(%s) = %v, expect %v", kase.Name, got, kase.IsExported)
		}
	}
}
