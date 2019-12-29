package html

type Table struct {
	headers []Element
	rows    []*Row
}

func (t *Table) AddHeader(text string) {
	t.headers = append(t.headers, Element{Text: text})
}
func (t *Table) AddHeaderWithClass(text string, class string) {
	t.headers = append(t.headers, Element{Text: text, Class: class})
}

func (t *Table) FillHeader(limit int) {
	for len(t.headers) < limit {
		t.AddHeader("")
	}
}

func (t *Table) NewRow() *Row {
	row := &Row{}
	t.rows = append(t.rows, row)
	return row
}

func (t *Table) ToHtml() string {
	buf := "<table>"
	buf += "<thead><tr>"
	buf += "\n"
	for _, v := range t.headers {
		if v.Class != "" {
			buf += "<th class='" + v.Class + "'>" + v.Text + "</th>"
		} else {
			buf += "<th>" + v.Text + "</th>"
		}
		buf += "\n"
	}
	buf += "</tr></thead>\n"

	buf += "<tbody>\n"
	for _, r := range t.rows {
		buf += "<tr>"
		for _, e := range r.elements {
			if e.Class != "" {
				buf += "<" + e.Tag + " class='" + e.Class + "'>" + e.Text + "</" + e.Tag + ">"
			} else {
				buf += "<" + e.Tag + ">" + e.Text + "</" + e.Tag + ">"
			}
		}
		buf += "</tr>\n"
	}
	buf += "</tbody></table>"

	return buf
}
