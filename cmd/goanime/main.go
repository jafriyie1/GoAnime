package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jafriyie1/goanime/animescrapper"
	"github.com/jafriyie1/goanime/animetries"
)

func main() {

	f, err := os.Open("../../Data/merge/full.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(f))

	_, builtTrie, _ := animetries.BuildAnimeTrie(r)
	var option string

	scanner := bufio.NewScanner(os.Stdin)
	searchedShow := animescrapper.GetShow(builtTrie)
	searchedShow = strings.TrimSpace(searchedShow)

	fmt.Println()
	//season := animescrapper.GetSeason()
	season := ""
	fmt.Println("Here is a list of episodes for the given show and season (please wait):")

	animescrapper.GetEpisodeList(searchedShow, season)

	fmt.Println()
	time.Sleep(5)
	fmt.Println("Scroll up to view episodes (and please ignore other messages).\nWould you like to watch one episode or mutliple (1 for episode, 2 for multiple)")

	if scanner.Scan() {
		option = scanner.Text()
	}
	option = strings.TrimSpace(option)
	var lowerLimitEpisode string

	upperLimitEpisode := " "

	if option == "1" {
		lowerLimitEpisode, lowerLimitEpisode = animescrapper.GetRangeOfEpisodes(false)
		upperLimitEpisode = lowerLimitEpisode
	} else {
		//fmt.Println("WARNING: You can only get a maximum of 2 episodes.\nOutside of that you will get wonky behavior.")

		lowerLimitEpisode, upperLimitEpisode = animescrapper.GetRangeOfEpisodes(true)
	}

	lowerEpisode, _ := strconv.Atoi(lowerLimitEpisode)
	upperEpisode, _ := strconv.Atoi(upperLimitEpisode)

	fmt.Println("Please wait....")
	for i := lowerEpisode; i < upperEpisode+1; i++ {

		c, ctxt := animescrapper.DoGoAnime()
		fmt.Println("Getting Episode", i)
		loopedEpisode := strconv.Itoa(i)
		animescrapper.ConcurrentEpisodes(loopedEpisode, upperLimitEpisode, searchedShow, season, c, ctxt)
	}
}
