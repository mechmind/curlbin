package main

import (
	"io"
	"os/exec"
)

const PYGMENTIZE = "pygmentize"

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderTo(dst io.Writer, src io.Reader) error {
	cmd := exec.Command(PYGMENTIZE, "-g")
	cmd.Stdin = src
	cmd.Stdout = dst

	return cmd.Run()
}
