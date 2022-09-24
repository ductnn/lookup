package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ammario/ipisp/v2"
	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/olekukonko/tablewriter"
)

type Crtsr struct {
	CommonName string `json:"common_name"`
	NameValue  string `json:"name_value"`
}

var banner = `
 _                _    _   _
| |    ___   ___ | | _| | | |_ __
| |   / _ \ / _ \| |/ / | | | '_ \
| |__| (_) | (_) |   <| |_| | |_) |
|_____\___/ \___/|_|\_\\___/| .__/
                            |_|

[+] by @ductnn
[+] https://github.com/ductnn
[-] Usage: ./loo
`

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

// Get Subdomain
func GetJsonFromCrt(domain string) ([]string, error) {

	resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	sb := []byte(body)
	var subdomains []Crtsr
	err = json.Unmarshal(sb, &subdomains)
	if err != nil {
		panic(err)
	}

	output := make([]string, 0)
	for _, subdomains := range subdomains {
		output = append(output, subdomains.CommonName)
		output = append(output, subdomains.NameValue)
	}

	return output, nil

}
func removeDuplicateValues(strSlice []string) []string {
	// https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func get_subdomain(name string) {
	var domainName string
	color.HiGreen("\n[✔] " + color.HiYellowString("Subdomain"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Domain", "Subdomain"})
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)

	subdomain, err := GetJsonFromCrt(name)
	if err != nil {
		panic(err)
	}

	removeDuplicateValuesSlice := removeDuplicateValues(subdomain)

	for _, uniquesubdomain := range removeDuplicateValuesSlice {
		// fmt.Println(uniquesubdomain)
		domainName = get_domain(name)

		data := [][]string{
			{color.HiCyanString(domainName), color.HiGreenString(uniquesubdomain)},
		}

		for _, v := range data {
			table.Append(v)
		}
	}

	table.Render()
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

	color.HiGreen("%s", banner)
	color.HiCyan("\nEnter subdomain or domain name:")

	fmt.Scanf("%s", &name)

	var choice int

	color.Cyan("\nEnter your choice")
	fmt.Printf("\n")

	color.Green("[1] - CNAME lookup")
	color.Green("[2] - Subdomain lookup")
	color.Green("[3] - TXT Records lookup")
	color.Green("[4] - IP informstion lookup")
	color.Green("[5] - NameServers lookup")
	color.Green("[6] - MX Records lookup")

	fmt.Printf("\n")
	fmt.Scanf("%d", &choice)

	if choice == 1 {
		get_cname(name)
	} else if choice == 2 {
		get_subdomain(name)
	} else if choice == 3 {
		get_txt_record(name)
	} else if choice == 4 {
		get_ip(name)
	} else if choice == 5 {
		get_ns(name)
	} else if choice == 6 {
		get_mx_record(name)
	}

	end_program()
}
