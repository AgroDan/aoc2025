package reactor

import (
	"fmt"
	"strings"
)

// First, we'll define the datatype:
// a struct that takes a string as a name,
// and a slice of strings as connections

type Device struct {
	Name        string
	Connections []string
}

func NewDevice(line string) *Device {
	// Parse the line to create a new device
	d := &Device{}
	objs := strings.Split(line, " ")
	d.Name = strings.TrimRight(objs[0], ":")
	for _, conn := range objs[1:] {
		d.Connections = append(d.Connections, conn)
	}
	return d
}

func (d *Device) String() string {
	return fmt.Sprintf("Device %s connected to %v", d.Name, d.Connections)
}
