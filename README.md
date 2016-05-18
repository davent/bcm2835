Go language package to control the Broadcom BCM 2835 as used in the Raspberry
Pi. Builds on Mike McCauley's C lib with the same name.

The package needs Golang 1.1 since the stable 1.0 release doesn't provide a
stable version of cgo. Needs to be run as sudo.

Example:

```go
package main

import (
	"examples/gpioTest/bcm2835"
	"fmt"
	"time"
)

const PIN = bcm2835.RPI_V2_GPIO_P1_29 // Raspberry Pi V2, pin 29

func main() {
	err := bcm2835.Init() // Initialize the library
	if err != nil {
		fmt.Println(err)
		return
	}
	defer bcm2835.Close()                 // Run close when returning
	bcm2835.GpioFsel(PIN, bcm2835.Output) // Set pin 29 to output

	for i := 0; i < 10; i++ { // Loop 10 times
		bcm2835.GpioSet(PIN) // Set pin 29 high
		time.Sleep(500 * time.Millisecond)
		bcm2835.GpioClr(PIN) // Set pin 29 low
		time.Sleep(500 * time.Millisecond)
	}
}
package main

import (
  "github.com/railsme/bcm2835"
  "time"
  "fmt"
)
```
