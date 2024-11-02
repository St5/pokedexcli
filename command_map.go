package main

import (
	"fmt"
	"time"
)

func callbackMap(cnf *config, params ...string) error {
	return renderMap(cnf, false)
}

func callbackMapB(cnf *config, params ...string) error {
	return renderMap(cnf, true)
}

func renderMap(cnf *config, isBack bool) error {
	
	pageLocalURL := cnf.nextLocUrl
	errMsg := "first"
	if(isBack){
		pageLocalURL = cnf.prevLocUrl
		errMsg = "last"
	}

	if(pageLocalURL == nil && (cnf.prevLocUrl != nil || cnf.nextLocUrl != nil)){
		return fmt.Errorf("This is %s page", errMsg)
	}
	start := time.Now()
	Locations, err := cnf.poketapiClient.ListLocationAreas(pageLocalURL)

	elapsed := time.Since(start)

	fmt.Printf("Request taken: %v \n", elapsed)
	if(err != nil){
		return err
	}
	cnf.nextLocUrl = Locations.Next
	cnf.prevLocUrl = Locations.Previous

	fmt.Println("Location areas:")

	for _,area := range Locations.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
 
	return nil
}