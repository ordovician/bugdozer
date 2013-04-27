// Copyright 2013 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Dumps out a pretty printed version of JSON
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	var obj interface{}
	
	json.Unmarshal(input, &obj)
	
	bytes, _ := json.MarshalIndent(obj, "", "    ")
	fmt.Print(string(bytes))
}
