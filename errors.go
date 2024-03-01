package main

import "errors"

var (
    // When app and tcell screen creation not possible
    ErrAppInitialization = errors.New("App: Unable to create screen for app")

    // When try to work with non-existent activity
    ErrActivityNotExists = errors.New("App: Activity not found")
)

