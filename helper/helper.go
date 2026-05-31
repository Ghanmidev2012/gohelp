package helper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"math/rand"
)

func SayHello(name string){
	fmt.Println("Hello, ", name)
}

func Clear()  {
	var cmd *exec.Cmd

	if  runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else{
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Wait(seconds int)  {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func Random(min, max int) int{
	if  min >= max {
			return min
	}
	return rand.Intn(max-min+1) + min
}

func Input(prompt string) string {
	var value string
	fmt.Print(prompt)

	fmt.Scan(&value)

	return value
}
