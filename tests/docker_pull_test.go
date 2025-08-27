package tests

import (
	"os/exec"
	"testing"
)

func TestDockerPull(t *testing.T) {
	image := "testcontainers/ryuk:0.12.0"

	cmd := exec.Command("docker", "pull", image)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to pull docker image %s: %v\nOutput: %s", image, err, string(output))
	}

	t.Logf("successfully pulled docker image: %s\nOutput: %s", image, string(output))
}
