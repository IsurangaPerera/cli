package soap

import (
	
	"net/http"
	"net/http/cookiejar"
	"crypto/tls"
	"net/url"
	"os"
)

var URL = "https://localhost:9443/services/UserAdmin.UserAdminHttpsSoap11Endpoint/"

func InitiateConnection() (c *http.Client, err error){

	jar, _ := cookiejar.New(nil)
	
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	client := &http.Client{
		Transport: tr,
		Jar: jar,
	}
	
	r, _ := http.NewRequest("POST", URL, nil)
	r.SetBasicAuth("admin", "admin")
	r.Header.Set("Content-Type", "text/xml")
	r.Header.Set("SOAPAction", `"urn:listAllUsers"`)

	u, _ := url.Parse(URL)

	resp, err := client.Do(r)
	
	if err != nil {
		return nil, err
		os.Exit(1)
	}
	
	defer resp.Body.Close()

	defer resp.Body.Close()
	jar.SetCookies(u, resp.Cookies())

	return client, err

}

func  CreateBasicAuthRequest(r* http.Request) {
	
	r.SetBasicAuth("admin", "admin")//need to be changed
}