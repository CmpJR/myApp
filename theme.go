package main

import "github.com/gdamore/tcell/v2"



type Theme struct {
    AppTheme    AppTheme
    UserStyle   map[string] tcell.Style
}



type AppTheme struct {
    AppStyle    tcell.Style
    BGColor     [3] int32
    FGColor     [3] int32
    Cursor      tcell.CursorStyle
}

func (t *AppTheme) CreateAppStyle() {
    t.AppStyle = tcell.StyleDefault.
        Background(tcell.NewRGBColor(t.BGColor[0], t.BGColor[1], t.BGColor[2])).
        Foreground(tcell.NewRGBColor(t.FGColor[0], t.FGColor[1], t.FGColor[2]))
}

