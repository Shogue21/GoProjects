package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	playerChoice   int
	playerHealth   int
	computerHealth int
	playAgain      string
	round          int
	choiceList     string
)

func printTypes() {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	brown := color.New(color.FgBlue, color.FgYellow).SprintFunc()
	fmt.Printf("[1] %s\n[2] %s\n[3] %s\n[4] %s\n[5] %s", red("Fire"), green("Grass"), brown("Rock"), blue("Ice"), yellow("Ground"))
}

func validType(choice int) bool {
	return choice == 1 || choice == 2 || choice == 3 || choice == 4 || choice == 5
}

func userSettings() int {
	for {
		fmt.Println("What type would you like to choose?")
		printTypes()
		fmt.Println("")
		fmt.Scanln(&playerChoice)
		if validType(playerChoice) {
			break
		} else {
			fmt.Println("Please choose a valid type!")
		}
	}
	return playerChoice
}
func comSettings() int {
	rand.Seed(time.Now().UnixNano())
	comChoice := rand.Intn(4)
	return comChoice
}

func playerDamage() int {
	pDamage := 20
	rand.Seed(time.Now().UnixNano())
	missAttack := rand.Intn(8)
	if missAttack == 1 {
		pDamage = 0
		fmt.Println("Your attack missed!")
	}
	critAttack := rand.Intn(8)
	if pDamage != 0 && critAttack == 1 {
		pDamage = pDamage * 2
		fmt.Println("You got a critical hit!")
	}
	return pDamage
}

func computerDamage() int {
	cDamage := 20
	rand.Seed(time.Now().UnixNano())
	missAttack := rand.Intn(8)
	if missAttack == 1 {
		cDamage = 0
		fmt.Println("The computer's attack missed!")
	}
	critAttack := rand.Intn(8)
	if cDamage != 0 && critAttack == 1 {
		cDamage = cDamage * 2
		fmt.Println("The computer got a critical hit!")
	}
	return cDamage
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else if num2 > num1 {
		return num2
	} else {
		return num1
	}
}

func game() {

	playerHealth := 100
	comHealth := 100
	fmt.Println("Here are your choices")
	printTypes()
	fmt.Println("")
	for {
		fmt.Println("")
		userChoice := userSettings()
		comChoice := comSettings()
		fmt.Println("")
		if userChoice == 1 {
			if comChoice == 0 {
				fmt.Println("You both chose fire! It's a tie!")
				playerHealth -= 0
				comHealth -= 0
			} else if comChoice == 1 {
				fmt.Println("The computer chose grass! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 2 {
				fmt.Println("The computer chose ice! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 3 {
				fmt.Println("The computer chose rock! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			} else if comChoice == 4 {
				fmt.Println("The computer chose ground! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			}
		} else if userChoice == 2 {
			if comChoice == 0 {
				fmt.Println("You both chose grass! It's a tie!")
				playerHealth -= 0
				comHealth -= 0
			} else if comChoice == 1 {
				fmt.Println("The computer chose rock! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 2 {
				fmt.Println("The computer chose ground! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 3 {
				fmt.Println("The computer chose fire! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			} else if comChoice == 4 {
				fmt.Println("The computer chose ice! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			}
		} else if userChoice == 3 {
			if comChoice == 0 {
				fmt.Println("You both chose rock! It's a tie!")
				playerHealth -= 0
				comHealth -= 0
			} else if comChoice == 1 {
				fmt.Println("The computer chose fire! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 2 {
				fmt.Println("The computer chose ice! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 3 {
				fmt.Println("The computer chose grass! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			} else if comChoice == 4 {
				fmt.Println("The computer chose ground! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			}
		} else if userChoice == 4 {
			if comChoice == 0 {
				fmt.Println("You both chose ice! It's a tie!")
				playerHealth -= 0
				comHealth -= 0
			} else if comChoice == 1 {
				fmt.Println("The computer chose grass! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 2 {
				fmt.Println("The computer chose ground! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 3 {
				fmt.Println("The computer chose fire! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			} else if comChoice == 4 {
				fmt.Println("The computer chose rock! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			}
		} else if userChoice == 5 {
			if comChoice == 0 {
				fmt.Println("You both chose ground! It's a tie!")
				playerHealth -= 0
				comHealth -= 0
			} else if comChoice == 1 {
				fmt.Println("The computer chose fire! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 2 {
				fmt.Println("The computer chose rock! You win!")
				comHealth -= playerDamage()
				comHealth = max(0, comHealth)
			} else if comChoice == 3 {
				fmt.Println("The computer chose ice! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			} else if comChoice == 4 {
				fmt.Println("The computer chose grass! You lose!")
				playerHealth -= computerDamage()
				playerHealth = max(0, playerHealth)
			}
		}
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		fmt.Println("")
		fmt.Printf("Player Health: %v", green(playerHealth))
		fmt.Println("")
		fmt.Printf("Computer Health: %v", red(comHealth))
		fmt.Println("")

		if playerHealth == 0 {
			fmt.Println("You were defeated by the computer!")
			fmt.Println("")
			break
		}
		if comHealth == 0 {
			fmt.Println("You defeated the computer!")
			fmt.Println("")
			break
		}
	}
}
func main() {
	game()
	for {
		fmt.Println("Would you like to play again? [Y/N]:")
		fmt.Scanln(&playAgain)
		playAgain := strings.ToLower(playAgain)
		if playAgain == "y" {
			round++
			fmt.Printf("Round %v", round)
			fmt.Println("")
			game()
		} else if playAgain == "n" {
			fmt.Println("I hope you had fun!")
			break
		} else {
			fmt.Println("Please choose a valid option!")
		}
	}
}
