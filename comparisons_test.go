package xmlcompare

import (
	"testing"
)

const simpleXML = `<human><firstname>chris</firstname><lastname>james</lastname><age>30</age></human>`
const comparableXML = `<human><firstname>christopher</firstname><lastname>james</lastname><age>15</age></human>`
const incomparableXML = `<wildebeest><name>Blue</name></wildebeest>`

func TestItKnowsTheSameXMLIsCompatible(t *testing.T) {
	assertCompatible(t, simpleXML, simpleXML)
}

func assertCompatible(t *testing.T, a, b string) {
	if compatible, err := IsCompatible(a, b); !compatible || err != nil {
		t.Errorf("%s should be compatible with %s (err = %v)", a, b, err)
	}
}

func assertIncompatible(t *testing.T, a, b string) {
	if compatible, err := IsCompatible(a, b); compatible || err != nil {
		t.Errorf("%s should not be compatible with %s (err = %v)", a, b, err)
	}
}
