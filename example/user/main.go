package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/iGuessImaDev/go_course_sdk/user"
)

func main() {
	userTrans := userSdk.NewHTTPClient("http://localhost:8081", "")

	user, err := userTrans.Get("c0468641-64e9-47b2-afasdf-b8bebfab0f61")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
