package policies

import "math"

type Scope uint8

const (
	ScopePublic = iota
	ScopeAdmin
)

type FileClassification int16

var MaxClassification = math.MaxInt16

const (
	ClassPublic = iota
	ClassProtected
)

type Perm uint8

const (
	PermRead Perm = 1 << iota
	PermWrite
	PermDelete
)

const (
	PermNoneField   = 0
	PermReadField   = PermRead
	PermUpdateField = PermRead | PermWrite
	PermDeleteField = PermRead | PermWrite | PermDelete
)

var policies = map[Scope]map[FileClassification]Perm{
	ScopePublic: {
		ClassPublic: PermReadField,
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
