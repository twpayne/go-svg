//go:generate go run ./internal/cmds/execute-template -data elements.yaml -output elements.gen.go elements.go.tmpl

// Package svg provides convenience methods for creating and writing SVG
// documents.
//
// See https://www.w3.org/TR/SVG2/.
package svg

import (
	"encoding/xml"
	"io"
	"sort"
)

// A Comment is a comment.
type Comment []byte

// MarshallXML implements encoding/xml.Marshaller.MarshalXML.
func (c Comment) MarshalXML(encoder *xml.Encoder, _ xml.StartElement) error {
	return encoder.EncodeToken(xml.Comment(c))
}

// A CharData is literal XML character data.
type CharData []byte

// MarshallXML implements encoding/xml.Marshaller.MarshalXML.
func (c CharData) MarshalXML(encoder *xml.Encoder, _ xml.StartElement) error {
	return encoder.EncodeToken(xml.CharData(c))
}

// WriteTo implements io.WriterTo.WriteTo. It writes encoding/xml.Header and
// then e to w.
func (e *SVGElement) WriteTo(w io.Writer) (int64, error) {
	return e.WriteToIndent(w, "", "")
}

// WriteToIndent writes encoding/xml.Header and then e to w, indenting with
// prefix and indent.
func (e *SVGElement) WriteToIndent(w io.Writer, prefix, indent string) (int64, error) {
	wc := &writeCounter{w: w}
	encoder := xml.NewEncoder(wc)
	encoder.Indent(prefix, indent)
	if err := encoder.Encode(e); err != nil {
		return int64(wc.bytesWritten), err
	}
	return int64(wc.bytesWritten), nil
}

// encodeElement is a helper function to encode a single element with its
// attributes and children.
func encodeElement(encoder *xml.Encoder, name string, attrs map[string]AttrValue, children []Element) error {
	localNames := make([]string, 0, len(attrs))
	for localName := range attrs {
		localNames = append(localNames, localName)
	}
	sort.Strings(localNames)

	xmlAttrs := make([]xml.Attr, 0, len(attrs))
	for _, localName := range localNames {
		value := attrs[localName].String()
		if value == "" {
			continue
		}
		xmlAttr := xml.Attr{
			Name:  xml.Name{Local: localName},
			Value: value,
		}
		xmlAttrs = append(xmlAttrs, xmlAttr)
	}

	startElement := xml.StartElement{
		Name: xml.Name{Local: name},
		Attr: xmlAttrs,
	}
	if err := encoder.EncodeToken(startElement); err != nil {
		return err
	}
	for _, child := range children {
		if err := child.MarshalXML(encoder, xml.StartElement{}); err != nil {
			return err
		}
	}
	if err := encoder.EncodeToken(startElement.End()); err != nil {
		return err
	}

	return nil
}
