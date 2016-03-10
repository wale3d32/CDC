package greetings 

//GreetingsString is used as a greeting string
var GreetingsString = "Hello World from Waleed :3"

//PrintGreetings prints the greetings on your screen (The title is pretty much self explanatory) 
func PrintGreetings (name string) string {	
	return GreetingsString + "-" + name; 
}
