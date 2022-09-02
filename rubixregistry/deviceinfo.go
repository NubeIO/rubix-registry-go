package rubixregistry

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type RubixRegistry struct {
	RubixRegistryDir            string
	RubixRegistryDeviceInfoFile string
	FileMode                    int
}

func New() *RubixRegistry {
	rubixRegistry := RubixRegistry{
		RubixRegistryDir:            "/data/rubix-registry",
		RubixRegistryDeviceInfoFile: "/data/rubix-registry/device_info.json",
		FileMode:                    0755,
	}
	return &rubixRegistry
}

func (inst *RubixRegistry) GetDeviceInfo() (*DeviceInfo, error) {
	data, err := os.ReadFile(inst.RubixRegistryDeviceInfoFile)
	if err != nil {
		return nil, err
	}
	deviceInfoDefault := DeviceInfoDefault{}
	err = json.Unmarshal(data, &deviceInfoDefault)
	if err != nil {
		return nil, err
	}
	return &deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo, nil
}

func (inst *RubixRegistry) UpdateDeviceInfo(deviceInfo DeviceInfo) (*DeviceInfo, error) {
	deviceInfoOld, err := inst.GetDeviceInfo()
	if err != nil {
		return nil, err
	}

	deviceInfo.GlobalUUID = deviceInfoOld.GlobalUUID
	deviceInfo.CreatedOn = deviceInfoOld.CreatedOn
	deviceInfo.UpdatedOn = strings.TrimSuffix(time.Now().UTC().Format(time.RFC3339Nano), "Z")

	deviceInfoDefault := DeviceInfoDefault{
		DeviceInfoFirstRecord: DeviceInfoFirstRecord{
			DeviceInfo: deviceInfo,
		},
	}
	deviceInfoDefaultRaw, err := json.Marshal(deviceInfoDefault)
	err = os.WriteFile(inst.RubixRegistryDeviceInfoFile, deviceInfoDefaultRaw, os.FileMode(inst.FileMode))
	if err != nil {
		return nil, err
	}
	return &deviceInfo, nil
}
