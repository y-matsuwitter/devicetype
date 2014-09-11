package devicetype

import (
	"fmt"
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

func BenchmarkIsIPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dd := DeviceDetector{"iPhone"}
		dd.IsIPhone()
	}
}

func ExampleIsIPhone() {
	dd := DeviceDetector{"iPhone"}
	fmt.Println(dd.IsIPhone())
	dd = DeviceDetector{"Android"}
	fmt.Println(dd.IsIPhone())

	// Output:
	// true
	// false
}

type isAndroidResult struct {
	in  string
	out bool
}

var isAndroidResults = []isAndroidResult{
	{"iPhone; iOS 7.1.1; Scale/2.00", false},
	{"Mozilla/5.0 (Linux; Android 4.2.2; HogePhone)", true},
}

func TestIsAndroid(t *testing.T) {
	for _, c := range isAndroidResults {
		dd := DeviceDetector{c.in}
		if dd.IsAndroid() != c.out {
			t.Errorf("IsAndroid should return %t for %s, got %t", c.out, c.in, dd.IsAndroid())
		}
	}
}
