package models

import "encoding/xml"

const BaseDir = "data"

type Bucket struct {
	XMLName          xml.Name `xml:"bucket"`
	Name             string   `xml:"name"`
	CreationTime     string   `xml:"creationTime"`
	LastModifiedTime string   `xml:"lastModifiedTime"`
	Status           string   `xml:"status"`
}

type Buckets struct {
	XMLName xml.Name `xml:"buckets"`
	Buckets []Bucket `xml:"bucket"`
}

type Object struct {
	XMLName      xml.Name `xml:"object"`
	ObjectKey    string   `xml:"objectKey"`
	ContentType  string   `xml:"contentType"`
	Size         string   `xml:"size"`
	LastModified string   `xml:"lastModified"`
}

type Objects struct {
	XMLName xml.Name `xml:"objects"`
	Objects []Object `xml:"object"`
}

type ErrResp struct {
	Code    int    `xml:"Code"`
	Message string `xml:"Message"`
}
