/*
 this Command Line program converts temperatures 
  -c num   converts F to C
  no flag  converts C to F
*/

package main

import (
    "flag"
    "fmt"
	"github.com/charmbracelet/lipgloss" // for color and styled box
	cf "github.com/i582/cfmt/cmd/cfmt"  // for color text formatting
)

var pl = fmt.Println 

// Define a flag for specifying the conversion direction
var toCelsius = flag.Bool("c", false, "Convert Fahrenheit to Celsius (default: false)")




func main() {
    // Parse command line flags
    flag.Parse()

    // Get the temperature value from arguments (optional)
    var temperature float64
    if len(flag.Args()) == 1 {
        fmt.Sscan(flag.Arg(0), &temperature) // Scan the argument for a float
    } else {
        fmt.Println("Usage: tempconv [-c] [temperature]")
        return
    }

    // Convert temperature based on the flag
    var convertedTemp float64
    if *toCelsius {
        convertedTemp = fahrenheitToCelsius(temperature)
    } else {
        convertedTemp = celsiusToFahrenheit(temperature)
    }

    // Print the result


	pl("")

	var style4 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Foreground(lipgloss.Color("#3333cc")).
		Bold(true).
		Padding(1).
		Margin(1/2).
		Width(25)
	
	
	pl(style4.Render("Temperature Conversion"))

	gunc := getUnit(!*toCelsius)
	cvt := convertedTemp
	guc := getUnit(*toCelsius)

    cf.Printf("{{%.2f}}::#66ffff|bold {{%s}}::#66ffff|bold  = {{%.2f}}::#99ff66|bold degrees {{%s}}::#99ff66|bold \n",
        temperature, guc, cvt, gunc)
}

// Function to convert Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5 / 9
}

// Function to convert Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
    return (celsius * 9 / 5) + 32
}

// Function to get the unit string based on the conversion direction
func getUnit(toCelsius bool) string {
    if toCelsius {
        return "\u2109"//"Fahrenheit"
    }
    return "\u2103"//"Celsius"
}
