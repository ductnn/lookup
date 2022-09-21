package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/ammario/ipisp/v2"
	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/olekukonko/tablewriter"
)

//Find CNAME
func get_cname(name string) {
	var domainName string
	color.HiGreen("\n[✔] " + color.HiYellowString("CNAME"))
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
	color.HiGreen("\n[✔] " + color.HiYellowString("TXT records"))
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
	color.HiGreen("\n[✔] " + color.HiYellowString("IP Informations"))
	iprecords, _ := net.LookupIP(name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Domain", "IP", "City and Country", "Location", "Organization", "IPISP",
	})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	domainName = get_domain(name)

	for _, ip := range iprecords {
		// Get IPISP
		resp, err := ipisp.LookupIP(context.Background(), net.ParseIP(ip.String()))
		if err != nil {
			panic(err)
		}

		resp, err = ipisp.LookupASN(context.Background(), ipisp.ASN(resp.ASN))
		if err != nil {
			panic(err)
		}

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
			color.HiCyanString(resp.ISPName),
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
	color.HiGreen("\n[✔] " + color.HiYellowString("Name Servers"))
	nss, _ := net.LookupNS(name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "Name Servers"})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	domainName = get_domain(name)

	for _, ns := range nss {
		data := [][]string{
			{domainName, color.HiRedString(ns.Host)},
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
	color.HiGreen("\n[✔] " + color.HiYellowString("MX Records"))
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

// Function get domain
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

// End
func end_program() {
	var goodbye, githubURL, dockerURL string

	goodbye = "Star the project on GitHub if you liked this tool"
	githubURL = "https://github.com/ductnn/lookup"
	dockerURL = "https://hub.docker.com/r/ductn4/loo"

	color.HiWhite("\n🐳 You can pull docker image in: " + color.HiCyanString(dockerURL))
	color.HiGreen("\n⭐️ " + goodbye)
	color.HiYellow("\n👉 " + githubURL + " 👈")
	color.HiWhite("\n🎉 Thank you so much 🎉")

	fmt.Printf("\n")
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

	color.HiGreen("%s", banner)
	color.HiBlue("\nEnter subdomain or domain name:")

	fmt.Scanf("%s", &name)

	get_cname(name)
	get_txt_record(name)
	get_ip(name)
	get_ns(name)
	get_mx_record(name)

	end_program()
}
