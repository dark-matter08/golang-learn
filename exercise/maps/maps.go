//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func serverStatus(servers map[string]int) {
	fmt.Println()
	fmt.Println("Server Information")
	fmt.Println("Number of servers:", len(servers))

	stats := make(map[int]int)

	for server, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("Unhandled server status")
		}

		var statusName string
		if status == 0 {
			statusName = "Online"
		}
		if status == 1 {
			statusName = "Offline"
		}
		if status == 2 {
			statusName = "Maintenance"
		}
		if status == 3 {
			statusName = "Retired"
		}
		fmt.Println("==> Server Status:", server, "is ->", statusName)
	}

	fmt.Println("=>", stats[Online], "Servers are Online")
	fmt.Println("=>", stats[Offline], "Servers are Offline")
	fmt.Println("=>", stats[Maintenance], "Servers are under Maintenance")
	fmt.Println("=>", stats[Retired], "Servers are Retired")

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverMap := make(map[string]int)

	for _, server := range servers {
		serverMap[server] = Online
	}

	serverStatus(serverMap)

	serverMap["darkstar"] = Retired
	serverMap["aiur"] = Offline
	serverStatus(serverMap)

	for server := range serverMap {
		serverMap[server] = Maintenance
	}

	serverStatus(serverMap)

}
