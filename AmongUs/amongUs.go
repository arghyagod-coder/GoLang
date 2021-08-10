package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var bas string
	fmt.Println("Can we start?")
	fmt.Scanln(&bas)
	if strings.ToLower(bas) == "no" {
		return
	} else if strings.ToLower(bas) == "yes" {
		var latika string
		fmt.Println("Local and Freeplay options are under-development. You can only choose Online\nChoose: Local, Online or FreePlay ")
		fmt.Scanln(&latika)
		rand.Seed(time.Now().UnixNano())
		var (
			maun  int = rand.Intn(3-1+1) + 1
			baun  int = rand.Intn(7-1+1) + 1
			bawa  int = rand.Intn(3-1+1) + 1
			banan int = rand.Intn(45-15+1) + 15
		)
		if strings.ToLower(latika) == "online" {
			fmt.Println("You are in waiting room.")
			fmt.Println("SHOWING MEETING SETTINGS....")
			fmt.Printf("Impostors: %v\n", maun)
			fmt.Println("Emergency Meetings:", baun)
			fmt.Println("Player Speed: ", bawa, "x")
			fmt.Println("Impostor Speed: ", bawa, "x")
			fmt.Println("Crew Light:", bawa, "x")
			fmt.Println("Impostor Light: ", bawa, "x")
			fmt.Println("Kill Cooldown:", banan, "s")
			fmt.Println("Common Tasks: ", baun)
			fmt.Println("Long Tasks: ", baun)
			fmt.Println("Short Tasks: ", baun)
			var pyu string
			fmt.Println("See the meeting settings and type Continue: ")
			fmt.Scanln(&pyu)
			if strings.ToLower(pyu) == "continue" {
				var rant [7]string
				rant[0] = "White"
				rant[1] = "Red"
				rant[2] = "Green"
				rant[3] = "Blue"
				rant[4] = "Black"
				rant[5] = "Brown"
				rant[6] = "Lime"
				var kpop = rant[rand.Intn(len(rant))]
				fmt.Println("You are", kpop)
				fmt.Println("Starting in 5...")
				fmt.Println("Starting in 4...")
				fmt.Println("Starting in 3...")
				fmt.Println("Starting in 2...")
				fmt.Println("Starting in 1...")
				fmt.Println("Starting in 0...")
				var duck [2]string
				duck[0] = "You are a Crewmate"
				duck[1] = "You are an Impostor"
				dichk := duck[rand.Intn(len(duck))]
				var (
					poop   string = rant[rand.Intn(len(rant))]
					kpbap  string = rant[rand.Intn(len(rant))]
					tarap  string = rant[rand.Intn(len(rant))]
					duffer string = rant[rand.Intn(len(rant))]
					// decki  string = rant[rand.Intn(len(rant))]
				)
				fmt.Println(dichk)
				var pau string
				fmt.Println("type Continue: ")
				fmt.Scanln(&pau)
				if strings.ToLower(pyu) == "continue" {
					var pint [4]string
					pint[0] = "You are in Electrical\nFixing Lights"
					pint[1] = "You are in Admin\nDownloading Files"
					pint[2] = "You are in Reactor\nReactor fixing."
					pint[3] = "You are in Hall\nDownloading Files"
					ghuji := pint[rand.Intn(len(pint))]
					fmt.Println(ghuji)
					var paint [4]string
					paint[0] = "Electrical"
					paint[1] = "Admin"
					paint[2] = "Hall"
					paint[3] = "Reactor"
					ghap := paint[rand.Intn(len(paint))]
					if strings.Contains(dichk, "Impostor") == true {
						fmt.Println(poop, "is passing by.")
						var ighf string
						fmt.Println("Will you Kill Him.... Type Yes or No; ")
						fmt.Scanln(&ighf)
						if strings.Contains(ighf, "yes") == true{
							fmt.Println(poop, "was killed")
							fmt.Println(poop, "'s dead body was reported by", tarap)
							fmt.Println(tarap, ": I found his body in ", ghap, "I guess it is ", kpbap)
							fmt.Println(duffer, ": I think it is", tarap, "or", kpop, "maybe")
							var damg string
							fmt.Println("A player is suspecting you. Type yes or no in your Defence: ")
							fmt.Scanln(&damg)
							fmt.Println(kpop, ":", damg, "I was in Elec")
							fmt.Println("Skip Vote. Yes or No? " )
							var ghazi string
							fmt.Scanln(&ghazi)
							var sgd[3] string
							sgd[0] = "The Other Impostor killed most of the mates\nYou Win\nCongratulations!"
							sgd[1] = "The Other Impostor left the game. \nYou were voted out\nYou lost."
							sgd[2] = "Reactor Meltdown Successful. \nYou win\nCongratulations!"
							ghfd := sgd[rand.Intn(len(sgd))]
							fmt.Println(ghfd)
							fmt.Println("Run the program again to play")
					} else if strings.Contains(dichk, "Crewmate"){
					
					var(
						tapap   string = rant[rand.Intn(len(rant))]
						tarip  string = rant[rand.Intn(len(rant))]
						rap  string = rant[rand.Intn(len(rant))]
						karap string = rant[rand.Intn(len(rant))]
					)
					fmt.Println(tapap, "called an emergency meeting")
					fmt.Println(karap, "was killed")
					fmt.Println(karap, "'s dead body was reported by", tapap)
					fmt.Println(tapap, ": I found his body in ", ghap, "I guess it is ", rap)
					fmt.Println(tarip, ": I think it is", rap, "or", kpop, "maybe")
					fmt.Println("A player is suspecting you. Type yes or no in your Defence: ")
					var damg string
					fmt.Scanln(&damg)
					fmt.Println(kpop, ":", damg, "I was in Elec")
					fmt.Println("Skip Vote. Yes or No? " )
					var sgd[3] string
					sgd[0] = "The Impostor killed you\nReactor Melted\nYou Lost"
					sgd[1] = "One Impostor left the game. \nDoing tasks in electrical\nEmergency Meeting\nYou were voted out.\nYou lost."
					sgd[2] = "Impostors voted out \nYou win\nCongratulations!"
					ghfd := sgd[rand.Intn(len(sgd))]
					fmt.Println(ghfd)
					fmt.Println("Run the program again to play")
	
				}
				}
				}
			}
		}
	}
}
