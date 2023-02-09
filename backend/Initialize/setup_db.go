package Initialize

import (
	"fmt"
	"github.com/fexcel/fexcel-backend/db"
)

func Setup_db() {
	err := db.DBsetup()
	if err != nil {
		fmt.Println(err)
		return
	}

}
