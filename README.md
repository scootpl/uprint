# uPrint

uPrint (microPrint) is a simple TinyGo library to print text on OLED display. Drivers implementing 'drivers.Displayer' are supported. Tested on Raspberry Pi Pico and ssd1306.


## Example

```go
	p := uprint.New(&device, "font7x10")
    p.Print("Hello!", 0, 0, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})

	device.Display()
```