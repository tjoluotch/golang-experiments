package main

import (
	"fmt"
)

// go basic type byte is an alias for uint8  - 8 bit unsigned integer values
type IpAddr [4]byte

/*	TODO: Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad
	TODO: For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
 */

 // pointer to ipAddr as the receiver
func(ipAddr *IpAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

 func main(){
 	// map of key - type string and value - type IpAddr
 	hosts := map[string]IpAddr{
 		"loopback": {127, 0, 0, 1},
 		"googleDNS": {8, 8, 8, 8},
	}


 	fmt.Println("How the map key & values looked before using the Stringer interface{ String() string }")
 	fmt.Printf("\n")

	 // loop through the entire map
 	for name, ip := range hosts{
 		fmt.Printf("%v: %v\n", name, ip)
	}

 	 fmt.Printf("\n")
	 fmt.Println("How the map key & values looked after implementing stringer")
	 fmt.Printf("\n")

 	for name, ip := range hosts {
 		// & indicates we pass the address of local variables name and ip
 		fmt.Printf("%v: %v\n", &name, &ip)
	}

 }
