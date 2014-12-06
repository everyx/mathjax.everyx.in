package main

import (
	"encoding/base64"
	"net/http"
	"os/exec"

	"github.com/go-martini/martini"
)

func mathJax(expressionType string, outType string, expression string) string {
	var mathJaxCmd *exec.Cmd
	if expressionType == "am" {
		if outType == "mml" {
			mathJaxCmd = exec.Command("am2mml", expression)
		} else if outType == "png" {
			mathJaxCmd = exec.Command("am2png", expression)
		} else if outType == "svg" {
			mathJaxCmd = exec.Command("am2svg", expression)
		}
	} else if expressionType == "mml" {
		if outType == "mml" {
			mathJaxCmd = exec.Command("mml2mml", expression)
		} else if outType == "png" {
			mathJaxCmd = exec.Command("mml2png", expression)
		} else if outType == "svg" {
			mathJaxCmd = exec.Command("mml2svg", expression)
		} else if outType == "svg-html5" {
			mathJaxCmd = exec.Command("mml2svg-html5", expression)
		}
	} else if expressionType == "page" {
		if outType == "mml" {
			mathJaxCmd = exec.Command("page2mml", expression)
		} else if outType == "png" {
			mathJaxCmd = exec.Command("page2png", expression)
		} else if outType == "svg" {
			mathJaxCmd = exec.Command("page2svg", expression)
		}
	} else if expressionType == "tex" {
		if outType == "mml" {
			mathJaxCmd = exec.Command("tex2mml", expression)
		} else if outType == "png" {
			mathJaxCmd = exec.Command("tex2png", expression)
		} else if outType == "svg" {
			mathJaxCmd = exec.Command("tex2svg", expression)
		} else if outType == "svg-filter" {
			mathJaxCmd = exec.Command("tex2svg-filter", expression)
		}
	}
	mathJaxOut, err := mathJaxCmd.Output()
	if err != nil {
		panic(err)
	}
	return string(mathJaxOut)
}

func mathJaxHandler(res http.ResponseWriter, params martini.Params) string {
	var contentType string
	switch params["outType"] {
	case "svg":
		contentType = "image/svg+xml"
	case "png":
		contentType = "image/png"
	case "svg-html5":
		contentType = "text/html"
	default:
		contentType = "text/html"
	}
	res.Header().Set("Content-Type", contentType)
	expression, _ := base64.URLEncoding.DecodeString(params["expression"])
	return mathJax(params["expressionType"], params["outType"], string(expression))
}

func main() {
	m := martini.Classic()
	m.Get("/:expressionType/:outType/:expression", mathJaxHandler)
	m.Run()
}
