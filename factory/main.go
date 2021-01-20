package main

import "fmt"

func main() {
	ak47Gun, _ := getGun("ak47")
	maverickGun, _ := getGun("maverick")
	printDetails(ak47Gun)
	printDetails(maverickGun)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
