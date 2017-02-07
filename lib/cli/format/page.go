package format

import (
	"fmt"
	"strings"
)

type Column struct {
	Name   string //data field name
	Label  string //display name
	Format string //format column
}

type Page struct {
	List    string //list field name
	Columns []Column
}

func (p *Page) Head() string {
	heads := []string{}
	for _, col := range p.Columns {
		if col.Label == "" {
			heads = append(heads, col.Name)
		} else {
			heads = append(heads, col.Label)
		}
	}
	return strings.Join(heads, "\t")
}

func (p *Page) Row() string {
	cells := []string{}
	for _, col := range p.Columns {
		cells = append(cells, col.Format)
	}
	return strings.Join(cells, "\t")
}

func (p *Page) Print(page map[string]interface{}) {
	data := page[p.List]
	if data == nil {
		fmt.Println("could not find filed %s in page", p.List)
		return
	}

	list, ok := data.([]interface{})
	if !ok {
		fmt.Println("page's filed %s isn't a slice", p.List)
		return
	}
	fmt.Println(p.Head())
	for _, row := range list {
		cells, ok := row.(map[string]interface{})
		if !ok {
			fmt.Println("page's filed %s's item isn't a map", p.List)
		}

		for _, col := range p.Columns {
			fmt.Printf(col.Format, cells[col.Name])
			fmt.Print("\t")
		}
		fmt.Println("")
	}
}

