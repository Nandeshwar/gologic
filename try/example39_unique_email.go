package main

import (
	"fmt"
	"strings"
)

func main() {
	emailList := []string{"nks.sah@yahoo.com","nks.sah+abc@yahoo.com", "nanshu@gmail.com"}

	//uniqueEmails := make(map[string]struct{})
	// line above and a line below - both are same
	uniqueEmails := map[string]struct{}{}
	for _, email := range emailList {
		fixedEmail := fixEmail(email)
		uniqueEmails[fixedEmail] = struct{}{}
	}
	fmt.Println(uniqueEmails)

}

func fixEmail(email string) string {
	emailPartList := strings.Split(email, "@")
	p1 := strings.ReplaceAll(emailPartList[0], ".", "")
	indexPlus := strings.Index(p1, "+")
	if indexPlus >= 0 {
		p1 = p1[0:indexPlus]
	}
	newEmail := p1 + "@" + emailPartList[1];
	return newEmail
}