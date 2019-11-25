package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MarinX/go-klingon/stapi"
	"github.com/MarinX/go-klingon/translate"
)

var apiKey = flag.String("apikey", "", "stapi api key")

func usage() {
	fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS] argument\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	name := strings.Join(os.Args[1:len(os.Args)], " ")

	if err := handleTranslate(name); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if err := handleStapi(name, *apiKey); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func handleTranslate(name string) error {
	hValue, err := translate.New(name).Klingon()
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, hValue)
	return nil
}

func handleStapi(name, apiKey string) error {
	cl := stapi.New(apiKey, nil)

	// search the character
	search, err := cl.Character.Search(struct {
		Name string `url:"name"`
	}{name})
	if err != nil {
		return err
	}

	// no character found
	if len(search.Characters) <= 0 {
		fmt.Fprintln(os.Stdout, "Unknown")
		return nil
	}

	// character found, take the first
	ch, err := cl.Character.Get(struct {
		UID string `url:"uid"`
	}{search.Characters[0].UID})
	if err != nil {
		return err
	}

	// no species found
	if ch == nil || len(ch.CharacterSpecies) <= 0 {
		fmt.Fprintln(os.Stdout, "Uknown")
		return nil
	}

	fmt.Fprintln(os.Stdout, ch.CharacterSpecies[0].Name)
	return nil
}
