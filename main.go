package main

import "time"

func main() {
	app := newApp(time.Second * 9)
	app.run(":8080")
}
