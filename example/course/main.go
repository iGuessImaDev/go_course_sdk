package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/iGuessImaDev/go_course_sdk/course"
)

func main() {
	courseTrans := courseSdk.NewHTTPClient("http://localhost:8082", "")

	course, err := courseTrans.Get("d370f50d-ffdb-4f0e-b5f5-fe6019646e31")
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error", err.Error())
		os.Exit(1)
	}

	fmt.Println(course)
}
