package logger

import (
	"time"
	"github.com/fatih/color"
	"log"
	"math/rand"
)


var acceptableColors = []color.Attribute{color.FgBlue,color.FgWhite,color.FgGreen}


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
	return Logger{Title:title,Color:getColor()}
}


func getColor() color.Attribute {
	random := rand.Intn(len(acceptableColors))
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
