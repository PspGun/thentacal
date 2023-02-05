package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PspGun/thentacal/db"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := db.DBsetup()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// dsn := "host=beta.pspgun.com user=fexel password=fexel1212312121 dbname=postgres-test port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	
	// 	if err != nil {
	// 	panic(err)
	// 	}else {
    //     	fmt.Println("Connecting successful")
    //     }
	// userRepository := repository.NewRepositoryDB(db)
	// userService := service.NewUserService(userRepository)
	// userHandler := hd.NewUserHandler(userService)

	// http.HandleFunc("/signup", userHandler.SignUp)
	// http.HandleFunc("/signin", userHandler.SignIn)
	// http.HandleFunc("/listuser", userHandler.ListUsers)

	http.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
		response_value := map[string]any{"Message": "Hello, World"}
		response, _ := json.Marshal(response_value)
		w.Write(response)
	}))

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}