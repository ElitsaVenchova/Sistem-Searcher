package system.tree

import (

)

type SystemTree interface{
	//Return root path like a string
	func GetRoot() string;
	//Set root path
	func SetRoot(root string) err error;
	//Get tree
	func GetTree();
	//Set tree
	func SetTree();
}
