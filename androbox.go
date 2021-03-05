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



/// COLOR functions
///
func printRed(input string){
	//Red color
	fmt.Println(string("\033[31m"), input,string("\033[0m"))
	}
func printGreen(input string){
	//green color
	fmt.Println(string("\033[32m"), input,string("\033[0m"))
	}
func printYellow(input string){
	//Red color
	fmt.Println(string("\033[33m"), input,string("\033[0m"))
	}
	
func printBlue(input string){
	//Blue color
	fmt.Println(string("\033[34m"), input,string("\033[0m"))
	}

func printPurple(input string){
	//Purple color
	fmt.Println(string("\033[35m"), input,string("\033[0m"))
	}
func printCyan(input string){
	//cyan color
	fmt.Println(string("\033[36m"), input,string("\033[0m"))
	}

func printWhite(input string){
	//cyan color
	fmt.Println(string("\033[37m"), input,string("\033[0m"))
	}
///

func print(input string){
	fmt.Print(" ", "\033[37m", input, "\n")
	
	}
////

func boost(){
	fmt.Println("Boost CPU")
	
	
	
	}
	
	
	

func device_info(){
	stopp:
	for true {
		CallClear()
		printBlue("             Device Info ")
		print("  ")
		printGreen("••••SYSTEM••••")
		if _, err := os.Stat("/sbin/magisk"); err == nil {
    	print("- Magisk=Installed");  
  		} else {
    	printRed("- Magisk=Not found");  
    	}
		if _, err := os.Stat("/system/xbin/busybox"); err == nil {
    	print("- Busybox=Installed");  
  		} else {
    	printRed("- Busybox=Not found");  
    	}
    	print(" ")
    	printGreen("••••KERNEL••••")
    	if _, err := os.Stat("/data/adb/busybox"); err == nil {
    	print("- Busybox=Installed");  
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
		printGreen("             About ")
		fmt.Println("  ")
		printCyan("Androbox 1.0 (1) Beta")
		print("Androbox Lincense GPL2")
    	print(" ")
    	printGreen("Contributors :")
    	print("• wahyu6070 (wahyu kurniawan) •> Founder")
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
		printGreen("            Androbox")
		fmt.Println(" ")
		fmt.Println(" 1.Boost")
		fmt.Println(" 2.Device Info")
		fmt.Println(" 3.About")
		fmt.Println(" 4.Exit")
		fmt.Println(" ")
		fmt.Print("  Select Menu : ")
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
			print(" ")
			print(" ")
			printRed("!!! Please select menu")
			time.Sleep(1 * time.Second)
			
			}
		
		}
	
	
	}
