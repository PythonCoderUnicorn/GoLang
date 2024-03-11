package main

import (
	"fmt"
	//"time"
	"github.com/ttacon/chalk"   // no http:// allowed
	
	//"github.com/alecthomas/colour"
	"github.com/gookit/color"

	//"github.com/theckman/yacspin"

)

//  default colors: black, red, green, yellow, blue, magenta, cyan and white

//  text styles: bold, dim, italic, underline, inverse, hidden and strikethrough


var pl = fmt.Println

func main() {

	fmt.Println("hello")
	
	lime := chalk.Green.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).
		Style
		
	fmt.Println( lime("this is lime chalk" ) )
	
	pl( chalk.Bold.TextStyle("so bold!") )
	
	pl( chalk.Italic.TextStyle("italic seasoning") )
	
	c := chalk.Magenta.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).Style
	pl( c("magenta text" ) )
	
	
	bl := chalk.Black.NewStyle().
		WithBackground(chalk.Yellow).
		WithTextStyle(chalk.Underline).Style
	
	pl( bl("[*] black text on yellow ") )
	
	
	red := chalk.Red.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Italic).Style
		
	pl( red(" this is a test of the emergency code red ") )
	
	
	//colour.Printf(" ^3 yellow text from colour pkg")
	

	//--------------
	color.Cyanln("\ncyan line from color")
	color.Warn.Println("Warn.Println -- warning text")
	color.Info.Println("Info.Println -- this is info")

	color.Styles["debug"].Println("oh no! a bug was found in code")
	color.Styles["danger"].Println("there is danger in the code")
	color.Styles["primary"].Println("primary message!")
	color.Styles["success"].Println("this was a success!")

	
	



	


	
	
	
	
	
	
	
	
	
}
