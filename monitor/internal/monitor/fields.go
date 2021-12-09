package monitor

type StatsData struct {
	Cpu_stats	CpuStats
	Precpu_stats	CpuStats
	Memory_stats	MemoryStats
}

type CpuStats struct {
	Cpu_usage	CpuUsage
	System_cpu_usage	int64
	Online_cpus	int8
}

type CpuUsage struct {
	Total_usage	int64
}

type MemoryStats struct {
	Usage	int64
	Limit	int64
}
