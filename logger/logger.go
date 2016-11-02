package logger

import (
	"time"
	"github.com/fatih/color"
	"math/rand"
	jww "github.com/spf13/jwalterweatherman"
)

var Debug bool = true
func init() {
	jww.UseTempLogFile("api")
	jww.SetStdoutThreshold(jww.LevelInfo)
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

func ShowLocation()  {
	Debug = true
	return
}

func New(title string) Logger {
	return Logger{Title:title, Color:getColor()}
}

func NewColor(title string, color color.Attribute) Logger {
	for u, cor := range acceptableColors {
		if color == cor {
			removeColor(u)
		}
	}
	return Logger{Title:title, Color:color}
}

func removeColor(random int) {
	acceptableColors = append(acceptableColors[:random], acceptableColors[random + 1:]...)
	if len(acceptableColors) == 0 {
		renew()
	}
}
func getColor() color.Attribute {
	random := rand.Intn(len(acceptableColors))
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
	jww.TRACE.Println(msg)
}
