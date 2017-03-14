package soap

import (
	//"bytes"
	"encoding/xml"
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// A RoundTripper executes a request passing the given req as the SOAP
// envelope body. The HTTP response is then de-serialized onto the resp
// object. Returns error in case an error occurs serializing req, making
// the HTTP request, or de-serializing the response.
type RoundTripper interface {
	RoundTrip(req, resp Message) error
}

// Message is an opaque type used by the RoundTripper to carry XML
// documents for SOAP.
type Message interface{}

// Header is an opaque type used as the SOAP Header element in requests.
//type Header interface{}

// AuthHeader is a Header to be encoded as the SOAP Header element in
// requests, to convey credentials for authentication.
type AuthHeader struct {
	Namespace string `xml:"xmlns:ns,attr"`
	Username  string `xml:"ns:username"`
	Password  string `xml:"ns:password"`
}

// Client is a SOAP client.
type Client struct {
	URL         string              // URL of the server
	Namespace   string              // SOAP Namespace
	Envelope    string              // Optional SOAP Envelope
	Header      Header              // Optional SOAP Header
	ContentType string              // Optional Content-Type (default text/xml)
	Config      *http.Client        // Optional HTTP client
	Pre         func(*http.Request) // Optional hook to modify outbound requests
}

func generateXML(v *Envelope) {

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

func formatRequest(r *http.Request) {
 // Create return string
 var request []string
 // Add the request string
 url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
 request = append(request, url)
 // Add the host
 request = append(request, fmt.Sprintf("Host: %v", r.Host))
 // Loop through headers
 for name, headers := range r.Header {
   name = strings.ToLower(name)
   for _, h := range headers {
     request = append(request, fmt.Sprintf("%v: %v", name, h))
   }
 }
 
 // If this is a POST, add post data
 if r.Method == "POST" {
    r.ParseForm()
    request = append(request, "\n")
    request = append(request, r.Form.Encode())
 } 
  // Return the request as a string
  fmt.Println(strings.Join(request, "\n"))
}


// RoundTrip implements the RoundTripper interface.
func (c *Client) RoundTrip(in, out Message) error {
	GetSessionCookie()
	/*req := &Envelope{
		EnvelopeAttr: c.Envelope,
		NSAttr:       c.Namespace,
		Header:       Header{},
		Body:         Body{Message: in},
	}
	if req.EnvelopeAttr == "" {
		req.EnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
	}
	if req.NSAttr == "" {
		req.NSAttr = "http://org.apache.axis2/xsd"
	}
	var b bytes.Buffer

	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		return err
	}
	ct := c.ContentType
	if ct == "" {
		ct = "text/xml"
	}
	cli := c.Config
	if cli == nil {
		cli = http.DefaultClient
	}
	r, err := http.NewRequest("POST", c.URL, &b)
	if err != nil {
		return err
	}

	//generateXML(req)

	//CreateBasicAuthRequest(r)

	r.Header.Set("Content-Type", ct)
	r.Header.Set("SOAPAction", `"urn:listAllUsers"`)
	r.Header.Set("User-Agent", "Apache-HttpClient/4.1.1 (java 1.5)")
	r.Header.Set("Connection", "Keep-Alive")


	if c.Pre != nil {
		c.Pre(r)
	}

	//formatRequest(r)

	resp, err := cli.Do(r)
	if err != nil {
		return err
	}

	//ac := resp.Cookies()
	//fmt.Println(ac[0])
	
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// read only the first Mb of the body in error case
		limReader := io.LimitReader(resp.Body, 1024*1024)
		body, _ := ioutil.ReadAll(limReader)
		return fmt.Errorf("%q: %q", resp.Status, body)
	}
	return xml.NewDecoder(resp.Body).Decode(out)*/
	return nil
}

// Envelope is a SOAP envelope.
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

// Body is the body of a SOAP envelope.
type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Message Message
}