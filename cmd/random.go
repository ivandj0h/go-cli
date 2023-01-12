/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Random command",
	Long:  `This is a new Random Command`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandJoke() Joke {
	resp, err := http.Get("https://icanhazdadjoke.com/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var joke Joke
	json.NewDecoder(resp.Body).Decode(&joke)

	return joke
}

func getJokeData(baseAPI string) []byte {
	resp, err := http.Get(baseAPI)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var jokeData []byte
	json.NewDecoder(resp.Body).Decode(&jokeData)

	return jokeData
}
