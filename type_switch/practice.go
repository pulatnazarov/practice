package type_switch

import "fmt"

func TypeSwitch() {
	do("hello")
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("v's type is int and v:", v)
	case float32:
		fmt.Println("v's type is float32 and v:", v)
	case string:
		fmt.Println("v's type is string and v:", v)
	}
}
