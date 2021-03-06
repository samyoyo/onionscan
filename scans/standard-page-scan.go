package scans

import (
	"github.com/s-rah/onionscan/report"
	"github.com/s-rah/onionscan/utils"
	"log"
	"regexp"
	"strings"
)

func StandardPageScan(scan Scanner, page string, status int, contents string, report *report.OnionScanReport) {
	log.Printf("Scanning %s%s\n", report.HiddenService, page)
	if status == 200 {
		log.Printf("\tPage %s%s is Accessible\n", report.HiddenService, page)

		domains := utils.ExtractDomains(contents)

		for _,domain := range domains {
			if !strings.HasPrefix(domain, "http://"+report.HiddenService) {
				log.Printf("Found Related URL %s\n", domain)
				// TODO: Lots of information here which needs to be processed.
				// * Links to standard sites - google / bitpay etc.
				// * Links to other onion sites
				// * Links to obscure clearnet sites.
			} else {
				// * Process Internal links
			}
		} 

		log.Printf("\tScanning for Images\n")
		r := regexp.MustCompile("src=\"(" + "http://" + report.HiddenService + "/)?((.*?\\.jpg)|(.*?\\.png)|(.*?\\.jpeg)|(.*?\\.gif))\"")
		foundImages := r.FindAllStringSubmatch(string(contents), -1)
		for _, image := range foundImages {
			log.Printf("\t Found image %s\n", image[2])
			scan.ScanPage(report.HiddenService, "/"+image[2], report, CheckExif)
		}
	} else if status == 403 {
		log.Printf("\tPage %s%s is Forbidden\n", report.HiddenService, page)
	} else if status == 404 {
		log.Printf("\tPage %s%s is Does Not Exist\n", report.HiddenService, page)
	}
}
