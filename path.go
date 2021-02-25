package cgr

import "path"

// Eliminates . and .. elements
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)

	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}

	return np
}

func appendSlash(path string) string{
	return path + "/"
}