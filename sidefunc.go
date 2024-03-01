package main

import (
    "fmt"
	"errors"

	"golang.org/x/sys/unix"
)



var TerminalSizeError = errors.New("App: Error while getting terminal size")



type WinSize struct {
    Height  uint16
    Width   uint16
}

// Return WinSize & error.
// WinSize -> Height, Width
func GetTermSize() (WinSize, error) {
    ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
    if err != nil {
        return WinSize{}, TerminalSizeError 
    }
    return WinSize{Height: ws.Row, Width: ws.Col}, nil
}

// Clear current line and go previous one
func ClrLine(lines int) {
    for i := 1; i <= lines; i++ {
        fmt.Printf("\x1b[1A\x1b[2K")
    }
}

// Return rgb escape sequence color string.
// Type : f = foreground,  b = background
func Rgb(r uint8, g uint8, b uint8, Type rune) string {
    typeInt := 38
    if Type == 'b'  { typeInt = 48 }
    return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", typeInt, r, g, b)
}



// Reset all text properties.
var RstClr = "\x1b[0m"
// Make text bold.
var BoldText = "\x1b[1m"
// Make text italic.
var ItalicText = "\x1b[3m"
// Add underline on text
var Underline = "\x1b[4m"

// Log theme struct
type LogTheme struct {
    Bg      string
    Fg      string 
    Tc      string 
}

// All log themes.
//
// Theme : red, blue, green, yellow, black, white
var LogThemes = map[string] LogTheme {
    "red" : {
        Bg: Rgb(200, 0,   25,  'b'),
        Fg: Rgb(230, 230, 220, 'f'),
        Tc: Rgb(235, 75,  75,  'f') + BoldText,
    },
    "blue" : {
        Bg: Rgb(60,  80,  150, 'b'),
        Fg: Rgb(0,   200, 250, 'f'),
        Tc: Rgb(0,   200, 250, 'f'),
    },
    "green" : {
        Bg: Rgb(90,  140, 35,  'b'),
        Fg: Rgb(255, 255, 255, 'f'),
        Tc: Rgb(200, 220, 0,   'f'),
    },
    "yellow" : {
        Bg: Rgb(250, 185, 50,  'b'),
        Fg: Rgb(50,  50,  50,  'f'),
        Tc: Rgb(250, 240, 230, 'f'),
    },
    "black" : {
        Bg: Rgb(40,  50,  55,  'b'),
        Fg: Rgb(240, 210, 190, 'f'),
        Tc: Rgb(245, 245, 240, 'f'),
    },
    "white" : {
        Bg: Rgb(230, 225, 200, 'b'),
        Fg: Rgb(5,   55,  80,  'f'),
        Tc: Rgb(230, 235, 140, 'f') + BoldText,
    },
}

// Print log.
// 
// Theme : red, blue, green, yellow, black, white
func Log(ico string, text string, theme string) {
    t, exists := LogThemes[theme]
    if !exists  { t = LogThemes["yellow"] }
    fmt.Printf("%v%v %v %v%v %v %v\n", t.Bg, t.Fg, ico, RstClr, t.Tc, text, RstClr)
}

// Return log.
// 
// Theme : red, blue, green, yellow, black, white
func Slog(ico string, text string, theme string) string {
    t, exists := LogThemes[theme]
    if !exists  { t = LogThemes["yellow"] }
    return fmt.Sprintf("%v%v %v %v%v %v %v\n", t.Bg, t.Fg, ico, RstClr, t.Tc, text, RstClr)
}
