package policies

import (
	"testing"
)

func TestPermShorthands(t *testing.T) {
	testCases := []struct {
		name     string
		perm     Perm
		expected uint8
	}{
		{"PermReadFull", PermReadFull, 3},
		{"PermUpdateField", PermUpdateField, 7},
		{"PermDeleteField", PermDeleteField, 15},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.perm != Perm(tc.expected) {
				t.Errorf("Expected %d, got %d", tc.expected, tc.perm)
			}
		})
	}
}

func TestIsAllowed(t *testing.T) {
	testCases := []struct {
		name     string
		scope    Scope
		class    FileClassification
		need     Perm
		expected bool
	}{
		{"Public - Public - ReadFull", ScopePublic, ClassPublic, PermReadFull, true},
		{"Admin - Protected - DeleteField", ScopeAdmin, ClassProtected, PermDeleteField, true},
		{"Admin - Public - DeleteField", ScopeAdmin, ClassPublic, PermDeleteField, true},
		{"Invalid scope and class", Scope(10), FileClassification(-1), PermReadFull, false},
		{"Invalid permission", ScopePublic, ClassPublic, 16, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsAllowed(tc.scope, tc.class, tc.need); got != tc.expected {
				t.Errorf("IsAllowed(%v, %v, %v) = %v want %v", tc.scope, tc.class, tc.need, got, tc.expected)
			}
		})
	}
}

func TestStringToClassification(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected FileClassification
		err      bool
	}{
		{"Valid - Public", "Public", ClassPublic, false},
		{"Valid - Admin", "Admin", ClassProtected, false},
		{"Invalid - Unknown", "Unknown", -1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := StringToClassification(tc.input)
			if (err != nil) != tc.err {
				t.Errorf("StringToClassification(%q) error: got %v, want %v", tc.input, err, tc.err)
			}
			if result != tc.expected {
				t.Errorf("StringToClassification(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
