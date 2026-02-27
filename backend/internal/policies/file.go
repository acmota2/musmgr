package policies

import (
	"fmt"
	"math"
)

type Scope uint8

const (
	ScopePublic = iota
	ScopeAdmin
)

// file classification in this domain can be ordered in a lattice
type FileClassification int16

var MaxClassification = math.MaxInt16

const (
	ClassPublic FileClassification = iota
	ClassProtected
)

type Perm uint8

const (
	PermList Perm = 1 << iota
	PermGet
	PermWrite
	PermDelete
)

// permission lattice shortcuts
const (
	PermNone        = 0 // for completeness
	PermReadFull    = PermList | PermGet
	PermUpdateField = PermList | PermGet | PermWrite
	PermDeleteField = PermList | PermGet | PermWrite | PermDelete
)

var policies = map[Scope]map[FileClassification]Perm{
	ScopePublic: {
		ClassPublic: PermReadFull,
	},
	ScopeAdmin: {
		ClassProtected: PermDeleteField,
		ClassPublic:    PermDeleteField,
	},
}

func IsAllowed(scope Scope, class FileClassification, need Perm) bool {
	if scope == ScopeAdmin {
		return true
	}
	return policies[scope][class]&need == need
}

func StringToClassification(s string) (FileClassification, error) {
	switch s {
	case "Public":
		return ClassPublic, nil
	case "Admin":
		return ClassProtected, nil
	default:
		return -1, fmt.Errorf("Invalid FileClassification format")
	}
}
