package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
)

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
	color.Yellow("\nA and AAA & ipInfo")
	color.Yellow("+-----------------------------------------+")
	iprecords, _ := net.LookupIP(name)
	for _, ip := range iprecords {
		city, err := ipinfo.GetIPCity(ip)
		if err != nil {
			log.Fatal(err)
		}

		country, err := ipinfo.GetIPCountry(ip)
		if err != nil {
			log.Fatal(err)
		}

		location, err := ipinfo.GetIPLocation(ip)
		if err != nil {
			log.Fatal(err)
		}

		organization, err := ipinfo.GetIPOrg(ip)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(
			"[+]",
			ip,
		)
		fmt.Println(
			color.HiGreenString("=>"),
			color.HiCyanString(city),
			color.HiCyanString(country),
			color.HiCyanString(location),
			color.HiCyanString(organization),
		)
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
	color.Blue(" * A and AAA Records & ipInfo")
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

	fmt.Printf("\n")
}
