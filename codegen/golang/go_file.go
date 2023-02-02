package golang

type GoFile struct {
	Package string
	Imports []string
	Code    string
	Name    string
}
