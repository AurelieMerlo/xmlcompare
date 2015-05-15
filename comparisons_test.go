package xmlcompare

import (
	"testing"
)

const baseXML = `<human><firstname>chris</firstname><lastname>james</lastname><age>30</age></human>`
const compatibleXML = `<human><firstname>christopher</firstname><lastname>james</lastname><age>15</age></human>`
const differentElementNamesXML = `<wildebeest><name>Blue</name></wildebeest>`
const differentValueTypeXML = `<human><firstname>kristofferson</firstname><lastname>james</lastname><age>old</age></human>`

func TestIdenticalXMLIsCompatible(t *testing.T) {
	assertCompatible(t, baseXML, baseXML)
}

func TestXMLWithSameElementNamesAndValueTypesIsCompatible(t *testing.T) {
	assertCompatible(t, baseXML, compatibleXML)
}

func TestXMLWithDifferentElementNamesIsIncompatible(t *testing.T) {
	assertIncompatible(t, baseXML, differentElementNamesXML)
}

func TestXMLWithDifferentValueTypesIsIncompatible(t *testing.T) {
	assertIncompatible(t, baseXML, differentValueTypeXML)
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
