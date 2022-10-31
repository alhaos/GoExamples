package main

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildList(m)

	err := m.OutputFileAndClose("report.pdf")
	if err != nil {
		panic(err)
	}

}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/image.png", props.Rect{
					Center:  true,
					Percent: 75,
				})
				if err != nil {
					panic(err)
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("some text can be placed here", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func buildList(m pdf.Maroto) {
	tableHeadling := []string{"Accession", "TestName", "TestResult"}
	contents := [][]string{{"2229802630", "Covid", "Detected"}, {"2229802630", "Covid", "Detected"}}

	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadling, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Center,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurpleColor,
	})
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getTealColor() color.Color {

	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
