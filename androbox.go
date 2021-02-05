// Androbox by wahyu kurniawan (wahyu6070)
// written the first time from zero = 04-02-2021 01.44 WIB FROM asus zenfone max pro m1 (X00T) (android mobilr)
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


func main(){
	berhentii:
	for true {
		CallClear()
		fmt.Println("      Androbox")
		fmt.Println(" ")
		fmt.Println("1.Boost")
		fmt.Println("2.About")
		fmt.Println("3.Exit")
		fmt.Println(" ")
		fmt.Println("Select Menu : \n")
		var inputterm int
		fmt.Scanln(&inputterm)
		switch inputterm {
			case 1:
			pwdd := exec.Command("sh", "asw.sh")
			pwdd.Stdout = os.Stdout
			pwdd.Run()
			fmt.Println("lol")
			time.Sleep(3 * time.Second)
			case 2:
			fmt.Println("lol")
			case 3:
			break berhentii
			default:
			fmt.Println("!!! Please select menu")
			time.Sleep(1 * time.Second)
			
			}
		
		}
	
	
	}
