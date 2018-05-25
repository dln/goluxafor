package goluxafor

import (
	"github.com/google/gousb"
	"log"
)

type Luxafor struct {
	ctx *gousb.Context

	Devices []*Device
}

func NewLuxafor() Luxafor {
	ctx := gousb.NewContext()

	devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return desc.Vendor == vendorId && desc.Product == productId
	})

	if err != nil {
		log.Fatalf("OpenDevices(): %v", err)
	}
	if len(devs) == 0 {
		log.Fatalf("no devices found matching VID %s and PID %s", vendorId, productId)
	}

	devices := make([]*Device, len(devs))
	for i, d := range devs {
		d.SetAutoDetach(true)
		log.Printf("Opened device: %s", d.Desc)
		devices[i] = newDevice(d)
	}

	return Luxafor{
		ctx:     ctx,
		Devices: devices,
	}
}

func (l *Luxafor) Close() {
	for _, d := range l.Devices {
		d.Close()
	}

	l.ctx.Close()
}
