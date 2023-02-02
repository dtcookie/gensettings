package collections

import (
	"fmt"
	"strings"
)

type Set[T any] []T

func (s Set[T]) Contains(v T) bool {
	switch typedV := any(v).(type) {
	case string:
		for _, elem := range s {
			if any(elem).(string) == typedV {
				return true
			}
		}
	}
	return false
}

func (s Set[T]) ToString() string {
	result := ""
	for _, elem := range s {
		result = result + fmt.Sprintf(" %v", elem)
	}
	return strings.TrimSpace(result)
}
