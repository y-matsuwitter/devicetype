/*
Package devicetype detects device type like iPhone, Android, etc... from UserAgent.

Example:
  package main
  import (
    "fmt"

    "github.com/y-matsuwitter/devicetype"
  )

  func main() {
    ua := "Mozilla/5.0 (Linux; Android 4.2.2; HogePhone)"
    dd := devicetype.DeviceType{ua}
    fmt.Println(dd.IsAndroid()) // true
  }
*/
package devicetype
