package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)



type Content struct {
    X           int 
    Y           int
    MainRune    rune
    CombRune    rune
    Style       tcell.Style
}
var ScreenBuffer [][] Content
func Def(names ...string) {
    for _, name := range names {
        fmt.Println(name)
    }
}



var MyScreen tcell.Screen
var MyApp App
var MyTheme = Theme {
    AppTheme: AppTheme{
        BGColor: [3]int32 {30, 20, 50},
        FGColor: [3]int32 {230, 220, 250},
    },
}



func main() {

    // MyApp, MyScreen, err := NewApp(&MyTheme)
    MyApp, _, err := NewApp(&MyTheme)
    if err != nil {
        MyApp.Exit(1)
    }


    MyApp.CreateActivity("MainActivity", func(app *App) {
        time.Sleep(time.Second * 3)
    })



    MyApp.LastLogs = append(MyApp.LastLogs,
        Slog("!", "Working ...", "green"),
    )
    MyApp.Run()
    // Def("Hello!", "Hi! This is JR. What's your name?", "I'm RJ.")
}

