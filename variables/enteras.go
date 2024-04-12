package variables

import "fmt"
func ShowIntegers() {
	var intcommon int
	intof32 := int32(10)
	intof64 := int64(100)
	fmt.Println(intcommon, intof32, intof64)
	fmt.Println("intcommon = ", intcommon)
	fmt.Println("intof32 = ", intof32)
	fmt.Println("intof64 = ", intof64)

}