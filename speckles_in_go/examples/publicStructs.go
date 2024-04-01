package animals

import (
	"fmt"
)

type Dog struct {
	Name string
	Age  uint16
}

func (d Dog) Bark() string {
	fmt.Println("Woof! Woof!")

	return "Woof! Woof!"
}

func main() {
	ivy := Dog{Name: "Ivy", Age: 4}
	oakley := Dog{Name: "Oakley", Age: 4}
	juniper := Dog{Name: "Juniper", Age: 4}

	var ruby Dog = Dog{Name: "Ruby", Age: 4}

	ruby.Bark()
	ivy.Bark()
	oakley.Bark()
	juniper.Bark()
}
