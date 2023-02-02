package golang

import "strings"

func Comments(comments []string) []string {
	if len(comments) == 0 {
		return comments
	}
	for idx, comment := range comments {
		comments[idx] = strings.TrimSpace(comment)
	}
	if len(comments[0]) == 0 {
		comments = comments[1:]
	}
	if len(comments) == 0 {
		return comments
	}
	if len(comments[len(comments)-1]) == 0 {
		comments = comments[:len(comments)-1]
	}
	return comments
}

func Comment(comment string) string {
	comment = strings.TrimSpace(comment)
	return strings.ReplaceAll(strings.ReplaceAll(comment, "\\", "\\\\"), "\"", "\\\"")
}
