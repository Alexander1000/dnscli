package models

// Canonicalize returns canonicalized string
func Canonicalize(name string) string {
	if name != "" && name[len(name)-1:] != "." {
		return name + "."
	}
	return name
}

// DeCanonicalize returns not canonicalized string
func DeCanonicalize(name string) string {
	if name != "" && name[len(name)-1:] == "." {
		return name[:len(name)-1]
	}
	return name
}
