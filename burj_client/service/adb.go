package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/alex19861108/burj/burj_center/iris/proto"
)

func Exec(name string) (string, error) {
	log.Println(name)
	ns := strings.Split(name, " ")
	cmd := exec.Command(ns[0], ns[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return string(bs), nil
}

const (
	DEVICE_ROOT_PATH  = "/sdcard/pic/"
	DEVICE_MEDIA_PATH = "/data/media"
)

var (
	ERROR_ZK_SERVER_EMPTY = errors.New("[ERROR]: server list is empty")
	ERROR_MAKE_DIRS       = errors.New("[ERROR]: make dirs")
	ERROR_PUSH_DATA       = errors.New("[ERROR]: push data")
	ERROR_INSTALL_APK     = errors.New("[ERROR]: install apk")
	ERROR_START_APK       = errors.New("[ERROR]: start apk")
)

type ADB struct{}

func (a *ADB) install() (string, error) {
	cmd := `brew install Caskroom/cask/android-platform-tools`
	return Exec(cmd)
}

func (a *ADB) Serials() (serials []string) {
	cmd := `adb devices`
	r, err := Exec(cmd)
	if err != nil {
		panic(err)
		return []string{}
	}
	for _, line := range strings.Split(r, "\n") {
		if line != "" {
			pieces := strings.Split(line, "\t")
			if len(pieces) == 2 {
				serial := pieces[0]
				serials = append(serials, serial)
			}
		}
	}
	return serials
}

func (a *ADB) deviceInfo(serial string) (*proto.Device, error) {
	buf := fmt.Sprintf(`adb -s %s shell cat /system/build.prop | grep product`, serial)
	r, err := Exec(buf)
	if err != nil {
		return &proto.Device{}, err
	}

	var (
		model         string
		brand         string
		name          string
		device        string
		board         string
		abi           string
		abilist       string
		abilist32     string
		abilist64     string
		locale        string
		firstApiLevel string
		manufacturer  string
		cuptsm        string
	)
	for _, line := range strings.Split(r, "\n") {
		if !strings.HasPrefix(line, "#") {
			pieces := strings.Split(line, "=")
			if len(pieces) == 2 {
				key, value := pieces[0], pieces[1]
				switch key {
				case "ro.product.model":
					model = value
				case "ro.product.brand":
					brand = value
				case "ro.product.name":
					name = value
				case "ro.product.device":
					device = value
				case "ro.product.board":
					board = value
				case "ro.product.cpu.abi":
					abi = value
				case "ro.product.cpu.abilist":
					abilist = value
				case "ro.product.cpu.abilist32":
					abilist32 = value
				case "ro.product.cpu.abilist64":
					abilist64 = value
				case "ro.product.locale":
					locale = value
				case "ro.product.first_api_level":
					firstApiLevel = value
				case "ro.product.manufacturer":
					manufacturer = value
				case "ro.product.cuptsm":
					cuptsm = value
				}
			}
		}
	}
	return &proto.Device{
		Model:  model,
		Brand:  brand,
		Name:   name,
		Device: device,
		Board:  board,
		Cpu: &proto.CPU{
			Abi:       abi,
			Abilist:   abilist,
			Abilist32: abilist32,
			Abilist64: abilist64,
		},
		Locale:        locale,
		FirstApiLevel: firstApiLevel,
		Manufacturer:  manufacturer,
		Cuptsm:        cuptsm,
		Serial:        serial,
	}, nil
}

func (a *ADB) deviceMemory(serial string) (*proto.Memory, error) {
	buf := fmt.Sprintf(`adb -s %s shell df`, serial)
	r, err := Exec(buf)
	if err != nil {
		return &proto.Memory{}, err
	}

	var (
		memTotal     string
		memUsed      string
		memAvailable string
	)
	for _, line := range strings.Split(r, "\n") {
		if strings.HasPrefix(line, DEVICE_MEDIA_PATH) {
			pieces := strings.Fields(line)
			if len(pieces) == 6 {
				memTotal, memUsed, memAvailable = pieces[1], pieces[2], pieces[3]
			}
		}
	}
	return &proto.Memory{
		Total:     memTotal,
		Used:      memUsed,
		Available: memAvailable,
	}, nil
}

func (a *ADB) deviceMakeDir(serial string, d string) error {
	buf := fmt.Sprintf(`adb -s %s shell mkdir %s`, serial, d)
	_, err := Exec(buf)
	return err
}

func (a *ADB) devicePushData(serial string, src string, dst string) error {
	buf := fmt.Sprintf(`adb -s %s push %s %s`, serial, src, dst)
	_, err := Exec(buf)
	return err
}

func (a *ADB) deviceRmDir(serial string, d string) error {
	buf := fmt.Sprintf(`adb -s %s shell rm -r %s`, serial, d)
	_, err := Exec(buf)
	return err
}

func (a *ADB) deviceInstallApk(serial string, apk string) error {
	buf := fmt.Sprintf(`adb -s %s install %s`, serial, apk)
	_, err := Exec(buf)
	return err
}

func (a *ADB) deviceUninstallApk(serial string, apk string) error {
	buf := fmt.Sprintf(`adb -s %s uninstall %s`, serial, apk)
	_, err := Exec(buf)
	return err
}

func (a *ADB) deviceStartApk(serial string, pkgName string, clsName string, jid string) error {
	buf := fmt.Sprintf(`adb -s %s shell am start %s/.%s --es jid %s`, serial, pkgName, clsName, jid)
	_, err := Exec(buf)
	return err
}

func (a *ADB) Devices() ([]*proto.Device, error) {
	var devices []*proto.Device
	serials := a.Serials()
	for _, serial := range serials {
		device, err := a.deviceInfo(serial)
		if err != nil {
			return devices, err
		}

		memory, err := a.deviceMemory(serial)
		if err != nil {
			return devices, err
		}
		device.Memory = memory
		log.Printf("%#v\n", device)
		devices = append(devices, device)
	}
	return devices, nil
}

func (a *ADB) Run(serial string, dataPath string, apkPath string, apkPkgName string, apkClsName string, jid string) (err error) {
	a.deviceRmDir(serial, DEVICE_ROOT_PATH)
	err = a.deviceMakeDir(serial, DEVICE_ROOT_PATH)
	if err != nil {
		return ERROR_MAKE_DIRS
	}
	err = a.devicePushData(serial, dataPath, DEVICE_ROOT_PATH)
	if err != nil {
		return ERROR_PUSH_DATA
	}
	a.deviceUninstallApk(serial, apkPkgName)
	err = a.deviceInstallApk(serial, apkPath)
	if err != nil {
		return ERROR_INSTALL_APK
	}
	err = a.deviceStartApk(serial, apkPkgName, apkClsName, jid)
	if err != nil {
		return ERROR_START_APK
	}
	return
}
