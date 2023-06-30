package printer

import (
	"fmt"
)

type Printable interface {
	GetName() string
	GetId() int
}

func Print(u Printable, msg string) {
	fmt.Printf("(%d, %s) -> %s\n", u.GetId(), u.GetName(), msg)
}
