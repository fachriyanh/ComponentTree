package main

import (
	"fmt"
)

type item struct {
	name      string
	component []*component
	part      []*part
}

type part struct {
	name      string
	component []*component
	part      []*part
}

type component struct {
	name string
}

func defineItem(name string) *item {
	itm := &item{
		name: name,
	}
	return itm
}

func (i *item) insertPart(partName string) *part {
	p := &part{
		name: partName,
	}

	i.part = append(i.part, p)

	return p
}

func (pa *part) insertPart(partName string) *part {
	p := &part{
		name: partName,
	}

	pa.part = append(pa.part, p)

	return p
}

func (i *item) insertComponent(compName string) {
	c := &component{
		name: compName,
	}

	i.component = append(i.component, c)
}

func (p *part) insertComponent(compName string) {
	c := &component{
		name: compName,
	}

	p.component = append(p.component, c)
}

func (i *item) printComponent() {
	for _, comp := range i.component {
		fmt.Println(comp.name)
	}

	for _, part := range i.part {
		if len(part.component) > 0 {
			for _, c := range part.component {
				fmt.Println(c.name)
			}
		}

		if len(part.part) > 0 {
			for _, p := range part.part {
				p.printComponent()
			}
		}
	}
}

func (p *part) printComponent() {
	if len(p.component) > 0 {
		for _, c := range p.component {
			fmt.Println(c.name)
		}
	}

	if len(p.part) > 0 {
		for _, p := range p.part {
			p.printComponent()
		}
	}
}

func main() {
	barang := defineItem("barang")
	part1 := barang.insertPart("part 1")
	part1.insertComponent("component 1")
	part1.insertComponent("component 2")
	part1.insertComponent("component 3")

	part2 := barang.insertPart("part 1")
	partA := part2.insertPart("part A")
	partA.insertComponent("component 4")
	partA.insertComponent("component 5")
	partB := part2.insertPart("part B")
	partB.insertComponent("component 6")
	partB.insertComponent("component 7")
	part2.insertComponent("component 8")
	barang.insertComponent("component 9")

	barang.printComponent()
}
