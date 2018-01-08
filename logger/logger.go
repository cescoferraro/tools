package logger

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
	jww "github.com/spf13/jwalterweatherman"
)

// Debug TODO: NEEDS COMMENT INFO
var Debug bool = true

func init() {
	jww.SetStdoutThreshold(jww.LevelTrace)
}

// Logger TODO: NEEDS COMMENT INFO
type Logger struct {
	Title string
	Color color.Attribute
	Time  time.Time
}

// Now TODO: NEEDS COMMENT INFO
func (block Logger) Now() Logger {
	block.Time = time.Now()
	return block
}

<<<<<<< HEAD
// ShowLocation TODO: NEEDS COMMENT INFO
=======
>>>>>>> 75b7584b484022a9aabb18fe754d142e3e73e48f
func ShowLocation() {
	Debug = true
	return
}

// New TODO: NEEDS COMMENT INFO
func New(title string) Logger {
	return Logger{Title: title, Color: getColor()}
}

// NewColor TODO: NEEDS COMMENT INFO
func NewColor(title string, color color.Attribute) Logger {
	for u, cor := range acceptableColors {
		if color == cor {
			removeColor(u)
		}
	}
	return Logger{Title: title, Color: color}
}

func removeColor(random int) {
	acceptableColors = append(acceptableColors[:random], acceptableColors[random+1:]...)
	if len(acceptableColors) == 0 {
		renew()
	}
}

func getColor() color.Attribute {
	random := rand.Intn(len(acceptableColors))
	defer removeColor(random)
	return acceptableColors[random]
}

// Print TODO: NEEDS COMMENT INFO
func (block Logger) Print(message string) {
	colore := color.New(block.Color).SprintFunc()
	msg := "[" + colore(block.Title) + "] " + message
	empty := time.Time{}
	if block.Time != empty {
		msg = msg + " " + colore("+") + colore(time.Since(block.Time))
	}
	jww.TRACE.Println(msg)
}
