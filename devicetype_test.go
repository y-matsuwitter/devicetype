package devicetype

import (
	"testing"
)

type deviceTypeResult struct {
	in  string
	out DeviceType
}

var deviceTypeResults = []deviceTypeResult{
	{"iPhone; iOS 7.1.1; Scale/2.00", IPHONE},
	{"Mozilla/5.0 (Linux; Android 4.2.2; HogePhone)", ANDROID},
	{"-", OTHER},
}

func TestDeviceType(t *testing.T) {
	for _, result := range deviceTypeResults {
		dd := DeviceDetector{result.in}
		if dd.DeviceType() != result.out {
			t.Errorf("DeviceType should return %t, got %t", result.out, dd.DeviceType())
		}
	}
}
