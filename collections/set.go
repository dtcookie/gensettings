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
	sep := ""
	for _, elem := range s {
		if any(elem).(string) == "environment" {
			continue
		}
		result = result + sep + fmt.Sprintf("%v", elem)
		sep = ", "
	}
	return strings.TrimSpace(result)
}
