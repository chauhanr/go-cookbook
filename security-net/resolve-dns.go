package main

import (
	"flag"
	"log"
	"net"
)

func resolveHostname(hostname string) ([]string,error) {
	ips, err := net.LookupHost(hostname)
	if err != nil {
		log.Printf("Error in resolving the host %s\n", err.Error())
		return []string{}, err
	}
	return ips, nil
}

func resolveIPAddr(ipAddr string) ([]string, error){
	names, err := net.LookupAddr(ipAddr)
	if err != nil {
		return []string{}, err
	}
	return names, nil
}

func resolveNameserver(host string) ([]string, error){
	ns, err := net.LookupNS(host)
	if err != nil{
		return []string{}, nil
	}
	hosts := []string{}
	for _, h := range ns{
		hosts = append(hosts,h.Host)
	}
	return hosts, nil
}

func main(){
	host := flag.String("h", "", "hostname to resolve")
	ipAddr := flag.String("ip", "", "IP address to resolve")
	ns := flag.String("ns", "", "Name server to resolve")

	flag.Parse()

	if *host != ""{
		ips, err := resolveHostname(*host)
		if err != nil {
			log.Fatalf("Unable to resolve the hostname : %s", err.Error())
		}
		for _, ip := range ips{
			log.Printf("Host name %s was resolved to %s\n", *host, ip)
		}
	}
	if *ipAddr != ""{
		hosts, err := resolveIPAddr(*ipAddr)
		if err != nil {
			log.Printf("Error while resolving the IP address : %s", err.Error())
		}
		for _,h :=range hosts {
			log.Printf("Host for ip address %s is %s",*ipAddr, h)
		}
	}

	if *ns != "" {
		hservers, err := resolveNameserver(*ns)
		if err != nil{
			log.Fatalf("Error in resolving name server %s : %s",*ns, err.Error())
		}
		for _, server := range hservers{
			log.Printf("Host for name server %s is %s", *ns, server)
		}
	}

}