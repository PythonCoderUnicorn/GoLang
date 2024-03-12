/*
  this program tests out terminal font color formatting
*/

package main

import (
	"fmt"
	"time"
	"github.com/ttacon/chalk"   // no http:// allowed
	
	//"github.com/alecthomas/colour"
	"github.com/gookit/color"

	//"github.com/theckman/yacspin"

	"github.com/charmbracelet/lipgloss"

	"github.com/mingrammer/cfmt"         
	cf "github.com/i582/cfmt/cmd/cfmt"  // the better cfmt okg

	"github.com/schollz/progressbar/v3"

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

	
	

	//------ lip gloss
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	pl("")
	fmt.Println(style.Render("Hello, kitty"))
	pl("")

	var style2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#3333cc")).
		Background(lipgloss.Color("#00cc00")).
		PaddingTop(2).
		PaddingLeft(4).
		PaddingBottom(2).
		Width(30)

	pl("")
	pl(style2.Render("put on some lipgloss"))
	pl("")
	

    var style3 = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(lipgloss.Color("#3333cc")).
		Background(lipgloss.Color("#00cc00")).
		Faint(true).
		Blink(true).
		//Strikethrough(true).
		//Underline(true).
		Reverse(true)

	pl("")
	pl(style3.Render("lipgloss message blink ! "))
	pl("")
	
	


	// Add a purple, rectangular border
	var style4 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))

	// Set a rounded, yellow-on-purple border to the top and left
	var style5 = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color("#9933ff")). // purple bg
		Bold(true).
		Blink(true).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00ffcc")).
		BorderBackground(lipgloss.Color("#9933ff")). // purple bg
		//BorderTop(true).
		//BorderLeft(true)
		Padding(2).
		Margin(3).
		Width(60).
		Border(lipgloss.DoubleBorder(), true, true, true, true)



	pl("")
	pl(style4.Render("style 4 lipgloss"))
	//pl("")
	pl(style5.Render("style 5 lipgloss: Insert Coins to Play"))
	pl("this text is not blinking")
	


	pl("")
	cfmt.Success("cfmt import was a success!\n")
	cfmt.Infoln("cfmt information [:] just some information")
	cfmt.Warningf("cfmt warning text\n")

	// the other cfmt pkg
	cf.Println("\nthis is a {{red color}}::red")
	cf.Println("this is another {{cfmt pkg}}::cyan|bold")
	cf.Println("This is a {{hex pink color}}::#ff66ff")
	cf.Println("cfmt {{background text}}::bgRed  ")
	cf.Println("cfmt {{background hex color}}::bg#6600ff")

	// green yellow red magenta blue = apple logo
	cf.Print(`
        {{                 -/+:.          }}::#34BE2D
        {{                :++++.          }}::#34BE2D
        {{               /+++/.           }}::#34BE2D
        {{       .:-::- .+/:-''.::-       }}::#34BE2D
        {{    .:/++++++/::::/++++++//:    }}::#34BE2D
        {{  .:///////////////////////:    }}::#FFB400
        {{  ////////////////////////      }}::#FFB400
        {{ -+++++++++++++++++++++++       }}::#FF7A00
        {{ /++++++++++++++++++++++/       }}::#FF7A00
        {{ /sssssssssssssssssssssss.      }}::#F41E34
        {{ :ssssssssssssssssssssssss-     }}::#F41E34
        {{  osssssssssssssssssssssssso/   }}::#A2359C
        {{   syyyyyyyyyyyyyyyyyyyyyyyy+   }}::#A2359C
        {{    ossssssssssssssssssssss/    }}::#00A0E2
        {{      :ooooooooooooooooooo+.    }}::#00A0E2
        {{       :+oo+/:-..-:/+o+/-       }}::#00A0E2
`)

cf.Println("this cfmt {{blinks i think}}::#00cc99|blink")
cf.Println("reverse {{reverse}}::lightYellow|reverse")
cf.Println("this is {{concealed}}::lightCyan|bold|concealed")

pl("")

//-- progress bar
//bar := progressbar.Default(100)
bar := progressbar.NewOptions(1000,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionShowBytes(true),
	progressbar.OptionSetWidth(15),
	progressbar.OptionSetDescription("[magenta]updating LCARS ..."),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer: 		"[green]=[reset]",
		SaucerHead: 	"[green]>[reset]",
		SaucerPadding: 	" ",
		BarStart: 		"[magenta][",
		BarEnd: 		"[magenta]]",
	}))

for i := 0; i < 800; i++ {
    bar.Add(1)
    time.Sleep(10 * time.Millisecond)
}

bar2 := progressbar.NewOptions(1000,
    //progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //you should install "github.com/k0kubun/go-ansi"
    progressbar.OptionEnableColorCodes(true),
    progressbar.OptionShowBytes(true),
    progressbar.OptionSetWidth(15),
    progressbar.OptionSetDescription("[cyan]writing to file..."),
    progressbar.OptionSetTheme(progressbar.Theme{
        Saucer:        "[green]=[reset]",
        SaucerHead:    "[green]>[reset]",
        SaucerPadding: " ",
        BarStart:      "[cyan][",
        BarEnd:        "[cyan]]",
    }))
for i := 0; i < 1000; i++ {
    bar2.Add(1)
    time.Sleep(5 * time.Millisecond)
}
	
	
	
	
	
}
