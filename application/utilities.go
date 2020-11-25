// Copyright 2020 Adegbenga Adeye. All rights reserved.
// Use of this source code is governed by a GNU GENERAL PUBLIC LICENSE
// that can be found in the LICENSE file.

/*
	Package application implements all needed functionalities in the application using a DDD programming style.
*/

package application

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	entity "jarvisESB/domain/entity"

	"github.com/joho/godotenv"
)

type model entity.ResponseMessage

// GetEnv returns the value based on the environment key passed in
func GetEnv(key string) string {
	return os.Getenv(key)
}

// BasePath gets and returns the rooted path name corresponding to the current directory
func BasePath() string {
	cwd := ""
	var err error

	//  Returns a rooted path name corresponding to the current directory
	if cwd, err = os.Getwd(); err != nil {
		CheckErr("Unable to get Application Working Directory.", err, 5)
	}

	return cwd
}

// GetEnvironment loads the preferred environment file
func GetEnvironment(fileName string) {
	var err error

	// load the .env file
	if err = godotenv.Load(filepath.Join(BasePath(), ".env")); err != nil {
		CheckErr("Error loading .env file", err, 5)
	}
}

// ErrorMessage returns a standard Error response
func (response *model) ErrorMessage() ([]byte, error) {
	response.Success = "false"
	return json.Marshal(response)
}

// SuccessMessage returns a standard Success response
func (response *model) SuccessMessage() ([]byte, error) {
	response.Success = "true"
	response.Message = "Data was successfully processed via the datastore."
	return json.Marshal(response)
}

// CheckErr checks if the err is not nil and panics
func CheckErr(message string, err error, level int) {

	if err == nil {
		return
	}

	// level 0 is for testing/we can ignore
	// level 2 is for minor file issues [None show stoppers e.g unable to load a json file]
	// level 3 is http error [We can recover from them gracefully]
	// level 4 is a fatal error (log.fatal), database errors, conversion errors
	// level 5 is critical and panics

	// file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(file)
	// log.Println(message + " > " + err.Error())

	if err != nil {
		switch level {
		case 0:
		case 1:
		case 2:
		case 3:
		case 4:
			log.Println(message + " > " + err.Error())
			// fmt.Errorf("%v > %v", message, err.Error())
		case 5:
			panic(err)
		}
	}
}
