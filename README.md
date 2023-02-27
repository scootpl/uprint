# uPrint

uPrint (microPrint) is a simple TinyGo library to print text on OLED display. Drivers implementing 'drivers.Displayer' are supported. Tested on Raspberry Pi Pico and ssd1306.


## Example

```go
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 10 * 1024 * 1024,
	})

	time.Sleep(time.Second)

	device := ssd1306.NewSPI(machine.SPI0, machine.Pin(17), machine.Pin(20), machine.Pin(16))
	device.Configure(ssd1306.Config{})

	device.ClearDisplay()

	p := uprint.New(&device, "font7x10")
	p.Print("Hello!", 0, 0, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})

	device.Display()

	for {
		time.Sleep(time.Second)
	}
```

---
fonts downloaded from: https://github.com/Nondzu/ssd1306_font