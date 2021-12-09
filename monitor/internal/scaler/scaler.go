package scaler

import (
	"context"
	"docker_stats/internal/monitor"
	"fmt"
	"os/exec"
	// "strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	cpu_threshold	float64 = 10.0
	mem_threshold	float64 = 10.0
)

func RunContainer(ctx context.Context, container_config *types.ContainerCreateConfig, name string) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	

	resp, err := cli.ContainerCreate(ctx, container_config.Config,
										container_config.HostConfig,
										container_config.NetworkingConfig,
										container_config.Platform,
										name)
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}



}

func CheckThreshold (usage *monitor.ResourceUsage) bool {
	if (usage.Cpu_usage > cpu_threshold) {
		return true
	}
	if (usage.Mem_usage > mem_threshold) {
		return true;
	}

	return false
}

func SpawnNewContainer (ctx context.Context) {
	path := "/home/icecream/Develop/go_dev/assignment2/standalone/docker-compose.yaml"
	fmt.Println(path)
	// exec.Command("/bin/bash", "echo", "ball").Output()
	// cmd := exec.Command("/bin/bash", "docker-compose", "-f", path, "up -d", "--scale go_server=6")
	cmd := exec.Command("/bin/bash", "run.bash", "6")

	// cmd := exec.Command("/bin/bash", "-c", "echo $PWD")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(stdout))
}

