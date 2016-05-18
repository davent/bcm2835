// Package bcm2835 provides functions for the bcm2835 as used in the Raspberry Pi
package bcm2835

//#include "bcm2835.h"
import "C"
import "errors"

const (
	Low            = 0
	High           = 1
	Input          = 0
	Output         = 1
	RPI_GPIO_P1_03 = 0  /*!< Version 1, Pin P1-03 */
	RPI_GPIO_P1_05 = 1  /*!< Version 1, Pin P1-05 */
	RPI_GPIO_P1_07 = 4  /*!< Version 1, Pin P1-07 */
	RPI_GPIO_P1_08 = 14 /*!< Version 1, Pin P1-08, defaults to alt function 0 UART0_TXD */
	RPI_GPIO_P1_10 = 15 /*!< Version 1, Pin P1-10, defaults to alt function 0 UART0_RXD */
	RPI_GPIO_P1_11 = 17 /*!< Version 1, Pin P1-11 */
	RPI_GPIO_P1_12 = 18 /*!< Version 1, Pin P1-12, can be PWM channel 0 in ALT FUN 5 */
	RPI_GPIO_P1_13 = 21 /*!< Version 1, Pin P1-13 */
	RPI_GPIO_P1_15 = 22 /*!< Version 1, Pin P1-15 */
	RPI_GPIO_P1_16 = 23 /*!< Version 1, Pin P1-16 */
	RPI_GPIO_P1_18 = 24 /*!< Version 1, Pin P1-18 */
	RPI_GPIO_P1_19 = 10 /*!< Version 1, Pin P1-19, MOSI when SPI0 in use */
	RPI_GPIO_P1_21 = 9  /*!< Version 1, Pin P1-21, MISO when SPI0 in use */
	RPI_GPIO_P1_22 = 25 /*!< Version 1, Pin P1-22 */
	RPI_GPIO_P1_23 = 11 /*!< Version 1, Pin P1-23, CLK when SPI0 in use */
	RPI_GPIO_P1_24 = 8  /*!< Version 1, Pin P1-24, CE0 when SPI0 in use */
	RPI_GPIO_P1_26 = 7  /*!< Version 1, Pin P1-26, CE1 when SPI0 in use */

	/* RPi Version 2 */
	RPI_V2_GPIO_P1_03 = 2  /*!< Version 2, Pin P1-03 */
	RPI_V2_GPIO_P1_05 = 3  /*!< Version 2, Pin P1-05 */
	RPI_V2_GPIO_P1_07 = 4  /*!< Version 2, Pin P1-07 */
	RPI_V2_GPIO_P1_08 = 14 /*!< Version 2, Pin P1-08, defaults to alt function 0 UART0_TXD */
	RPI_V2_GPIO_P1_10 = 15 /*!< Version 2, Pin P1-10, defaults to alt function 0 UART0_RXD */
	RPI_V2_GPIO_P1_11 = 17 /*!< Version 2, Pin P1-11 */
	RPI_V2_GPIO_P1_12 = 18 /*!< Version 2, Pin P1-12, can be PWM channel 0 in ALT FUN 5 */
	RPI_V2_GPIO_P1_13 = 27 /*!< Version 2, Pin P1-13 */
	RPI_V2_GPIO_P1_15 = 22 /*!< Version 2, Pin P1-15 */
	RPI_V2_GPIO_P1_16 = 23 /*!< Version 2, Pin P1-16 */
	RPI_V2_GPIO_P1_18 = 24 /*!< Version 2, Pin P1-18 */
	RPI_V2_GPIO_P1_19 = 10 /*!< Version 2, Pin P1-19, MOSI when SPI0 in use */
	RPI_V2_GPIO_P1_21 = 9  /*!< Version 2, Pin P1-21, MISO when SPI0 in use */
	RPI_V2_GPIO_P1_22 = 25 /*!< Version 2, Pin P1-22 */
	RPI_V2_GPIO_P1_23 = 11 /*!< Version 2, Pin P1-23, CLK when SPI0 in use */
	RPI_V2_GPIO_P1_24 = 8  /*!< Version 2, Pin P1-24, CE0 when SPI0 in use */
	RPI_V2_GPIO_P1_26 = 7  /*!< Version 2, Pin P1-26, CE1 when SPI0 in use */
	RPI_V2_GPIO_P1_29 = 5  /*!< Version 2, Pin P1-29 */
	RPI_V2_GPIO_P1_31 = 6  /*!< Version 2, Pin P1-31 */
	RPI_V2_GPIO_P1_32 = 12 /*!< Version 2, Pin P1-32 */
	RPI_V2_GPIO_P1_33 = 13 /*!< Version 2, Pin P1-33 */
	RPI_V2_GPIO_P1_35 = 19 /*!< Version 2, Pin P1-35 */
	RPI_V2_GPIO_P1_36 = 16 /*!< Version 2, Pin P1-36 */
	RPI_V2_GPIO_P1_37 = 26 /*!< Version 2, Pin P1-37 */
	RPI_V2_GPIO_P1_38 = 20 /*!< Version 2, Pin P1-38 */
	RPI_V2_GPIO_P1_40 = 21 /*!< Version 2, Pin P1-40 */

	/* RPi Version 2, new plug P5 */
	RPI_V2_GPIO_P5_03 = 28 /*!< Version 2, Pin P5-03 */
	RPI_V2_GPIO_P5_04 = 29 /*!< Version 2, Pin P5-04 */
	RPI_V2_GPIO_P5_05 = 30 /*!< Version 2, Pin P5-05 */
	RPI_V2_GPIO_P5_06 = 31 /*!< Version 2, Pin P5-06 */

	/* RPi B+ J8 header, also RPi 2 40 pin GPIO header */
	RPI_BPLUS_GPIO_J8_03 = 2  /*!< B+, Pin J8-03 */
	RPI_BPLUS_GPIO_J8_05 = 3  /*!< B+, Pin J8-05 */
	RPI_BPLUS_GPIO_J8_07 = 4  /*!< B+, Pin J8-07 */
	RPI_BPLUS_GPIO_J8_08 = 14 /*!< B+, Pin J8-08, defaults to alt function 0 UART0_TXD */
	RPI_BPLUS_GPIO_J8_10 = 15 /*!< B+, Pin J8-10, defaults to alt function 0 UART0_RXD */
	RPI_BPLUS_GPIO_J8_11 = 17 /*!< B+, Pin J8-11 */
	RPI_BPLUS_GPIO_J8_12 = 18 /*!< B+, Pin J8-12, can be PWM channel 0 in ALT FUN 5 */
	RPI_BPLUS_GPIO_J8_13 = 27 /*!< B+, Pin J8-13 */
	RPI_BPLUS_GPIO_J8_15 = 22 /*!< B+, Pin J8-15 */
	RPI_BPLUS_GPIO_J8_16 = 23 /*!< B+, Pin J8-16 */
	RPI_BPLUS_GPIO_J8_18 = 24 /*!< B+, Pin J8-18 */
	RPI_BPLUS_GPIO_J8_19 = 10 /*!< B+, Pin J8-19, MOSI when SPI0 in use */
	RPI_BPLUS_GPIO_J8_21 = 9  /*!< B+, Pin J8-21, MISO when SPI0 in use */
	RPI_BPLUS_GPIO_J8_22 = 25 /*!< B+, Pin J8-22 */
	RPI_BPLUS_GPIO_J8_23 = 11 /*!< B+, Pin J8-23, CLK when SPI0 in use */
	RPI_BPLUS_GPIO_J8_24 = 8  /*!< B+, Pin J8-24, CE0 when SPI0 in use */
	RPI_BPLUS_GPIO_J8_26 = 7  /*!< B+, Pin J8-26, CE1 when SPI0 in use */
	RPI_BPLUS_GPIO_J8_29 = 5  /*!< B+, Pin J8-29,  */
	RPI_BPLUS_GPIO_J8_31 = 6  /*!< B+, Pin J8-31,  */
	RPI_BPLUS_GPIO_J8_32 = 12 /*!< B+, Pin J8-32,  */
	RPI_BPLUS_GPIO_J8_33 = 13 /*!< B+, Pin J8-33,  */
	RPI_BPLUS_GPIO_J8_35 = 19 /*!< B+, Pin J8-35,  */
	RPI_BPLUS_GPIO_J8_36 = 16 /*!< B+, Pin J8-36,  */
	RPI_BPLUS_GPIO_J8_37 = 26 /*!< B+, Pin J8-37,  */
	RPI_BPLUS_GPIO_J8_38 = 20 /*!< B+, Pin J8-38,  */
	RPI_BPLUS_GPIO_J8_40 = 21 /*!< B+, Pin J8-40,  */
)

// Init initialise the library by opening /dev/mem and getting pointers to the
// internal memory for BCM 2835 device registers. You must call this
// (successfully) before calling any other functions in this library (except
// SetDebug)
func Init() (err error) {
	if C.bcm2835_init() == 0 {
		return errors.New("Init: failed")
	}
	return
}

// Close closes the library, deallocating any allocaterd memory and closing
// /dev/mem
func Close() (err error) {
	if C.bcm2835_close() == 0 {
		return errors.New("Close: failed")
	}
	return
}

// SetDebug sets the debug level of the library.  A value of 1 prevents mapping
// to /dev/mem, and makes the library print out what it would do, rather than
// accessing the GPIO registers.  A value of 0, the default, causes normal
// operation.  Call this before calling Init()
func SetDebug(debug int) {
	C.bcm2835_set_debug(C.uint8_t(debug))
}

// GpioFsel sets the function select register for the given pin, which
// configures the pin as either Input or Output
func GpioFsel(pin, mode int) {
	C.bcm2835_gpio_fsel(C.uint8_t(pin), C.uint8_t(mode))
}

// GpioSet sets the specified pin output to high.
func GpioSet(pin int) {
	C.bcm2835_gpio_set(C.uint8_t(pin))
}

// GpioClr sets the specified pin output to low.
func GpioClr(pin int) {
	C.bcm2835_gpio_clr(C.uint8_t(pin))
}

// GpioLev reads the current level on the specified pin and returns either high
// or low. Works whether or not the pin is an input or an output.
func GpioLev(pin int) int {
	return int(C.bcm2835_gpio_lev(C.uint8_t(pin)))
}

func GpioEds(pin int) int {
	return int(C.bcm2835_gpio_eds(C.uint8_t(pin)))
}

func GpioSetEds(pin int) {
	C.bcm2835_gpio_set_eds(C.uint8_t(pin))
}

func GpioRen(pin int) {
	C.bcm2835_gpio_ren(C.uint8_t(pin))
}

func GpioClrRen(pin int) {
	C.bcm2835_gpio_clr_ren(C.uint8_t(pin))
}

func GpioFen(pin int) {
	C.bcm2835_gpio_fen(C.uint8_t(pin))
}

func GpioClrFen(pin int) {
	C.bcm2835_gpio_clr_fen(C.uint8_t(pin))
}

func GpioHen(pin int) {
	C.bcm2835_gpio_hen(C.uint8_t(pin))
}

func GpioClrHen(pin int) {
	C.bcm2835_gpio_clr_hen(C.uint8_t(pin))
}

func GpioLen(pin int) {
	C.bcm2835_gpio_len(C.uint8_t(pin))
}

func GpioClrLen(pin int) {
	C.bcm2835_gpio_clr_len(C.uint8_t(pin))
}

func GpioAren(pin int) {
	C.bcm2835_gpio_aren(C.uint8_t(pin))
}

func GpioClrAren(pin int) {
	C.bcm2835_gpio_clr_aren(C.uint8_t(pin))
}

func GpioAfen(pin int) {
	C.bcm2835_gpio_afen(C.uint8_t(pin))
}

func GpioClrAfen(pin int) {
	C.bcm2835_gpio_clr_afen(C.uint8_t(pin))
}

func GpioPud(pud int) {
	C.bcm2835_gpio_pud(C.uint8_t(pud))
}

func GpioPudclk(pin, on int) {
	C.bcm2835_gpio_pudclk(C.uint8_t(pin), C.uint8_t(on))
}

func GpioPad(group int) uint32 {
	return uint32(C.bcm2835_gpio_pad(C.uint8_t(group)))
}

func GpioSetPad(group int, control uint32) {
	C.bcm2835_gpio_set_pad(C.uint8_t(group), C.uint32_t(control))
}

/// GpioWrite sets the output state of the specified pin
func GpioWrite(pin, on int) {
	C.bcm2835_gpio_write(C.uint8_t(pin), C.uint8_t(on))
}

func GpioSetPud(pin, pud int) {
	C.bcm2835_gpio_set_pud(C.uint8_t(pin), C.uint8_t(pud))
}

func SpiBegin() {
	C.bcm2835_spi_begin()
}

func SpiEnd() {
	C.bcm2835_spi_end()
}

func SpiSetBitOrder(order int) {
	C.bcm2835_spi_setBitOrder(C.uint8_t(order))
}

func SpiSetClockDivider(divider uint16) {
	C.bcm2835_spi_setClockDivider(C.uint16_t(divider))
}

func SpiSetDataMode(mode int) {
	C.bcm2835_spi_setDataMode(C.uint8_t(mode))
}

func SpiChipSelect(cs int) {
	C.bcm2835_spi_chipSelect(C.uint8_t(cs))
}

func SpiSetChipSelectPolarity(cs, active int) {
	C.bcm2835_spi_setChipSelectPolarity(C.uint8_t(cs), C.uint8_t(active))
}

func SpiTransfer(value int) int {
	return int(C.bcm2835_spi_transfer(C.uint8_t(value)))
}

/*
func SpiTransfernb() {
  C.bcm2835_spi_transfernb(char* tbuf, char* rbuf, uint32_t len)
}

func SpiTransfern() {
  C.bcm2835_spi_transfern(char* buf, uint32_t len)
}
*/
