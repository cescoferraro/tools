package logger


import (
	"time"
	"github.com/fatih/color"
	"log"
	"math/rand"
)

var acceptableColors  []color.Attribute


func init() {
	renew()
}

func renew() {
	acceptableColors = []color.
	Attribute{
		//color.FgBlack,
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		//color.FgWhite,
		color.BgBlack,
		color.BgRed,
		color.BgGreen,
		color.BgYellow,
		color.BgBlue,
		color.BgMagenta,
		color.BgCyan,
		color.BgWhite,
		color.FgHiBlack,
		color.FgHiRed,
		color.FgHiGreen,
		color.FgHiYellow,
		color.FgHiBlue,
		color.FgHiMagenta,
		color.FgHiCyan,
		color.FgHiWhite,
	}
}

type Logger struct {
	Title string
	Color color.Attribute
	Time  time.Time
}

func (block Logger) Now() Logger {
	block.Time = time.Now()
	return block
}

func New(title string) Logger {
	return Logger{Title:title, Color:getColor()}
}

func NewColor(title string, color color.Attribute) Logger {
	for u, cor := range acceptableColors {
		if color == cor {
			acceptableColors = append(acceptableColors[:u], acceptableColors[u + 1:]...)
		}
	}
	return Logger{Title:title, Color:color}
}



func getColor() color.Attribute {
	random := rand.Intn(len(acceptableColors))
	removeColor := func(random int) {
		acceptableColors = append(acceptableColors[:random], acceptableColors[random + 1:]...)
		if len(acceptableColors) == 0 {
			renew()
		}
	}
	defer removeColor(random)
	return acceptableColors[random]
}

func (block Logger) Print(message string) {
	colore := color.New(block.Color).SprintFunc()
	msg := "[" + colore(block.Title) + "] " + message
	empty := time.Time{}

	if block.Time != empty {
		msg = msg + " " + colore("+") + colore(time.Since(block.Time))
	}

	log.Println(msg)
}
