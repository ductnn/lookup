package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var name string
	banner := `
     /\/| __   /\/| __
    |/\/ /_/  |/\/ /_/
      ___     ___
     / _ \   / _ \
    | (_) | | (_) |
     \___/   \___/
 	`

	color.Green("%s", banner)
	color.Blue("What it looks for : ")
	color.Blue(" * CNAME")
	color.Blue(" * TXT Records")
	color.Blue(" * A and AAA Records")
	color.Blue(" * NS Records")
	color.Blue(" * MX Records")
	color.Blue("\nEnter domain name:")

	fmt.Scanf("%s", &name)

	hostName := strings.TrimSpace(name)
	hostParts := strings.Split(hostName, ".")

	lengthOfHostParts := len(hostParts)

	if lengthOfHostParts == 3 {
		name = strings.Join([]string{hostParts[1], hostParts[2]}, ".")
	} else if lengthOfHostParts == 4 {
		name = strings.Join([]string{hostParts[2], hostParts[3]}, ".")
	}

	get_cname(name)
	get_txt_record(name)
	get_aaa_record(name)
	get_ns(name)
	get_mx_record(name)
}

//Find CNAME
func get_cname(name string) {
	color.Yellow("\nCNAME")
	color.Yellow("+-----------------------------------------+")
	cname, _ := net.LookupCNAME(name)
	fmt.Println("[+]", cname)
}

//Find txt records
func get_txt_record(name string) {
	color.Yellow("\nTXT records")
	color.Yellow("+-----------------------------------------+")
	txtrecords, _ := net.LookupTXT(name)

	for _, txt := range txtrecords {
		fmt.Println("[+]", txt)
	}
}

//Find txt A and AAAA Records
func get_aaa_record(name string) {
	color.Yellow("\nA and AAA")
	color.Yellow("+-----------------------------------------+")
	iprecords, _ := net.LookupIP(name)
	for _, ip := range iprecords {
		fmt.Println("[+]", ip)
	}
}

//Find nameserver(s)
func get_ns(name string) {
	color.Yellow("\nName Server")
	color.Yellow("+-----------------------------------------+")
	nss, _ := net.LookupNS(name)
	for _, ns := range nss {
		fmt.Println("[+]", ns)
	}
}

//Find MX record
func get_mx_record(name string) {
	color.Yellow("\nMX")
	color.Yellow("+-----------------------------------------+")
	mxrecords, _ := net.LookupMX(name)
	for _, mx := range mxrecords {
		fmt.Println("[+]", mx.Host, mx.Pref)
	}
}
