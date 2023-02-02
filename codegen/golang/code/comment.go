package code

import "strings"

type Comment []string

func (me Comment) Multiline() []string {
	lines := []string{}
	for _, line := range me {
		lines = append(lines, strings.Split(line, "\n")...)
	}
	return lines
}

type Commentable interface {
	Comment() Comment
}
