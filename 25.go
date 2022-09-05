package task

import (
	"time"
)

// Sleep1 перестает ждать, когда через d канал из time.After вернёт время
func Sleep1(d time.Duration) {
	<-time.After(d)
}

// Sleep2 перестает ждать, когда через d канал C вернёт время
func Sleep2(d time.Duration) {
	<-time.NewTicker(d).C
}

// Sleep3 перестает ждать, когда через d канал C вернёт время
func Sleep3(d time.Duration) {
	<-time.NewTimer(d).C
}
