package svg

import "encoding/xml"

type Element interface {
	xml.Marshaler
}
