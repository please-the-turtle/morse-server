package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/please-the-turtle/morse"
	"github.com/please-the-turtle/morse/wave"
)

type MorseQuery struct {
	Message   string        `form:"m" binding:"required,max=50"`
	Frequency float64       `form:"freq" binding:"omitempty,gte=20.0,lte=15000.0"`
	DotLen    time.Duration `form:"dotlen" binding:"omitempty,gte=20ms,lte=2000ms"`
}

func main() {
	route := gin.Default()
	route.GET("/", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var morseQuery MorseQuery

	if err := c.ShouldBind(&morseQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	tr := morse.NewDefaultTranslator()
	p := morse.Parse(morseQuery.Message, tr)
	conv := wave.DefaultWavConverter()

	if morseQuery.Frequency > 0 {
		conv.Frequency = morseQuery.Frequency
	}

	if morseQuery.DotLen > 0 {
		conv.DotLen = morseQuery.DotLen
	}

	wav, err := conv.Convert(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when trying to get audio morse code",
		})

		return
	}

	c.Header("Content-Disposition", "attachment; filename=morse_message.wav")
	c.Data(http.StatusOK, "application/octet-stream", wav)
}
