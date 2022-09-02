package rubixregistry

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func (inst *RubixRegistry) CreateDeviceInfoIfDoesNotExist() error {
	dirExist := DirExists(inst.RubixRegistryDir)
	if !dirExist {
		if err := os.MkdirAll(inst.RubixRegistryDir, os.FileMode(inst.FileMode)); err != nil {
			panic(err)
		}
	}
	fileExist := FileExists(inst.RubixRegistryDeviceInfoFile)
	if !fileExist {
		deviceInfoDefault := DeviceInfoDefault{}
		currentDate := strings.TrimSuffix(time.Now().UTC().Format(time.RFC3339Nano), "Z")
		deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.GlobalUUID = ShortUUID("glb")
		deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.CreatedOn = currentDate
		deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.UpdatedOn = currentDate
		deviceInfoDefaultRaw, err := json.Marshal(deviceInfoDefault)
		if err != nil {
			return err
		}
		err = os.WriteFile(inst.RubixRegistryDeviceInfoFile, deviceInfoDefaultRaw, os.FileMode(inst.FileMode))
		if err != nil {
			return err
		}
	}
	return nil
}

func DirExists(dirPath string) bool {
	f, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func FileExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

func ShortUUID(prefix ...string) string {
	u := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, u)
	if n != len(u) || err != nil {
		return "-error-uuid-"
	}
	uuid := fmt.Sprintf("%x%x", u[0:4], u[4:6])
	if len(prefix) > 0 {
		uuid = fmt.Sprintf("%s_%s", prefix[0], uuid)
	}
	return uuid
}
