package main

import (
	"fmt"
	"strings"
)

/*
	ignore dot in local part of the email
	everything after + including + upto @ in local part of email is ignored
	
*/

func main() {
	emails := []string{"radha@golok.org", "radha.@golok.org", "krishna+gopis@golok.org", "krishna@golok.org", "ram@saketlok.org"}
	// output: UniqueEmails= [radha@golok.org krishna@golok.org ram@saketlok.org]
	uniqueEmails := findUniqueEmails(emails)
	fmt.Println("UniqueEmails=", uniqueEmails)
}

func findUniqueEmails(emails []string) []string {
	var uniqueEmails []string

	m := make(map[string]struct{})

	for _, v := range emails {
		emailsParts := strings.Split(v, "@")
		localPart, domain := emailsParts[0], emailsParts[1]
		localPart = strings.ReplaceAll(localPart, ".", "")
		localPartList := strings.Split(localPart, "+")
		newEmail := localPartList[0] + "@" + domain
		m[newEmail] = struct{}{}
	}

	for key, _ := range m {
		uniqueEmails = append(uniqueEmails, key)
	}
	return uniqueEmails
}
