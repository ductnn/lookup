package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/olekukonko/tablewriter"
)

func get_domain(name string) string {
	hostName := strings.TrimSpace(name)
	hostParts := strings.Split(hostName, ".")

	lengthOfHostParts := len(hostParts)

	if lengthOfHostParts == 3 {
		name = strings.Join([]string{hostParts[1], hostParts[2]}, ".")
	} else if lengthOfHostParts == 4 {
		name = strings.Join([]string{hostParts[2], hostParts[3]}, ".")
	} else if lengthOfHostParts == 5 {
		name = strings.Join([]string{hostParts[3], hostParts[4]}, ".")
	}

	return name
}

//Find CNAME
func get_cname(name string) {
	var domainName string
	color.Yellow("\nCNAME")
	cname, _ := net.LookupCNAME(name)

	domainName = get_domain(name)

	data := [][]string{
		{name, color.HiCyanString(domainName), color.HiCyanString(cname)},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain or Subdomain", "Domain", "CNAME"})
	table.SetRowLine(true)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

//Find txt records
func get_txt_record(name string) {
	var domainName string
	color.Yellow("\nTXT records")
	txtrecords, _ := net.LookupTXT(name)

	domainName = get_domain(name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "TXT Records"})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	for _, txt := range txtrecords {
		data := [][]string{
			{domainName, color.HiMagentaString(txt)},
		}

		for _, v := range data {
			table.Append(v)
		}
	}

	table.Render()
}

//Find IP informations
func get_ip(name string) {
	var domainName string
	color.Yellow("\nIP Informations")
	iprecords, _ := net.LookupIP(name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Domain", "IP", "City and Country", "Location", "Organization",
	})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	domainName = get_domain(name)

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

		data := [][]string{{
			domainName,
			ip.String(),
			color.HiGreenString(city + ", " + country),
			color.HiGreenString(location),
			color.HiGreenString(organization),
		}}

		for _, v := range data {
			table.Append(v)
		}
	}

	table.Render()
}

//Find nameserver(s)
func get_ns(name string) {
	var domainName string
	color.Yellow("\nName Servers")
	nss, _ := net.LookupNS(name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "Name Servers"})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	domainName = get_domain(name)

	for _, ns := range nss {
		data := [][]string{
			{domainName, color.HiYellowString(ns.Host)},
		}
		for _, v := range data {
			table.Append(v)
		}
	}

	table.Render()
}

//Find MX record
func get_mx_record(name string) {
	var domainName string
	color.Yellow("\nMX Records")
	mxrecords, _ := net.LookupMX(name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "MX Records"})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	domainName = get_domain(name)

	for _, mx := range mxrecords {
		s := mx.Pref
		data := [][]string{{
			domainName,
			color.HiCyanString(mx.Host) + " " + color.HiGreenString(strconv.Itoa(int(s))),
		}}
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
	color.Blue("\nEnter subdomain or domain name:")

	fmt.Scanf("%s", &name)

	get_cname(name)
	get_txt_record(name)
	get_ip(name)
	get_ns(name)
	get_mx_record(name)

	fmt.Printf("\n")
}
