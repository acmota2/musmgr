package model

type SavableModel[T any] interface {
	Save() (*T, error)
}
