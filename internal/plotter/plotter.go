package plotter

import (
	"currency_bot/internal/entities"
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)


func GetPlot(data *entities.CurrencyPlot) (string, error) {
	// Преобразуем данные в формат plotter.XYs
	points := make(plotter.XYs, len(data.DateList))

	for i, d := range data.DateList {
		date, _ := time.Parse("2006-01-02", d)
		points[i].X = float64(date.Unix())
		points[i].Y = data.RatesList[i]
	}

	// Создаем график
	p := plot.New()
	p.Title.Text = "Exchange Rate Over Time"
	p.X.Label.Text = "Dates"
	p.Y.Label.Text = fmt.Sprintf("Rates (%s)", data.RateCode)

	// Настройка отображения времени на оси X
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	p.X.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		ticks := make([]plot.Tick, len(data.DateList))
		for i, d := range data.DateList {
			date, _ := time.Parse("2006-01-02", d)
			ticks[i] = plot.Tick{
				Value: float64(date.Unix()),
				Label: date.Format("2006-01-02"),
			}
		}
		return ticks
	})
	p.X.Tick.Label.Rotation = 0.5 // Угол в радианах

	// Добавление линии графика
	line, err := plotter.NewLine(points)
	if err != nil {
		return "", err
	}
	line.LineStyle.Width = vg.Points(2)
	p.Add(line)

	// Добавление точек графика
	scatter, _ := plotter.NewScatter(points)
	scatter.GlyphStyle.Radius = vg.Points(3)
	p.Add(scatter)

	// Сохраняем график
	file := "exchange_rate.png"
	err = p.Save(8*vg.Inch, 8*vg.Inch, file)
	if err != nil {
		return "", err
	}

	return file, err
}
