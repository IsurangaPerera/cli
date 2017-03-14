package soap

import (
	
	"net/http"
	"fmt"
	//"log"
	//"net/http/cookiejar"
	//"io/ioutil"
	"crypto/tls"
	//"crypto/x509"
	//"net/url"
)

func GetSessionCookie() {

	/*// Load client cert
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

		// Load CA cert
	caCert, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

		// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()*/

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//transport := &http.Transport{TLSClientConfig: tlsConfig}
	//client := &http.Client{Transport: transport}


	URL := "https://localhost:9443/services/UserAdmin.UserAdminHttpsSoap11Endpoint/"

	r, _ := http.NewRequest("POST", URL, nil)
	r.SetBasicAuth("admin", "admin")
	r.Header.Set("Content-Type", "text/xml")
	r.Header.Set("SOAPAction", `"urn:listAllUsers"`)


	resp, err := client.Do(r)
	
	if err != nil {
		fmt.Println(resp == nil)
		fmt.Println(err)
		return
	}
	
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	a := resp.Cookies()
	fmt.Println(a)
	
}

func  CreateBasicAuthRequest(r* http.Request) {
	
	r.SetBasicAuth("admin", "admin")//need to be changed
}