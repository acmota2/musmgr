package policies

import (
	"testing"
)

func TestPermShorthands(t *testing.T) {
	tests := []struct {
		name     string
		perm     Perm
		expected uint8
	}{
		{"PermReadFull", PermReadFull, 3},
		{"PermUpdateField", PermUpdateField, 7},
		{"PermDeleteField", PermDeleteField, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.perm != Perm(tt.expected) {
				t.Errorf("Expected %d, got %d", tt.expected, tt.perm)
			}
		})
	}
}

func TestIsAllowed_ScopePublic(t *testing.T) {
	tests := []struct {
		name     string
		scope    Scope
		class    FileClassification
		need     Perm
		expected bool
	}{
		{"Public - Public - ReadFull", ScopePublic, ClassPublic, PermReadFull, true},
		{"Public - Public - Write", ScopePublic, ClassPublic, PermWrite, false},
		{"Public - Public - List", ScopePublic, ClassPublic, PermList, true},
		{"Public - Public - Get", ScopePublic, ClassPublic, PermGet, true},
		{"Public - Public - InvalidPerm", ScopePublic, ClassPublic, 16, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllowed(tt.scope, tt.class, tt.need); got != tt.expected {
				t.Errorf("IsAllowed(%v, %v, %v) = %v, want %v", tt.scope, tt.class, tt.need, got, tt.expected)
			}
		})
	}
}

func TestIsAllowed_ScopeAdmin(t *testing.T) {
	tests := []struct {
		name     string
		scope    Scope
		class    FileClassification
		need     Perm
		expected bool
	}{
		{"Admin - Protected - DeleteField", ScopeAdmin, ClassProtected, PermDeleteField, true},
		{"Admin - Public - DeleteField", ScopeAdmin, ClassPublic, PermDeleteField, true},
		{"Admin - Public - ReadFull", ScopeAdmin, ClassPublic, PermReadFull, true},
		{"Admin - Public - Write", ScopeAdmin, ClassPublic, PermWrite, true},
		{"Admin - Public - List", ScopeAdmin, ClassPublic, PermList, true},
		{"Admin - Public - Get", ScopeAdmin, ClassPublic, PermGet, true},
		{"Admin - Public - InvalidPerm", ScopeAdmin, ClassPublic, 16, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllowed(tt.scope, tt.class, tt.need); got != tt.expected {
				t.Errorf("IsAllowed(%v, %v, %v) = %v want %v", tt.scope, tt.class, tt.need, got, tt.expected)
			}
		})
	}
}

func TestStringToClassification(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected FileClassification
		err      bool
	}{
		{"Valid - Public", "Public", ClassPublic, false},
		{"Valid - Admin", "Admin", ClassProtected, false},
		{"Invalid - Unknown", "Unknown", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToClassification(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("StringToClassification(%q) error: got %v, want %v", tt.input, err, tt.err)
			}
			if result != tt.expected {
				t.Errorf("StringToClassification(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
