package main

import (
	"syscall"
	"unsafe"
	"fmt"
)

type SystemPowerStatus struct {
	AcLineStatus        byte
	BatteryFlag         byte
	BatteryLifePercent  byte
	SystemStatusFlag    byte
	BatteryLifeTime     uint32
	BatteryFullLifeTime uint32
}

var (
	kernel32, _       = syscall.LoadLibrary("kernel32.dll")
	getPowerStatus, _ = syscall.GetProcAddress(kernel32, "GetSystemPowerStatus")
)

func getSystemPowerStatus() {
	battery := new(SystemPowerStatus)
	var nargs uintptr = 1
	syscall.Syscall(uintptr(getPowerStatus), nargs, uintptr(unsafe.Pointer(battery)), 0, 0)
	dumpSystemPowerStatus(battery)

}

func dumpSystemPowerStatus(battery *SystemPowerStatus) {
	fmt.Printf("AcLineStatus: %d\n", battery.AcLineStatus)
	fmt.Printf("BatteryFlag: %d\n", battery.BatteryFlag)
	fmt.Printf("BatteryLifePercent: %d\n", battery.BatteryLifePercent)
	fmt.Printf("BatteryLifeTime: %d\n", battery.BatteryLifeTime)
}

func main() {
	getSystemPowerStatus()
}
