package compare

type StringSlice []string

func (me StringSlice) Equals(other []string) bool {
	if len(me) != len(other) {
		return false
	}
	for idx, a := range me {
		if other[idx] != a {
			return false
		}
	}
	return true
}
