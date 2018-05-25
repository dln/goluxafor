package goluxafor

import (
	"github.com/google/gousb"
	"log"
)

type Device struct {
	Desc     *gousb.DeviceDesc
	device   *gousb.Device
	intf     *gousb.Interface
	endpoint *gousb.OutEndpoint
}

const vendorId = gousb.ID(0x4d8)   // Microchip Technology Inc.
const productId = gousb.ID(0xf372) // LUXAFOR FLAG

type Led byte

const (
	LedAll Led = 0xff
	LedA   Led = 0x41
	LedB   Led = 0x42
	Led1   Led = 0x01
	Led2   Led = 0x02
	Led3   Led = 0x03
	Led4   Led = 0x04
	Led5   Led = 0x05
	Led6   Led = 0x06
)

type Wave byte

const (
	Wave1 Wave = 0x01
	Wave2 Wave = 0x02
	Wave3 Wave = 0x03
	Wave4 Wave = 0x04
	Wave5 Wave = 0x05
)

type Pattern byte

const (
	Pattern1 Pattern = 0x01
	Pattern2 Pattern = 0x02
	Pattern3 Pattern = 0x03
	Pattern4 Pattern = 0x04
	Pattern5 Pattern = 0x05
	Pattern6 Pattern = 0x06
	Pattern7 Pattern = 0x07
	Pattern8 Pattern = 0x08
)

func newDevice(dev *gousb.Device) *Device {
	intf, _, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("%s.DefaultInterface(): %v", dev, err)
	}

	ep, err := intf.OutEndpoint(1)
	if err != nil {
		log.Fatalf("%s.OutEndpoint(1): %v", intf, err)
	}

	return &Device{
		Desc:     dev.Desc,
		device:   dev,
		intf:     intf,
		endpoint: ep,
	}
}

func (device *Device) Close() {
	device.intf.Close()
	device.device.Close()
}

func (device *Device) writeCommand(command []byte) error {
	_, err := device.endpoint.Write(command)
	if err != nil {
		log.Printf("Error writing data: %s", err)
	}
	return err
}

func (device *Device) Color(led Led, red uint8, green uint8, blue uint8, fadeTime uint8) error {
	data := []byte{0x01, byte(led), red, green, blue, fadeTime, 0x0, 0x0}
	if fadeTime > 0 {
		data[1] = 0x02
	}
	return device.writeCommand(data)
}

func (device *Device) Strobe(led Led, red uint8, green uint8, blue uint8, speed uint8, repeat uint8) error {
	data := []byte{0x03, byte(led), red, green, blue, speed, 0x0, repeat}
	return device.writeCommand(data)
}

func (device *Device) Wave(wave Wave, red uint8, green uint8, blue uint8, speed uint8, repeat uint8) error {
	data := []byte{0x04, byte(wave), red, green, blue, 0x0, repeat, speed}
	return device.writeCommand(data)
}

func (device *Device) Pattern(pattern Pattern, repeat uint8) error {
	data := []byte{0x06, byte(pattern), repeat, 0x0, 0x0, 0x0, 0x0, 0x0}
	return device.writeCommand(data)
}
