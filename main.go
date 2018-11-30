package main

import "malu/client"

func main() {
	client.Start()
	<-make(chan (struct{}))
}
