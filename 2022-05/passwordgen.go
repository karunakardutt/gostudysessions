package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randPwdGen(a int) string {
	var retval string
	pwdchar := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_"
	for ; a > 0; a-- {
		rand.Seed(time.Now().UnixNano())
		retval = retval + string(pwdchar[rand.Intn(len(pwdchar))])
	}
	return retval
}

func main() {
	pwd_len := 25
	if len(os.Args) == 2 {
		pwd_len, _ = strconv.Atoi(os.Args[1])
	}
	PWD := randPwdGen(pwd_len)
	fmt.Println(PWD)
}
