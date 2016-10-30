package tools

import (
	"time"
	"github.com/spf13/viper"
	"github.com/fatih/color"
	"log"
)

var toolsLogger = Logger{Title:"TOOLS",Color:color.FgBlue}


type Logger struct {
	Title string
	Color color.Attribute
	Time  time.Time
}

func (block Logger) Now() Logger {
	block.Time = time.Now()
	return block
}

func (block Logger) Print(message string) {
	if viper.GetBool("verbose") {
		colore := color.New(block.Color).SprintFunc()
		msg := "[" + colore(block.Title) + "] " + message
		empty := time.Time{}
		if block.Time != empty {
			msg = msg + " " + colore("+") + colore(time.Since(block.Time))
		}
		log.Println(msg)
	}
}