package repository

import "fmt"

type RepositoryError struct {
	msg string
}

func NewRepositoryError(msg string) *RepositoryError {
	return &RepositoryError{msg}
}

func (re *RepositoryError) Error() string {
	return fmt.Sprintf("Repository error: %s", re.msg)
}
