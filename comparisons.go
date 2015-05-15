package xmlcompare

import (
	"fmt"
	"reflect"

	"github.com/clbanning/mxj"
)

// IsCompatible checks that two XML strings are structurally the same so that they are compatible. The first string should be your "correct" json, if there are extra fields in B then they will still be seen as compatible
func IsCompatible(a, b string) (compatible bool, err error) {

	aMap, err := mxj.NewMapXml([]byte(a), true)
	if err != nil {
		return
	}
	fmt.Println("aMap: ", aMap)

	bMap, err := mxj.NewMapXml([]byte(b), true)
	if err != nil {
		return
	}
	fmt.Println("bMap: ", bMap)

	return isStructurallyTheSame(aMap, bMap)
}

func isStructurallyTheSame(a, b map[string]interface{}) (compatible bool, err error) {
	if reflect.DeepEqual(a, b) {
		compatible = true
		return
	}
	for keyInA, v := range a {
		fmt.Println("keyInA: ", keyInA)
		fmt.Println("b[keyInA]: ", b[keyInA])

		if b[keyInA] == nil {
			fmt.Println("Element '%s' exists in 1st XML but not in 2nd")
			return
		}
		fmt.Println("v: ", v)

		switch v.(type) {
		case string:
			fmt.Println("v is string")
			if _, isString := b[keyInA].(string); !isString {
				return
			}
		case bool:
			fmt.Println("v is bool")
			if _, isBool := b[keyInA].(bool); !isBool {
				return
			}
		case float64:
			fmt.Println("v is float64")
			if _, isFloat := b[keyInA].(float64); !isFloat {
				return
			}

		case interface{}:
			fmt.Println("v is interface")
			aArr, aIsArray := a[keyInA].([]interface{})
			fmt.Println("aArr: ", aArr)
			fmt.Println("aIsArray: ", aIsArray)

			bArr, bIsArray := b[keyInA].([]interface{})
			fmt.Println("bArr: ", bArr)
			fmt.Println("bIsArray: ", bIsArray)

			if !bIsArray &&
				aIsArray {
				return
			}

			if !bIsArray &&
				!aIsArray {
				fmt.Println("Neither a or b is array. I should do something else but I don't know what!")
				return
			}

			aLeaf, aIsMap := aArr[0].(map[string]interface{})
			bLeaf, bIsMap := bArr[0].(map[string]interface{})

			if aIsMap && bIsMap {
				return isStructurallyTheSame(aLeaf, bLeaf)
			} else if aIsMap && !bIsMap {
				return
			} else if reflect.TypeOf(aArr[0]) != reflect.TypeOf(bArr[0]) {
				return
			}
		default:
			err = fmt.Errorf("Unmatched datatype in XML found, got a %v", reflect.TypeOf(v))
			return
		}
	}
	compatible = true
	return
}
