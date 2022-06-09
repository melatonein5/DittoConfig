package dittoconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//loadJSONFile is a local function to local a config file
func loadJSONFile(filename string) error {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}

	//Convert the contents of the file into bytes
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	//Interface to prase json into
	var dynamicJson interface{}
	err = json.Unmarshal(jsonFile.)
}
