package xmlfmt

import "encoding/xml"

type ListAllMyBucketsResult struct {
	XMLName xml.Name `xml:"ListAllMyBucketsResult"`
	Buckets Buckets  `xml:"Buckets"`
}

type Buckets struct {
	Bucket []Bucket `xml:"Bucket"`
}

type Bucket struct {
	Name         string `xml:"Name"`
	CreationDate string `xml:"CreationDate"`
}
