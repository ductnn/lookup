package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/olekukonko/tablewriter"
)

//Find CNAME
func get_cname(name string) {
	color.Yellow("\nCNAME")
	cname, _ := net.LookupCNAME(name)
	data := [][]string{
		{name, cname},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "CNAME"})
	table.SetRowLine(true)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	// fmt.Println("[+]", cname)
}

//Find txt records
func get_txt_record(name string) {
	color.Yellow("\nTXT records")
	txtrecords, _ := net.LookupTXT(name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "TXT Records"})
	table.SetRowLine(true)

	for _, txt := range txtrecords {
		data := [][]string{
			{name, "[+] " + txt},
		}

		for _, v := range data {
			table.Append(v)
		}
		// fmt.Println("[+]", txt)
	}

	table.Render()
}

//Find txt A and AAAA Records
func get_aaa_record(name string) {
	color.Yellow("\nIP Informations")
	iprecords, _ := net.LookupIP(name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "IP", "IP informations"})
	table.SetRowLine(true)

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

		// fmt.Println(
		// 	"[+]",
		// 	ip,
		// )
		// fmt.Println(
		// 	color.HiGreenString("=>"),
		// 	color.HiCyanString(city),
		// 	color.HiCyanString(country),
		// 	color.HiCyanString(location),
		// 	color.HiCyanString(organization),
		// )

		data := [][]string{
			{name, ip.String(), color.HiCyanString(city + " " + country + " " + location + " " + organization)},
		}

		for _, v := range data {
			table.Append(v)
		}
	}

	table.Render()
}

//Find nameserver(s)
func get_ns(name string) {
	color.Yellow("\nName Server")
	nss, _ := net.LookupNS(name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "Name Servers"})
	table.SetRowLine(true)

	for _, ns := range nss {
		data := [][]string{
			{name, ns.Host},
		}
		for _, v := range data {
			table.Append(v)
		}
		// fmt.Println("[+]", ns)
	}

	table.Render()
}

//Find MX record
func get_mx_record(name string) {
	color.Yellow("\nMX")
	mxrecords, _ := net.LookupMX(name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "MX Records"})
	table.SetRowLine(true)
	for _, mx := range mxrecords {
		// fmt.Println("[+]", mx.Host, mx.Pref)
		var s string
		s = fmt.Sprint(mx.Pref)
		data := [][]string{
			{name, mx.Host + " " + s},
		}
		for _, v := range data {
			table.Append(v)
		}
	}
	table.Render()
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

	// table := tablewriter.NewWriter(os.Stdout)
	// table.Append(get_cname(name))
}
