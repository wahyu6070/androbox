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
//"runtime"
"os/exec"
"time"
)

///////Base functions


func CallClear() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}


func printRed(input string){
	//green color
	fmt.Pri
func printGreen(input string){
	//green color
	fmt.Println(string("\033[32m"), input,string("\033[0m"))
	}



func print(input string){
	fmt.Print(input, "\n")
	
	}
////

func boost(){
	fmt.Println("Boost CPU")
	
	
	
	}
	
	
	

func device_info(){
	stopp:
	for true {
		CallClear()
		fmt.Println("             Device Info ")
		print("  ")
		print("••••KERNEL••••")
		if _, err := os.Stat("/data/adb/busybox"); err == nil {
    	printg("- Busybox=Installed\n");  
  		} else {
    	printg("- Busybox=Not found");  
    	}
    	print("••••KERNEL••••")
    	if _, err := os.Stat("/data/adb/busybox"); err == nil {
    	print("- Busybox=Installed\n");  
  		} else {
    	print("- Busybox=Not found");  
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
