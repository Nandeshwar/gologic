/*
not related but some video for encrypt using Crypto : https://www.youtube.com/watch?v=jgTqR8PuWuU&t=663s
*/

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
	//uDec, err := b64.URLEncoding.DecodeString("MIIEpAIBAAKCAQEA8mbTZwLtqy5mgA/JFmizW2zZNs62CapuqGP0Xq7SFBMp9CVI\nBWypvAkvuY4ufc1pgpEv3o5gH/e3hXHght2d0VToXsqsvOaVpIPFSmz+OjcF/mhL\n1pMXg9yoqbZMsTSWSTwolFxptJgUhSoOe+m10iGWiQ/fjarbs3cSLJRjK9PtqfIz\n4HymHbAJZY4e80ydtjgvEmLap5C/Pl/B+r0f+UCnuPAVaHLP/zFZRj1vGjSi9ldC\nge+XO0UKRYojTWXzHtLfxH108JHu6Ap2/8ej2o+9MmSGhIfAeAYiQVmDe807L1u4\n5+TqXbaimZHNxbjYCr2pHuphQ205O2cM6TlwxQIDAQABAoIBAC9hXe0Cq4YMOcjK\nTRnuOTCjpAmMehSFlb+gWgGv1ixKWqb4Ko204bB5czSnz/qTDg6RZnPwDGXzO5H/\n4k06QEDMgecVYpJB1/oiL4wOpdqnKqrBpCE6+xuxI96YXYE9sCp8ccoeDiNYnpow\n0Ef0NK/wFMaRrba5MNxxZ8IC9yd3kdg+B4ao6udYKbsRg5v2RS6HgqnJO67y7j+C\niYGoLvu42fNTUX21nQyUwgL2DOalNhVTCh+PNIggLsdl/ZybbXgvyBMo25z2lidZ\nyQrakld6xWQ9G25/PRT635DwwhHtbNwFvA+qlhcvCMXnlglqLaiCw51TyuAa6fKj\nk10AScECgYEA+oRyZ0SDdFjlmR9ofX/Tu23Rwu/VIfXr3as5KRlw2mmPx49ejAsq\nv96x1ROwmDqeF5w5uCS3ytibySl6KA1/mGmnxN9gwCNL+cWffaqYvOcYtDsde5H2\n2zZm0QZma853vw+hbva6qBucqFCyiDfHu/Ad3Y1TnmiHgbsSAuBIv6kCgYEA97To\n4b0tuNuhsFv8K2TOWrFYITUkInKZ3aYpA9UnU+l2zruCVQZTzHS3V0c2eDajR6Jb\nQ2hXu0rJF2P9p9Kk9sDk6id7TAHeQZiRGKfObwDzwpnH1CiCrNDpSeE2lF2vrbx5\nLZOBWwsfNuYFrs/tDS6PZ84IlpJbtBzmPubvCb0CgYBZ5P8ceWRmcqPo/3FCX4U9\n45l6xw8HAbUitRds3Rk29txGMvctb7Bma2YK3OboqVgjhsbbgimFm5bDZ6PTDYz0\nxy0Ro0qXh9LyjOy/bmEioBaoTfI1blpTrUDVzuMf2lXz9IrsQ5MVUds0NsjpwoJk\nuTQuVVFlLYM6lUNQTuUsCQKBgQCsyGJRY/ZKpdkI/YDOfAh8tou0zi6gYLP3Kfoe\nFEbUf1tCJQVqbXlyek+Q77mM7P/D7foe8N+RYz8Vs8exkntDK5YBxvx4Li1sMBG1\n0wdp4o4lxcLfuEo+ZZL018WMhDUQyRD1u5hVe1KQpq58G8lMkpexXsQa04hoAGiS\no99FOQKBgQDOIqRvY2nrMg48mEQx5gaIbaNkRIOKKy3IpqJAf41lSaRwO10sn6uc\nnBOYCx/+7/2OUY/MbRJCk2jX78TxWXDKB1butN7hQDAqdBYKDCUWm9O1JwK8kbRq\nW87mY2Dndj/II1OOVUvbvrD6RjAOnqQ7VlasQ24nyrx4WYE+RINKbQ==")
	str := "MIIEpAIBAAKCAQEA8mbTZwLtqy5mgA/JFmizW2zZNs62CapuqGP0Xq7SFBMp9CVI\nBWypvAkvuY4ufc1pgpEv3o5gH/e3hXHght2d0VToXsqsvOaVpIPFSmz+OjcF/mhL\n1pMXg9yoqbZMsTSWSTwolFxptJgUhSoOe+m10iGWiQ/fjarbs3cSLJRjK9PtqfIz\n4HymHbAJZY4e80ydtjgvEmLap5C/Pl/B+r0f+UCnuPAVaHLP/zFZRj1vGjSi9ldC\nge+XO0UKRYojTWXzHtLfxH108JHu6Ap2/8ej2o+9MmSGhIfAeAYiQVmDe807L1u4\n5+TqXbaimZHNxbjYCr2pHuphQ205O2cM6TlwxQIDAQABAoIBAC9hXe0Cq4YMOcjK\nTRnuOTCjpAmMehSFlb+gWgGv1ixKWqb4Ko204bB5czSnz/qTDg6RZnPwDGXzO5H/\n4k06QEDMgecVYpJB1/oiL4wOpdqnKqrBpCE6+xuxI96YXYE9sCp8ccoeDiNYnpow\n0Ef0NK/wFMaRrba5MNxxZ8IC9yd3kdg+B4ao6udYKbsRg5v2RS6HgqnJO67y7j+C\niYGoLvu42fNTUX21nQyUwgL2DOalNhVTCh+PNIggLsdl/ZybbXgvyBMo25z2lidZ\nyQrakld6xWQ9G25/PRT635DwwhHtbNwFvA+qlhcvCMXnlglqLaiCw51TyuAa6fKj\nk10AScECgYEA+oRyZ0SDdFjlmR9ofX/Tu23Rwu/VIfXr3as5KRlw2mmPx49ejAsq\nv96x1ROwmDqeF5w5uCS3ytibySl6KA1/mGmnxN9gwCNL+cWffaqYvOcYtDsde5H2\n2zZm0QZma853vw+hbva6qBucqFCyiDfHu/Ad3Y1TnmiHgbsSAuBIv6kCgYEA97To\n4b0tuNuhsFv8K2TOWrFYITUkInKZ3aYpA9UnU+l2zruCVQZTzHS3V0c2eDajR6Jb\nQ2hXu0rJF2P9p9Kk9sDk6id7TAHeQZiRGKfObwDzwpnH1CiCrNDpSeE2lF2vrbx5\nLZOBWwsfNuYFrs/tDS6PZ84IlpJbtBzmPubvCb0CgYBZ5P8ceWRmcqPo/3FCX4U9\n45l6xw8HAbUitRds3Rk29txGMvctb7Bma2YK3OboqVgjhsbbgimFm5bDZ6PTDYz0\nxy0Ro0qXh9LyjOy/bmEioBaoTfI1blpTrUDVzuMf2lXz9IrsQ5MVUds0NsjpwoJk\nuTQuVVFlLYM6lUNQTuUsCQKBgQCsyGJRY/ZKpdkI/YDOfAh8tou0zi6gYLP3Kfoe\nFEbUf1tCJQVqbXlyek+Q77mM7P/D7foe8N+RYz8Vs8exkntDK5YBxvx4Li1sMBG1\n0wdp4o4lxcLfuEo+ZZL018WMhDUQyRD1u5hVe1KQpq58G8lMkpexXsQa04hoAGiS\no99FOQKBgQDOIqRvY2nrMg48mEQx5gaIbaNkRIOKKy3IpqJAf41lSaRwO10sn6uc\nnBOYCx/+7/2OUY/MbRJCk2jX78TxWXDKB1butN7hQDAqdBYKDCUWm9O1JwK8kbRq\nW87mY2Dndj/II1OOVUvbvrD6RjAOnqQ7VlasQ24nyrx4WYE+RINKbQ=="

	sDec, err2 := b64.StdEncoding.DecodeString(str)
	if err2 != nil {
		fmt.Println("error b64=", err2)
		return
	}
	fmt.Println("bs64=======", sDec)

	now := time.Now() // current local time
	sec := now.Unix()

	beginDate := "2023-06-05"
	url := fmt.Sprintf("https://cfeeds-content-publisher-api.content-feed.clusters.pluto.tv/content/aero/v1/epg/5dcc42446750e200093b15e2/%s?page=1&results_per_page=100&client_id=slingtv&sling_tv_timestamp=%d", beginDate, sec)
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
