package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"session"
	"time"

	"github.com/golang/freetype"
)

func Captcha(w http.ResponseWriter, r *http.Request) {

	const (
		width  = 300
		height = 50
	)

	rectangle := image.Rect(0, 0, width, height)

	img := image.NewNRGBA(rectangle)

	for x := 0; x < width; x++ {

		for y := 0; y < height; y++ {

			img.Set(x, y, color.NRGBA{

				R: uint8(0),
				G: uint8(255),
				B: uint8(0),
				A: uint8(255),
			})

		}

	}

	chars := "abcdefghijklmnopqrstuvwxyz0123456789"

	charlenth := len(chars)

	var code string

	for i := 0; i < 6; i++ {

		rand.Seed(time.Now().UnixNano())

		time.Sleep(50)

		randindex := rand.Intn(charlenth)

		code += string(chars[randindex])

	}

	sessionData := session.Read(w, r)

	sessionData["code"] = code

	session.Write(w, r, sessionData)

	a, err := ioutil.ReadFile("./src/fonts/PRISTINA.TTF")

	if nil != err {

		log.Println(err)
		return
	}

	font, err := freetype.ParseFont(a)

	if nil != err {

		log.Println(err)
		return
	}

	context := freetype.NewContext()

	context.SetFont(font)
	context.SetDPI(float64(72))
	context.SetFontSize(float64(32))
	context.SetDst(img)
	context.SetSrc(image.Black)
	context.SetClip(img.Bounds())

	pt := freetype.Pt(96, 32)

	_, err = context.DrawString(code, pt)

	if nil != err {
		log.Println(err)
		return
	}

	png.Encode(w, img)

}
