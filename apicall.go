package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	now2 := time.Now() // current local time
	sec2 := now2.Unix()

	fmt.Println("sec=", sec2)
	str := "nyrx4WYE+RINKbQ=="

	sDec, err2 := b64.StdEncoding.DecodeString(str)
	if err2 != nil {
		fmt.Println("error b64=", err2)
		return
	}
	fmt.Println("bs64=======", sDec)

	now := time.Now() // current local time
	sec := now.Unix()

	beginDate := "2023-06-05"
	url := fmt.Sprintf("my url", beginDate, sec)
	fmt.Println("url=", url)

	//block, input := pem.Decode([]byte(str))
	//if input != nil {
	//	fmt.Println("error in decoding=")
	//	return
	//}
	//fmt.Println("block=", block)

	//privateKey, err2 := rsa.GenerateKey(rand.Reader, 2048)
	//if err2 != nil {
	//	fmt.Println("error=", err2)
	//	return
	//}

	//hashed := sha256.Sum256([]byte(str))

	//fmt.Println("signature=", string(signature))

	//decryptedBytes, err2 := privateKey.Decrypt(nil, sDec, &rsa.OAEPOptions{Hash: crypto.SHA256})
	//if err2 != nil {
	//	fmt.Println("deencrypt error=", err2)
	//	return
	//}
	//
	//fmt.Println(string(decryptedBytes))

	var parsedKey interface{}
	if parsedKey, err2 = x509.ParsePKCS1PrivateKey(sDec); err2 != nil {
		if parsedKey, err2 = x509.ParsePKCS8PrivateKey(sDec); err2 != nil { // note this returns type `interface{}`
			log.Printf("Unable to parse RSA private key, generating a temp one :%s", err2.Error())
			return
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Printf("Unable to parse RSA private key, generating a temp one : %s", err2.Error())
		return
	}

	fmt.Println("privateKey=", privateKey)

	hash := crypto.Hash(crypto.SHA256).New()
	hash.Write([]byte(url))
	hashed := hash.Sum(nil)

	singedResult, err2 := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err2 != nil {
		fmt.Println("error while signing, ", err2)
		return
	}

	fmt.Println("singedResult=", singedResult)

	sEnc := b64.StdEncoding.EncodeToString(singedResult)
	fmt.Println(sEnc)

	//os.Exit(1)

	client := &http.Client{
		Timeout: time.Minute * 20,
	}

	req, err2 := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-sling-tv-signature", sEnc)

	response, err2 := client.Do(req)

	if err2 != nil {
		fmt.Println("error sending req=", err2)
		return
	}

	if response == nil {
		fmt.Println("response is nil...............")
		return
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("error reading body=", err2)
		return
	}
	//fmt.Println("response.StatusCode=", response.StatusCode)

	defer response.Body.Close()

	result := string([]byte(body))
	fmt.Println(result)

}
