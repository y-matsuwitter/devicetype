package devicetype

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
	return OTHER
}

// IsIPhone returns true if this ua represents iPhone.
func (d *DeviceDetector) IsIPhone() bool {
	return true
}

// IsAndroid returns true if this ua represents Android
func (d *DeviceDetector) IsAndroid() bool {
	return true
}
