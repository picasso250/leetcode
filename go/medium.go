package main

import "strings"

func simplifyPath(path string) string {
	a := strings.Split(path, "/")
	b := []string{a[0]}
	for i := 1; i < len(a); i++ {
		f := a[i]
		if f == "" {
			// do nothing
		} else if f == "." {
			// do nothing
		} else if f == ".." {
			if len(b) > 1 {
				b = b[:len(b)-1]
			} // else do nothing
		} else {
			b = append(b, f)
		}
	}
	if len(b) == 1 {
		return "/"
	}
	return strings.Join(b, "/")
}
