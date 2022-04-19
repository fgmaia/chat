package main

import (
	g "github.com/AllenDang/giu"
)

var (
	mensagem string
	servidor string
	username string
	content  string
)

func loop() {
	g.SingleWindow().Layout(
		g.Row(
			g.Label("Servidor:"),
			g.InputText(&servidor).Size(150),
			g.Label("User name:"),
			g.InputText(&username).Size(150),
			g.Button("conectar").OnClick(func() {
			}),
		),
		g.Row(
			g.Label("Mensagem:"),
			g.InputText(&mensagem).Size(400),
			g.Button("send").OnClick(func() {
			}),
		),
		g.InputTextMultiline(&content).Size(g.Auto, g.Auto),
	)
}

func main() {
	wnd := g.NewMasterWindow("Chat", 600, 600, 0)
	wnd.Run(loop)
}
