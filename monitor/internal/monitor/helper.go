package monitor

import (
	"context"
	"encoding/json"
	// "fmt"
	"strings"

	// "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	// "github.com/docker/docker/api/types/container"
)

func ExtractStats (raw_stats []byte) *StatsData {
	stats := &StatsData{}
	
	err := json.Unmarshal(raw_stats, stats)
	if err != nil {
		panic(err)
	}
	
	return stats
}

func IsGoServerContainer (ctx context.Context, cli *client.Client, id string) bool {
	instance, _ := cli.ContainerInspect(ctx, id)
	// fmt.Println(instance.Name)
	if strings.Contains(instance.Name, "go_server") {
		return true
	}
	return false
}
func CalcPercentCPU (stats *StatsData) float64 {
	cpu_delta := stats.Cpu_stats.Cpu_usage.Total_usage - stats.Precpu_stats.Cpu_usage.Total_usage
	system_cpu_delta := stats.Cpu_stats.System_cpu_usage - stats.Precpu_stats.System_cpu_usage
	number_cpus := stats.Cpu_stats.Online_cpus

	percentage_cpu := float64(cpu_delta)/float64(system_cpu_delta) * float64(number_cpus) * 100.0

	return percentage_cpu
}

func CalcPercentMemory (stats *StatsData) float64 {
	used_memory := stats.Memory_stats.Usage
	available_memory := stats.Memory_stats.Limit

	percentage_memory := float64(used_memory)/float64(available_memory) * 100.0

	return percentage_memory
}
