package name_test

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gimei-cli/name"
)

func ExampleNew() {
	fmt.Println(name.New("島根 哲平", "しまね テッペイ"))
	// Output:
	// &{島根 哲平 シマネ テッペイ しまね てっぺい shimane teppei}
}
