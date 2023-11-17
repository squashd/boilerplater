package ui

import "fmt"

const LogoArt = `                                           
 _       _ _             _     _           
| |_ ___|_| |___ ___ ___| |___| |_ ___ ___ 
| . | . | | | -_|  _| . | | .'|  _| -_|  _|
|___|___|_|_|___|_| |  _|_|__,|_| |___|_|  
                    |_|                   
`

func RenderLogo() {
	fmt.Println(Logo.Render(LogoArt))
}

func PrettyPrintChoice(str string) {
	fmt.Printf("%s\n", Unselected.Render(str))
}
