package configs

import "github.com/hisitra/confine/v2"

var isLoaded = false

// Get : Loads the configs if not already loaded, then returns them.
// Panics if loading fails.
func Get() *Configs {
	if isLoaded {
		return conf
	}

	err := confine.LoadMany(confineMap)
	if err != nil {
		panic(err)
	}

	isLoaded = true
	return conf
}
