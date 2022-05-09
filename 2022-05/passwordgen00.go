package main

/* This is code is adapted from publicly available blog: https://blogvali.com/password-generator-fyne-gui-golang-tutorial-29-30-31/
And associated youtube videos : https://www.youtube.com/watch?v=DWklQa8S24k&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=29
and https://www.youtube.com/watch?v=tzBeu3kGaJk&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=30&t=8s and 
https://www.youtube.com/watch?v=tzBeu3kGaJk&list=PL5vZ49dm2gshlo1EIxFNcQFBUfLFoHPfp&index=30&t=8s

The code is used for learning golang in a study session.
*/


import (
    "fmt"
    "image/color"
    "math/rand"
    "strconv"
    "time"
    // import fyne
    //    "fyne.io/fyne/v2"
    //"fyne.io/fyne/v2/app"
    //"fyne.io/fyne/v2/canvas"
    //"fyne.io/fyne/v2/container"
    //"fyne.io/fyne/v2/widget"
)
func main() {
    // Password Generator
    // Lower case
    lowCase := "abcdefghijklmnopqrstuvxyz"
    // Upper Case
    upCase := "ABCDEFGHIJKLMNOPQRSTUVXYZ"
    // Numbers
    Numbers := "0123456789"
    // Special characters
    SpecialChar := "£$&()*+[]@#^-_!?"
    // Password Length
    passwordLength := 8
    // variable for storing password
    password := ""
    
    // loop
    for n := 0; n < passwordLength; n++ {
        // Now random characters
        rand.Seed(time.Now().UnixNano())
        randNum := rand.Intn(4)
        fmt.Println(randNum)
        // Switch statment
        switch randNum {
        case 0:
            rand.Seed(time.Now().UnixNano())
            randNum := rand.Intn(len(lowCase))
            // len to find lenth of slice/array
            // Now we will store the generated passowrd character
            password = password + string(lowCase[randNum])
            // it is byte... we need string
        // first case completed
        case 1:
            rand.Seed(time.Now().UnixNano())
            randNum := rand.Intn(len(upCase))
            // len to find lenth of slice/array
            // Now we will store the generated passowrd character
            password = password + string(upCase[randNum])
            // it is byte... we need string
            // second done 🙂
        case 2:
            rand.Seed(time.Now().UnixNano())
            randNum := rand.Intn(len(Numbers))
            // len to find lenth of slice/array
            // Now we will store the generated passowrd character
            password = password + string(Numbers[randNum])
            // it is byte... we need string
            // 3rd done 🙂
        case 3:
            rand.Seed(time.Now().UnixNano())
            randNum := rand.Intn(len(SpecialChar))
            // len to find lenth of slice/array
            // NOw we will store the generated passowrd character
            password = password + string(SpecialChar[randNum])
            // it is byte... we need string
        } // end of switch
    } // end of for loop
    fmt.Println(password)
}
