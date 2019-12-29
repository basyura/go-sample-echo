package html

type Row struct {
	elements []Element
}

func (r *Row) Add(ele Element) {
	r.elements = append(r.elements, ele)
}

func (r *Row) AddCell(class string) {
	r.Add(Element{Tag: "td", Class: class})
}

func (r *Row) Count() int {
	return len(r.elements)
}

func (r *Row) FillCell(limit int) {
	for len(r.elements) < limit {
		r.elements = append(r.elements, Element{Tag: "td"})
	}
}
