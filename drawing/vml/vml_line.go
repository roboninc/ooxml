package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

// line is direct mapping of CT_Line
type line struct {
	XMLName xml.Name `xml:"line"`
	From    string   `xml:"from,attr,omitempty"`
	To      string   `xml:"to,attr,omitempty"`
	shapeAttributes
	shapeElements
}

//Line creates a new object with default values
func Line() *line {
	return &line{
		From: "0,0",
		To:   "10,10",
	}
}

func (s *line) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}