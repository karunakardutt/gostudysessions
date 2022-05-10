package main

/* This is code from publicly available blog: https://blogvali.com/password-generator-fyne-gui-golang-tutorial-29-30-31/
And associated youtube videos : https://www.youtube.com/watch?v=DWklQa8S24k&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=29
and https://www.youtube.com/watch?v=tzBeu3kGaJk&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=30&t=8s and 
https://www.youtube.com/watch?v=tzBeu3kGaJk&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=30&t=8s
*/



import (
    "fmt"
    "image/color"
    "math/rand"
    "strconv"
    "time"
    // import fyne
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)
func main() {
    // New App
    a := app.New()
    // New title and window
    w := a.NewWindow("Password Generator")
    // resize window
    w.Resize(fyne.NewSize(400, 400))
    title := canvas.NewText("Password Generator", color.Black)
    // input Box
    input := widget.NewEntry()
    input.SetPlaceHolder("Enter password length")
    //label to show password
    text := canvas.NewText("", color.Black)
    text.TextSize = 20
    // button to generate password
    btn1 := widget.NewButton("Generate", func() {
        // input
        passlength, _ := strconv.Atoi(input.Text) // convert string to integer
        text.Text = PasswordGenerator(passlength)
        text.Refresh()
    })
    // show content
    w.SetContent(
        container.NewVBox(
            // put all widgets here
            title,
            input,
            text,
            btn1,
        ),
    )
    w.ShowAndRun()
    // Fyne appp
}
// Converting the code to a function
func PasswordGenerator(passwordLength int) string {
    return "Dummy password"
}
