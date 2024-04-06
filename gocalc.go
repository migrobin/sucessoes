package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)
       

func getUserInput() (string, error) {
    
    scanner := bufio.NewScanner(os.Stdin)

      scanner.Scan()
  
    if err := scanner.Err(); err != nil {
        return "", err
    }

    return scanner.Text(), nil
}


func isValidInput(str string)([3]float32,error){
 var arr [3]float32


 _, err := fmt.Sscanf(str, "%v,%v,%v", &arr[0], &arr[1], &arr[2])
 if err != nil {
        return arr,err
    }
    
    return arr,nil      
}


func print (num float32, ter float32) string {
    if num == 0 {
        return ""
    }

    if ter == 1 {
        if num == 1 {
            return "n²"
        }
        return fmt.Sprintf("%gn²", num )
    }

    if ter == 2 {
        if num == 1 {
            return " +n"
        }
        if num > 1 {
            return fmt.Sprintf(" +%gn", num)
        }
        return fmt.Sprintf(" %gn", num)
    }

    if num > 0 {
        return fmt.Sprintf(" +%g", num)
    }
    return fmt.Sprintf(" %g", num)
}


func main() {

    fmt.Println("-------------------------------------------------------------------------")
 fmt.Println("\033[33;1mBem vindo ao goCalc\033[0m\n")
 fmt.Println("Calcula um termo geral de uma sucessão a partir de 3 termos consecutivos\n")
 var s[3]float32
 
 for {
 fmt.Println("introduz 3 termos separados por uma virgula\n")

 var userInput ,err = getUserInput()

if err != nil {
    fmt.Println("there was an error, abortar a coisa ", err)
    return
}


s, err = isValidInput(userInput)

if err != nil {
    fmt.Println("\nDados invalidos. \033[97;1mTás a gozar Bro ? \033[0m  Ctrl + C para sair")
    fmt.Println(err)
  
} else {
    break
}


}
  
fmt.Println("\n-------------------------------------------------------------------------")
    fmt.Println("Sucessão  -> \033[33;1m", s, "\033[0m   O termo geral será da forma \033[34;1m Un = an² + bn + c\033[0m")
    fmt.Println("\nVamos calcular os 3 primeiros termos :")

    var a,b,c float32

    a = (s[2] - 2*s[1] + s[0]) / 2
    b = (-5*s[0] + 8*s[1] - 3*s[2]) / 2
    c = (6*s[0] - 6*s[1] + 2*s[2]) / 2

    fmt.Println("\nU1: -> a(1²) + b(1) + c =\033[33m", s[0], "\033[0m    (1) -> a + b + c =\033[33m", s[0],"\033[0m")
    fmt.Println("U2: -> a(2²) + b(2) + c =\033[33m", s[1], "\033[0m    (2) -> 4a + 2b + c =\033[33m", s[1],"\033[0m")
    fmt.Println("U3: -> a(3²) + b(3) + c =\033[33m", s[2], "\033[0m    (3) -> 9a + 3b + c =\033[33m", s[2],"\033[0m")

    fmt.Println("\nSubtrair (2)-(1) =>  3a + b =", (s[1] - s[0]))
    fmt.Println("Subtrair (3)-(2) =>  5a + b =", (s[2] - s[1]))
    fmt.Println("                     -----------")
    fmt.Println("Subtraindo           2a = ", (s[2] - s[1] - (s[1] - s[0])), "     => a =", a, "      => b =", b, "  => c =", c)

    temp:= print(a, 1) + print(b, 2) + print(c, 3)
    if temp == "" {
        fmt.Println("\n\nO termo geral será  \033[33;1m  Un =", 0 ,"\033[0m\n")
    }  else {
    fmt.Println("\n\nO termo geral será  \033[33;1m  Un =", temp,"\033[0m\n")
}


    

    fmt.Println(("have a nice day!"))

    fmt.Println(("press any key to exit"))
    keyboard.GetSingleKey()



}

	