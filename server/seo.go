package server

import (
	_ "embed"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const baseURL = "https://CHANGEME.example.com"

const robotsTxt = `User-agent: *
Allow: /
Sitemap: %s/sitemap.xml
`

func robotsHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf(robotsTxt, baseURL))
}

// See: https://www.sitemaps.org/protocol.html

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	LastMod    string   `xml:"lastmod,omitempty"`
	ChangeFreq string   `xml:"changefreq,omitempty"`
	Priority   float64  `xml:"priority,omitempty"`
}

func sitemapHandler(c echo.Context) error {
	set := URLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs: []URL{
			{Loc: baseURL + "/users"},
		},
	}
	return c.XML(http.StatusOK, set)
}
