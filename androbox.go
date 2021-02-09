// Androbox by wahyu kurniawan (wahyu6070)
// written the first time from zero = 04-02-2021 01.44 WIB FROM asus zenfone max pro m1 (X00T) (android mobile)
// 
// latest update
// version=1.0
// version=code=1

package main
import(
"fmt"
"os"
"runtime"
"os/exec"
"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func boost(){
	fmt.Println("Boost CPU")
	
	
	
	}

func device_info(){
	stopp:
	for true {
		CallClear()
		fmt.Println("             Device Info ")
		fmt.Println("  ")
		if _, err := os.Stat("/data/adb/busybox"); err == nil {
    	fmt.Printf("- Busybox Installed\n");  
  		} else {
    	fmt.Printf("File does not exist\n");  
    	}
    	
    	fmt.Println(" ")
    	fmt.Printf("1.Back\n")
    	fmt.Println(" ")
    	fmt.Printf(" Select Menu : ")
    	var input1 int
		fmt.Scanln(&input1)
		switch input1 {
			case 1:
			break stopp
			default:
			break stopp
			}
	
	}
	}
func about(){
	stopp:
	for true {
		CallClear()
		fmt.Println("             About ")
		fmt.Println("  ")
		fmt.Println("Androbox 1.0 (1) Beta")
		fmt.Println("Androbox Lincense GPL2")
    	
    	fmt.Println(" ")
    	fmt.Printf("1.Back\n")
    	fmt.Println(" ")
    	fmt.Printf(" Select Menu : ")
    	var input1 int
		fmt.Scanln(&input1)
		switch input1 {
			case 1:
			break stopp
			default:
			break stopp
			}
	
	}
	
	}
func main(){
	berhentii:
	for true {
		CallClear()
		fmt.Println("      Androbox")
		fmt.Println(" ")
		fmt.Println(" 1.Boost")
		fmt.Println(" 2.Device Info")
		fmt.Println(" 3.About")
		fmt.Println(" 4.Exit")
		fmt.Println(" ")
		fmt.Printf("  Select Menu : ")
		var inputterm int
		fmt.Scanln(&inputterm)
		switch inputterm {
			case 1:
			boost()
			case 2:
			device_info()
			case 3:
			about()
			case 4:
			break berhentii
			default:
			fmt.Println("!!! Please select menu")
			time.Sleep(1 * time.Second)
			
			}
		
		}
	
	
	}
