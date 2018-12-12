package main

import (
	"bytes"
	"strings"
)

// 8/16 colors
var (
	colorDefault      = []byte("\033[39m")
	colorBlack        = []byte("\033[30m")
	colorRed          = []byte("\033[31m")
	colorGreen        = []byte("\033[32m")
	colorYellow       = []byte("\033[33m")
	colorBlue         = []byte("\033[34m")
	colorMagenta      = []byte("\033[35m")
	colorCyan         = []byte("\033[36m")
	colorLightGray    = []byte("\033[37m")
	colorDarkGray     = []byte("\033[90m")
	colorLightRed     = []byte("\033[91m")
	colorLightGreen   = []byte("\033[92m")
	colorLightYellow  = []byte("\033[93m")
	colorLightBlue    = []byte("\033[94m")
	colorLightMagenta = []byte("\033[95m")
	colorLightCyan    = []byte("\033[96m")
	colorWhite        = []byte("\033[97m")
)

var colorReplacer = strings.NewReplacer(
	string(colorDefault), "",
	string(colorBlack), "",
	string(colorRed), "",
	string(colorGreen), "",
	string(colorYellow), "",
	string(colorBlue), "",
	string(colorMagenta), "",
	string(colorCyan), "",
	string(colorLightGray), "",
	string(colorDarkGray), "",
	string(colorLightRed), "",
	string(colorLightGreen), "",
	string(colorLightYellow), "",
	string(colorLightBlue), "",
	string(colorLightMagenta), "",
	string(colorLightCyan), "",
	string(colorWhite), "",
)

type buffer struct {
	buf *bytes.Buffer
}

func (b *buffer) NewBuffer(buf *bytes.Buffer) *buffer {
	return &buffer{
		buf: buf,
	}
}

func (b *buffer) WriteColor(bs []byte, c []byte) *buffer {
	b.buf.Write(c)
	b.buf.Write(bs)
	b.buf.Write(colorDefault)
	return b
}
