package monitor

import (
	// "fmt"
	"io/ioutil"
	"context"

	"github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	// "github.com/docker/docker/pkg/stdcopy"
)

type ResourceUsage struct {
	Cpu_usage float64
	Mem_usage float64
}


func ListContainers (ctx context.Context, cli *client.Client) ([]types.Container, *client.Client) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	/* for _, container := range containers {
		fmt.Println(container.ID)
	} */

	return containers, cli
}

func GetResourceUsage(ctx context.Context, containerID string) *ResourceUsage {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	container_stats, err := cli.ContainerStats(ctx, containerID, false)
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(container_stats.Body)
	if err != nil {
		panic(err)
	}
	container_stats.Body.Close()

	stats := ExtractStats(buf)

	cpu_usage := CalcPercentCPU(stats)
	memory_usage := CalcPercentMemory(stats)

	return &ResourceUsage{Cpu_usage: cpu_usage, Mem_usage: memory_usage}
}
