package main

import (
	"os"
	"fmt"

	"github.com/gdamore/tcell/v2"
)



// Create new app
func NewApp(t *Theme) (App, tcell.Screen, error) {
    s, err := tcell.NewTerminfoScreenFromTtyTerminfo(nil, nil)
    if err != nil {
        if err == tcell.ErrNoScreen {
            Log("Error", "Unsupported screen", "red")
            return App {}, nil, ErrAppInitialization
        } else {
            Log("Error", err.Error(), "red")
            return App {}, nil, err
        }
    }

    ret, err := GetTermSize()
    if err != nil {
        Log("Error", err.Error(), "red")
        return App {}, nil, ErrAppInitialization
    }
    return App {
        Height:     ret.Height,
        Width:      ret.Width,
        Theme:      t,
        activities: make(map[string]Activity),
        Screen:     s,
    }, s, nil
}



type App struct {
    Height              uint16
    Width               uint16
    Theme               *Theme
    activities          map[string] Activity
    ActiveActivity      string
    Screen              tcell.Screen
    LastLogs            [] string
    isScreenInitialized bool
    SupportMouse        bool
}

// Run the app
func (app *App) CreateActivity(name string, OnRun func(*App)) {
    app.activities[name] = Activity{
        Run: OnRun,
    }
}

// Run specefic activity
func (app *App) RunActivity(name string) error {
    activity, exists := app.activities[name]
    if !exists {
        return ErrActivityNotExists
    }
    activity.Run(app)
    return nil
}

// Run the app
func (app *App) Run() {
    // Initialize the screen
    if err := app.Screen.Init(); err != nil {
        Log("Error", "Unable to initialize screen", "red")
        app.Exit(1)
    }
    app.isScreenInitialized = true

    // Set default screen/app style
    if app.Theme.AppTheme.AppStyle == tcell.StyleDefault {
        app.Theme.AppTheme.CreateAppStyle()
    }
    
    // Configure tcell screen
    if app.Screen.HasMouse() {
        app.Screen.EnableMouse(tcell.MouseButtonEvents)
        app.SupportMouse = true
    } else {
        app.Screen.DisableMouse()
    }
    app.Screen.DisableFocus()
    app.Screen.DisablePaste()
    app.Screen.SetStyle(app.Theme.AppTheme.AppStyle)
    app.Screen.SetCursorStyle(tcell.CursorStyleSteadyBar)

    if err := app.RunActivity("MainActivity"); err != nil {
        app.LastLogs = append(app.LastLogs, 
            Slog("Error", err.Error(), "red"),
        )
        app.Exit(1)
    }

    app.Exit(0)
}

// Exit app
func (app *App) Exit(code int) {
    if app.isScreenInitialized {
        app.Screen.SetCursorStyle(tcell.CursorStyleDefault)
        app.Screen.Fini()
    }
    for _, log := range app.LastLogs {
        fmt.Println(log)
    }
    os.Exit(code)
}

