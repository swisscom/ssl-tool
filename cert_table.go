package main

import "github.com/alexeyco/simpletable"

func certTable() *simpletable.Table {
	t := simpletable.New()
	t.SetStyle(simpletable.StyleCompactClassic)
	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: boldStyle.Render("Subject")},
			{Align: simpletable.AlignLeft, Text: boldStyle.Render("Issuer")},
		},
	}
	return t
}
