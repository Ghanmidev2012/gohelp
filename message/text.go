package message

import ("fmt")

func WelcomeMessage(name string) string {
	return fmt.Sprintf("Welcome, %s!",name )
}