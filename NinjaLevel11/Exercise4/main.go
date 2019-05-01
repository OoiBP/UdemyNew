package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		err := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  errors.New("You are a failure"),
		}
		return 0, err
	}
	return 42, nil
}

/*
Starting with this code, use the sqrt.Error struct as a value of type error. If you would like, use these numbers for your
	- lat "50.2289 N"
	- long "99.4656 W"
*/
