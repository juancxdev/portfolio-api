package main

import (
	"portfolio-api/api"
	"portfolio-api/internal/env"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	e := env.NewConfiguration()
	api.Start(e.Port)
}
