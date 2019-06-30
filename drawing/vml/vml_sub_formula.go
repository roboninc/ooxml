package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//Formulas is direct mapping of CT_Formulas
type Formulas struct {
	XMLName xml.Name  `xml:"formulas"`
	List    []Formula `xml:"f"`
}

//Formula is direct mapping of CT_F
type Formula string

type formula struct {
	Eqn string `xml:"eqn,attr"`
}

func (s *Formulas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s.List) > 0 {
		return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
	}

	return nil
}

func (s Formula) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(formula{Eqn: string(s)}, xml.StartElement{
		Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name),
	})
}

func (s *Formula) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if len(start.Attr) > 0 {
		*s = Formula(start.Attr[0].Value)
	}

	return d.Skip()
}
