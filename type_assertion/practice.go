package type_assertion

import "fmt"

type myStr string

func TypeAssertion() {
	var value interface{} = "hello world"
	var (
		i   interface{}
		str myStr = "hi"
	)
	i = str
	x, ok := i.(string)
	if ok {
		fmt.Println("uess")
	} else {
		println("panic will not be")
		fmt.Println("noo")
	}
	fmt.Println(x)

	s := value.(string)
	fmt.Println("value's type is string s:", s)

	s, ok = value.(string)
	if ok {
		fmt.Println("value's type is string and ok is true s:", s, "ok:", ok)
	}

	f, ok := value.(float64)
	if !ok {
		fmt.Println("value's type is not float64, f:", f, "ok:", ok)
	}

	fmt.Println("here will be panic")
	f = value.(float64)
	fmt.Println("f:", f)
}
