package xmlfmt

import "encoding/xml"

type ListAllMyObjectsResult struct {
	XMLName xml.Name `xml:"ListAllMyObjectsResult"`
	Objects Objects  `xml:"Objects"`
}

type Objects struct {
	Object []Object `xml:"Object"`
}

type Object struct {
	Name         string `xml:"Name"`
	Size         int64  `xml:"Size"`
	ContentType  string `xml:"ContentType"`
	CreationDate string `xml:"CreationDate"`
}
