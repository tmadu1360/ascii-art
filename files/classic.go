package ascii_art

import (
	"fmt"
	"strings"
)

func Classic(args []string)  {

	if len(args) < 2 {

		noArgument()
		classicUsage()
		errorExit()

	} else if len(args) > 2 {

		tooManyArguments()
		classicUsage()
		errorExit()

	}

	var text = args[1]
	var format = "standard"
	var result = ""

	if newlineCheck(text) { //  == true

		var splitText = strings.Split(text,"\\n")

		for _, v := range splitText {

			for i := 0; i < 8; i++ {

				for _, v2 := range v {

					result += scanner(format, 1+int(v2-32)*9+i)

				}

				fmt.Println(result)
				result = ""

			}

		}

	} else { //if newlineCheck(text) == false

			for i := 0 ; i < 8 ; i++ {

				for _, v := range text {

					result += scanner(format, 1 + int(v - 32) * 9 + i)

				}

				fmt.Println(result)
				result = ""

			}

	}

}