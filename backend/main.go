package main

import (
	"github.com/fexcel/fexcel-backend/Initialize"
)

func main() {
	Initialize.Load_env()
	Initialize.Init_gpt()
	Initialize.Setup_db()
	Initialize.Init_fiber()
}
