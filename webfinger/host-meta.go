package webfinger

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/mrbotchi-team/mrbotchi/utils"
)

type (
	hostMeta struct {
		XMLName xml.Name `xml:"XRD"`
		Text    string   `xml:",chardata"`
		Xmlns   string   `xml:"xmlns,attr"`
		Link    struct {
			Text     string `xml:",chardata"`
			Rel      string `xml:"rel,attr"`
			Type     string `xml:"type,attr"`
			Template string `xml:"template,attr"`
		} `xml:"Link"`
	}
	HostMetaHandler struct {
		Host string
	}
)

func newHostmeta(host string) *hostMeta {
	return &hostMeta{
		Xmlns: "http://docs.oasis-open.org/ns/xri/xrd-1.0",
		Link: struct {
			Text     string `xml:",chardata"`
			Rel      string `xml:"rel,attr"`
			Type     string `xml:"type,attr"`
			Template string `xml:"template,attr"`
		}{
			Rel:      "lrdd",
			Type:     "application/xrd+xml",
			Template: fmt.Sprintf("https://%s/.well-known/webfinger?resource={uri}", host),
		},
	}
}

func (h HostMetaHandler) Get(w http.ResponseWriter, r *http.Request) error {
	response := newHostmeta(h.Host)
	body, err := xml.Marshal(response)
	if nil != err {
		return err
	}

	utils.WriteBody(w, body, http.StatusOK, "application/xrd+xml")
	return nil
}
