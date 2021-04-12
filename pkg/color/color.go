/*
  Package to allow for colored text to be outputted to the terminal 
*/

package color

import (
	// Standard
	"runtime"
)

// Variables for all used colors, plus reset to default
var Reset	= "\033[0m"
var Red		= "\033[31m"
var Green	= "\033[32m"
var Yellow	= "\033[33m"
var Blue	= "\033[34m"
var Purple	= "\033[35m"
var Cyan	= "\033[36m"
var Gray	= "\033[37m"
var White	= "\033[97m"

// Makes all colors blank if host is Windows (Windows doesn't support colors by default)
func init() {
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""	
	}		
}