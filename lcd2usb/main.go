package lcd2usb

import (
	"errors"
	"fmt"

	"github.com/schleibinger/sio"
)

type cmd byte

const (
	cmdPrefix         cmd = 0xfe
	cmdBacklightOn        = 0x42
	cmdBacklightOff       = 0x46
	cmdBrightnes          = 0x99 // brightnes
	cmdContrast           = 0x50 // contrast
	cmdAutoscrollOn       = 0x51
	cmdAutoscrollOff      = 0x52
	cmdClearScreen        = 0x58
	cmdChangeSplash       = 0x40 // Write number of chars for splash
	cmdCursorPosition     = 0x47 // x, y
	cmdHome               = 0x48
	cmdCursorBack         = 0x4c
	cmdCursorForward      = 0x4d
	cmdUnderlineOn        = 0x4a
	cmdUnderlineOff       = 0x4b
	cmdBlockOn            = 0x53
	cmdBlockOff           = 0x54
	cmdBacklightColor     = 0xd0 // r, g, b
	cmdLCDSize            = 0xd1 // cols, rows
	cmdCreateChar         = 0x4e
	cmdSaveChar           = 0xc1
	cmdLoadChar           = 0xc0
)

type Device struct {
	rows uint8
	cols uint8
	p    *sio.Port
}

func Open(name string, rows, cols uint8) (*Device, error) {
	p, err := sio.Open(name, 9600)
	if err != nil {
		return nil, err
	}

	d := &Device{
		rows: rows,
		cols: cols,
		p:    p,
	}

	if err = d.send(cmdLCDSize, cols, rows); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Device) Close() error {
	if d.p != nil {
		err := d.p.Close()
		d.p = nil
		return err
	}
	return nil
}

func (d *Device) Write(buf []byte) (n int, err error) {
	if d.p == nil {
		return 0, errors.New("writing to closed deviced")
	}

	return d.p.Write(buf)
}

func (d *Device) send(c cmd, args ...byte) error {
	_, err := d.Write(append([]byte{byte(cmdPrefix), byte(c)}, args...))
	return err
}

func (d *Device) Backlight(on bool) error {
	if on {
		return d.send(cmdBacklightOn)
	} else {
		return d.send(cmdBacklightOff)
	}
}

func (d *Device) Brightnes(set uint8) error {
	return d.send(cmdBrightnes, set)
}

func (d *Device) Contrast(set uint8) error {
	return d.send(cmdContrast, set)
}

func (d *Device) Autoscroll(on bool) error {
	if on {
		return d.send(cmdAutoscrollOn)
	} else {
		return d.send(cmdAutoscrollOff)
	}
}

func (d *Device) Clear() error {
	return d.send(cmdClearScreen)
}

func (d *Device) ChangeSplash(set []byte) error {
	cells := int(d.rows) * int(d.cols)
	if len(set) > cells {
		return fmt.Errorf("wrong number of characters: expected %d", cells)
	}

	return d.send(cmdChangeSplash, set...)
}

func (d *Device) CursorPosition(x, y uint8) error {
	if x > d.cols || y > d.rows {
		return fmt.Errorf("setting cursor out of bounds")
	}

	return d.send(cmdCursorPosition, x, y)
}

func (d *Device) Home() error {
	return d.send(cmdHome)
}

func (d *Device) CursorBack() error {
	return d.send(cmdCursorBack)
}

func (d *Device) CursorForward() error {
	return d.send(cmdCursorForward)
}

func (d *Device) Underline(on bool) error {
	if on {
		return d.send(cmdUnderlineOn)
	} else {
		return d.send(cmdUnderlineOff)
	}
}

func (d *Device) Block(on bool) error {
	if on {
		return d.send(cmdBlockOn)
	} else {
		return d.send(cmdBlockOff)
	}
}

func (d *Device) Color(r, g, b uint8) error {
	return d.send(cmdBacklightColor, r, g, b)
}
