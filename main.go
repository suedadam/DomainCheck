 package main

 import (
         "fmt"
         "net"
         "os"
         "time"
 )

 func main() {
 		 var toTry []string
         addresses, err := net.LookupIP(os.Args[1])

         for _, addr := range addresses {
         	toTry = append(toTry, addr.String()+":25565")
         }

         _, srvs, err := net.LookupSRV("minecraft", "tcp", os.Args[1])
         if err == nil && len(srvs) > 0 {
         	for _, srv := range srvs {
         		srvaddresses, _ := net.LookupIP(srv.Target)
         		for _, srvaddr := range srvaddresses {
         			toTry = append(toTry, fmt.Sprintf("%s:%d", srvaddr.String(), srv.Port))
         		}
         	}
         }
         check(toTry)
 }

func check(addresses []string) {
	for _, addr := range addresses {
		conn, err := net.DialTimeout("tcp", addr, time.Duration(time.Second))
		if err != nil {
			fmt.Printf("%s isn't online!\n", addr)
			continue
		} else {
			fmt.Printf("%s is online!\n", addr)
			continue
		}
		conn.Close()
	}
}