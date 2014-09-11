package devicetype

import (
	"strings"
)

// DeviceType is a const for each device.
type DeviceType int

// Device type constants
const (
	IPHONE DeviceType = iota
	ANDROID
	OTHER
)

// DeviceDetector is a user agent parser.
type DeviceDetector struct {
	UserAgent string
}

// DeviceType represents this ua’s DeviceType code.
func (d *DeviceDetector) DeviceType() DeviceType {
	lower := strings.ToLower(d.UserAgent)
	if strings.Contains(lower, "iphone") {
		return IPHONE
	}

	if strings.Contains(lower, "android") {
		return ANDROID
	}
	return OTHER
}

// IsIPhone returns true if this ua represents iPhone.
func (d *DeviceDetector) IsIPhone() bool {
	return d.DeviceType() == IPHONE
}

// IsAndroid returns true if this ua represents Android
func (d *DeviceDetector) IsAndroid() bool {
	return d.DeviceType() == ANDROID
}
