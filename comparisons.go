package xmlcompare

import (
	"fmt"
	"reflect"

	"github.com/clbanning/mxj"
)

// IsCompatible checks that two XML strings are structurally the same so that they are compatible. The first string should be your "correct" XML, if there are extra elements in B then they will still be seen as compatible.
func IsCompatible(a, b string) (compatible bool, err error) {
	aMap, err := mxj.NewMapXml([]byte(a), true)
	if err != nil {
		return
	}
	bMap, err := mxj.NewMapXml([]byte(b), true)
	if err != nil {
		return
	}
	return isStructurallyTheSame(aMap, bMap)
}

func isStructurallyTheSame(a, b map[string]interface{}) (compatible bool, err error) {
	for keyInA, v := range a {
		switch v.(type) {
		case map[string]interface{}:
			bMap, bIsMap := b[keyInA].(map[string]interface{})
			if bIsMap {
				for vKey, vValue := range v.(map[string]interface{}) {
					if reflect.TypeOf(vValue) != reflect.TypeOf(bMap[vKey]) {
						return
					}
				}
				compatible = true
			}
			return
		default:
			err = fmt.Errorf("Unmatched datatype in XML found, got a %v", reflect.TypeOf(v))
			return
		}
	}
	compatible = true
	return
}
