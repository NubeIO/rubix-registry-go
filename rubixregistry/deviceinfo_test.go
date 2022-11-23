package rubixregistry

import (
	"fmt"
	"testing"
)

func Test_GetDeviceInfo(*testing.T) {
	rr := New()
	deviceInfo, err := rr.GetDeviceInfo()
	fmt.Println("err", err)
	fmt.Println("deviceInfo", deviceInfo)
	fmt.Println("deviceInfo.DeviceType", deviceInfo.DeviceType)
}

func Test_UpdateDeviceInfo(*testing.T) {
	rr := New()
	deviceInfo, _ := rr.GetDeviceInfo()
	deviceInfo.DeviceType = "cloud"
	deviceInfo, err := rr.UpdateDeviceInfo(*deviceInfo)
	fmt.Println("err", err)
	fmt.Println("deviceInfo", deviceInfo)
	fmt.Println("deviceInfo.DeviceType", deviceInfo.DeviceType)
}
