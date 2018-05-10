/*
 How to run
 Pass the Bluetooth name or address as first param:

        go run examples/minidrone_events.go "Mambo_1234"

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
        "fmt"
        "os"
        "time"
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/platforms/ble"
        "gobot.io/x/gobot/platforms/parrot/minidrone"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
        bleAdaptor := ble.NewClientAdaptor(os.Args[1])
        drone := minidrone.NewDriver(bleAdaptor)
	keys := keyboard.NewDriver()

        work := func() {
		
                keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			rotationSpeed:=100
			upDownSpeed:=100
			forwardBackSpeed:=100
			leftRightSpeed:=100
		
			switch key.Key{
				case keyboard.W:
					drone.Up(upDownSpeed)
				case keyboard.S:
					drone.Down(upDownSpeed)
				case keyboard.A:
					drone.Clockwise(rotationSpeed)
				case keyboard.D:
					drone.CounterClockwise(rotationSpeed)
				case keyboard.T:
					drone.TakeOff()
				case keyboard.L:
					drone.Land()
				case keyboard.P:
					drone.TakePicture()
				case keyboard.N: //Enter
					drone.GunControl(0)
				case 32: //Spacebar				
					drone.Stop()
				case 65: //Up Arrow
					drone.Forward(forwardBackSpeed)
				case 66: //Down Arrow
					drone.Backward(forwardBackSpeed)
				case 68: //Left Arrow
					drone.Left(leftRightSpeed)
				case 67: //Right Arrow
					drone.Right(leftRightSpeed)
				case keyboard.C:
					drone.ClawControl(0,1)
				case keyboard.V:
					drone.ClawControl(0,0)					
				case keyboard.Z: 
					drone.FrontFlip()
					drone.Down(10)				
	
			}
		})
		   
		
		drone.On(minidrone.Battery, func(data interface{}) {
                        fmt.Printf("battery: %d\n", data)
                })

                drone.On(minidrone.FlightStatus, func(data interface{}) {
                        fmt.Printf("flight status: %d\n", data)
                })
/*
                drone.On(minidrone.Takeoff, func(data interface{}) {
                        fmt.Println("taking off...")
                })

                drone.On(minidrone.Hovering, func(data interface{}) {
                        fmt.Println("hovering!")
                })
		
                drone.On(minidrone.Landing, func(data interface{}) {
                        fmt.Println("landing...")
                })

                drone.On(minidrone.Landed, func(data interface{}) {
                        fmt.Println("landed.")
                })*/

               time.Sleep(50 * time.Millisecond)
                
        }

	
        robot := gobot.NewRobot("minidrone",
                []gobot.Connection{bleAdaptor},
                []gobot.Device{drone,keys},
                work,
        )


        robot.Start()


}



