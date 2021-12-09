package main

import (
	"context"
	"docker_stats/internal/monitor"
	"docker_stats/internal/scaler"
	"fmt"
	"time"

	// "docker_stats/internal/scaler"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	// "github.com/docker/docker/container"
)

func main () {
	fmt.Println("docker")
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}


	ctx := context.Background()
	containers, cli := monitor.ListContainers(ctx, cli)

	for _, instance := range containers {
		if monitor.IsGoServerContainer(ctx, cli, instance.ID) {
			// fmt.Println("is a go server", instance.ID)
			resource := monitor.GetResourceUsage(ctx, instance.ID)
			fmt.Println(resource.Cpu_usage)
			fmt.Println(resource.Mem_usage)
		}

	}

	config := &types.ContainerCreateConfig{}
	config.Config = &container.Config{
		Image: "traefik:v2.0",
		Tty: true,
		}
	
	for i := 0; i > -1; i++ {
		containers, _ = monitor.ListContainers(ctx, cli)
		for _, instance := range containers {
			if monitor.IsGoServerContainer(ctx, cli, instance.ID) {
				resource := monitor.GetResourceUsage(ctx, instance.ID)
				if scaler.CheckThreshold(resource) {
					scaler.SpawnNewContainer(ctx)
				}

			}
		}
		time.Sleep(5)
	}
	// scaler.RunContainer(ctx, config, "trial_container")
	// scaler.SpawnNewContainer(ctx)
	// for 
}
