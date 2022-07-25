package system

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type (
	System struct {
		Host   *Host   `json:"host,omitempty"`
		CPU    *CPU    `json:"cpu,omitempty"`
		Disk   *Disk   `json:"disk,omitempty"`
		Load   *Load   `json:"load,omitempty"`
		Memory *Memory `json:"memory,omitempty"`
	}

	Host struct {
		Info *host.InfoStat `json:"info,omitempty"`
	}

	CPU struct {
		Count   int             `json:"count,omitempty"`
		Info    []cpu.InfoStat  `json:"info,omitempty"`
		Times   []cpu.TimesStat `json:"times,omitempty"`
		Percent []float64       `json:"percent,omitempty"`
	}

	Disk struct {
		Partitions []disk.PartitionStat `json:"partitions,omitempty"`
		Usages     []*disk.UsageStat    `json:"usage,omitempty"`
	}

	Load struct {
		Avg *load.AvgStat `json:"avg,omitempty"`
	}

	Memory struct {
		Virtual *mem.VirtualMemoryStat `json:"virtual,omitempty"`
		Swap    *mem.SwapMemoryStat    `json:"swap,omitempty"`
	}

	Net struct {
		IOCounters    []net.IOCountersStat
		Interfaces    []net.InterfaceStat
		Connections   []net.ConnectionStat
		ProtoCounters []net.ProtoCountersStat
	}

	Process struct {
		Processes []*process.Process
	}
)
