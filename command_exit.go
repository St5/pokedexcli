package main

import "os"


func callbackExit(cnf *config) error{
	os.Exit(0)
	return nil
}