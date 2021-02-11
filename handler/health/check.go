package health

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	// B Byte
	B = 1
	// KB KByte
	KB = 1024 * B
	// MB MByte
	MB = 1024 * KB
	// GB GByte
	GB = 1024 * MB
)

// Check shows `OK` as the ping-pong result
func Check(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

// DiskCheck check the disk usage
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")
	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)
	status := http.StatusOK
	tips := "Ok"

	if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		tips = "Warning"
	} else if usedPercent >= 95 {
		status = http.StatusInternalServerError
		tips = "Critical"
	}
	msg := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
		tips, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+msg)
}

// CPUCheck check the cpu usage
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)
	avg, _ := load.Avg()
	load1 := avg.Load1
	load5 := avg.Load5
	load15 := avg.Load15
	status := http.StatusOK
	tips := "Ok"

	if load5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		tips = "Critical"
	} else if load5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		tips = "Warning"
	}
	msg := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d",
		tips, load1, load5, load15, cores)
	c.String(status, "\n"+msg)
}

// MemCheck check the mem usage
func MemCheck(c *gin.Context) {
	vm, _ := mem.VirtualMemory()
	usedMB := int(vm.Used) / MB
	usedGB := int(vm.Used) / GB
	totalMB := int(vm.Total) / MB
	totalGB := int(vm.Total) / GB
	usedPercent := int(vm.UsedPercent)
	status := http.StatusOK
	tips := "Ok"

	if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		tips = "Warning"
	} else if usedPercent >= 95 {
		status = http.StatusInternalServerError
		tips = "Critical"
	}
	msg := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
		tips, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+msg)
}
