package goconfiglib

import "os"

func getBaseDir() string {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if len(baseDir) == 0 {
		baseDir = "~/.config"
	}
	// trim trailing slash if there
	if baseDir[len(baseDir)-1] == '/' {
		baseDir = baseDir[:len(baseDir)-1]
	}
	return baseDir
}

func prependXDGConfigPath(path string) string {
	for len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	baseDir := getBaseDir() + "/"

	return baseDir + path
}
