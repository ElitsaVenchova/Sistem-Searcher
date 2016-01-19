package system.search

import (

)

type Search interface{
	//Set searched text
	func SetSearchedText(text string)https://www.youtube.com/watch?v=KcEBNIH6AhY&ab_channel=TU154
	// Retutn searched text
	func GetSearchedText() string;
	// Return is seached text a regular expression
	func IsRegex() bool;
	//Set is seached text a regular expression 
	func SetIsGegex(bool) err error;
	//Return root path like a string
	func GetRoot() string;
	//Set root path
	func SetRoot(root string) err error;
	//Set max search depth (0- only in current dir)
	func SetSearchDepth(int depth) err error;
	//This start seach after read content in previouse fields and after thas return result set like slice of files
	func GetResultSet() ([]File, map[]);
	
}
