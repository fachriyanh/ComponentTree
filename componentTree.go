package main

import (
	"fmt"
)

type barang struct {
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

func defineItem(name string) *barang {
	brg := &barang{
		name: name,
	}
	return brg
}

func (b *barang) insertPart(partName string) *part {
	p := &part{
		name: partName,
	}

	b.part = append(b.part, p)

	return p
}

func (pa *part) insertPart(partName string) *part {
	p := &part{
		name: partName,
	}

	pa.part = append(pa.part, p)

	return p
}

func (b *barang) insertComponent(compName string) {
	c := &component{
		name: compName,
	}

	b.component = append(b.component, c)
}

func (p *part) insertComponent(compName string) {
	c := &component{
		name: compName,
	}

	p.component = append(p.component, c)
}

func (b *barang) printComponent() {
	for _, comp := range b.component {
		fmt.Println(comp.name)
	}

	for _, part := range b.part {
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
