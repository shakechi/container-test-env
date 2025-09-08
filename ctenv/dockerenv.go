package ctenv

import (
	"os"
	"path/filepath"
	"runtime"
)

func ensureDockerEnv() {
	if runtime.GOOS != "darwin" {
		return
	}
	if os.Getenv("DOCKER_HOST") != "" {
		return
	}
	home, _ := os.UserHomeDir()
	rancherSock := filepath.Join(home, ".rd", "docker.sock")
	if _, err := os.Stat(rancherSock); err == nil {
		_ = os.Setenv("DOCKER_HOST", "unix://"+rancherSock)
	}
}
