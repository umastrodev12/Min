// Hey! Min Run on terminal, he is NOT an App, i'm creating .exe version, enjoy the program~

package main

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func server() {
	r := gin.Default() // for routes

	// routes
	r.GET("/version", func(c* gin.Context) {
		c.JSON(200, gin.H{
			"version": "0.0.1 BETA",
			"author": "Astro",
		})
	})
}

func run() {
	err := errors.New("An Error Has Ocurred.")
	if err != nil {
		return
	}

	prefix := string("min://")
	port := string(":8080")

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Min is Starting...")
	fmt.Println("Current Version: 0.0.1 BETA")
	fmt.Println("Prefix:", prefix)
	fmt.Println("Port:", port)
	fmt.Println("Alert: Min can have some issues because, Min has started TOO EARLY, bye~")
	fmt.Println("------------------------------------------------------------------------")

	cmd := exec.Command("Min.exe")

	if err := cmd.Start(); err != nil {
		fmt.Println("Cannot Open Min.exe!")
		return
	}

	cmd.Wait() // close min
}
