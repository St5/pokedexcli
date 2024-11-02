package main

import "os"


func callbackExit(cnf *config, params ...string) error{
	os.Exit(0)
	return nil
}