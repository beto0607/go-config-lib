package goconfiglib

import (
	"fmt"
)

func (self *Configs) Print() {
	self.Root.Print()

	for s := range self.Root.Subsections {
		section := self.Root.Subsections[s]
		section.Print()
	}
}

func (self *Section) Print() {
	fmt.Printf("[%s]\n", self.Name)
	for k, v := range self.Values {
		fmt.Printf("\t%s = %s\n", k, v)
	}
}
