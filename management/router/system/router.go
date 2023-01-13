package system

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type SystemHandlerFunc func() (any, error)

func Route(rg *gin.RouterGroup) {
	router := rg.Group("/system")
	router.GET("/host", SystemHandler(hostInfo))
	router.GET("/cpu", SystemHandler(cpuInfo))
	router.GET("/disk", SystemHandler(diskInfo))
	router.GET("/load", SystemHandler(loadInfo))
	router.GET("/memory", SystemHandler(memoryInfo))
	router.GET("/net", SystemHandler(netInfo))
	router.GET("/process", SystemHandler(processInfo))
}

func SystemHandler(f SystemHandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := f()
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err) //nolint
		}
		context.JSON(http.StatusOK, info)
	}
}

func hostInfo() (iface any, err error) {
	h := &Host{}
	h.Info, err = host.Info()
	return h, err
}

func cpuInfo() (iface any, err error) {
	c := &CPU{}
	c.Count, err = cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	c.Info, err = cpu.Info()
	if err != nil {
		return nil, err
	}
	c.Times, err = cpu.Times(true)
	if err != nil {
		return nil, err
	}
	c.Percent, err = cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func diskInfo() (iface any, err error) {
	d := Disk{}
	d.Partitions, err = disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	for _, partition := range d.Partitions {
		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			return nil, err
		}
		d.Usages = append(d.Usages, usageStat)
	}
	return d, err
}

func loadInfo() (iface any, err error) {
	l := &Load{}
	l.Avg, err = load.Avg()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func memoryInfo() (iface any, err error) {
	m := Memory{}
	m.Virtual, err = mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	m.Swap, err = mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func netInfo() (iface any, err error) {
	n := &Net{}
	var ioCounters []net.IOCountersStat
	ioCounters, err = net.IOCounters(false)
	if err != nil {
		return nil, err
	}
	n.IOCounters = append(n.IOCounters, ioCounters...)
	ioCounters, err = net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	n.IOCounters = append(n.IOCounters, ioCounters...)

	n.Interfaces, err = net.Interfaces()
	if err != nil {
		return nil, err
	}

	n.Connections, err = net.Connections("")
	if err != nil {
		return nil, err
	}

	return n, nil
}

func processInfo() (iface any, err error) {
	p := &Process{}
	p.Processes, err = process.Processes()
	if err != nil {
		return nil, err
	}
	return p, nil
}
