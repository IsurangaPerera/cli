package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"../wsdlgo"
)

type RoundTripper interface {
	RoundTrip(req, resp Message) error
}

type Envelope struct {
	XMLName      xml.Name `xml:"soapenv:Envelope"`
	EnvelopeAttr string   `xml:"xmlns:soapenv,attr"`
	NSAttr       string   `xml:"xmlns:xsd,attr"`
	Header       Header   `xml:"soapenv:Header"`
	Body         Body
}

type Header struct {
	XMLName xml.Name `xml:"soapenv:Header"`
}

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Message Message
}

// Message is an opaque type used by the RoundTripper to carry XML
// documents for SOAP.
type Message interface{}

// AuthHeader is a Header to be encoded as the SOAP Header element in
// requests, to convey credentials for authentication.
type AuthHeader struct {
	Namespace string `xml:"xmlns:ns,attr"`
	Username  string `xml:"ns:username"`
	Password  string `xml:"ns:password"`
}

type Client struct {
	C *http.Client
}

func GenerateXML(v *Envelope) {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	dir = dir+"/resources"+"/configuration.xml"

	f, err := os.Create(dir)
	fmt.Println(err)
	defer f.Close()

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	_, err = f.Write(output)
}

func (m *Client) OneWayTrip(in Message, action string) error {
	
	c := m.C
	req := GenerateSoapRequest(in)
	

	var b bytes.Buffer

	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		return nil
	}

	r, err := http.NewRequest("POST", URL, &b)
	if err != nil {
		return err
	}

	setHeaders(r, action)

	u, _ := url.Parse(URL)
	r.AddCookie(c.Jar.Cookies(u)[0])

	resp, err := c.Do(r)
	if err != nil {
		return nil	
	}

	defer resp.Body.Close()
	var body []byte
	if resp.StatusCode != http.StatusOK {
		body, _ = ioutil.ReadAll(resp.Body)
	} 

	q := &wsdlgo.Envelope{}
	_ = xml.Unmarshal(body, q)

	fmt.Println(q.Body.Fault.FaultString)
	return nil

}
 
func (m *Client) RoundTrip(in, out Message, action string) error {

	c := m.C
	req := GenerateSoapRequest(in)

	var b bytes.Buffer

	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", URL, &b)
	if err != nil {
		return err
	}

	setHeaders(r, action)
	
	u, _ := url.Parse(URL)
	r.AddCookie(c.Jar.Cookies(u)[0])

	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()
	var body []byte
	if resp.StatusCode != http.StatusOK {
		body, _ = ioutil.ReadAll(resp.Body)
		q := &wsdlgo.Envelope{}
		_ = xml.Unmarshal(body, q)

		fmt.Println(q.Body.Fault.FaultString)

		return nil
	}
	body, _ = ioutil.ReadAll(resp.Body)

	_ = xml.Unmarshal(body, out)

	return nil
}