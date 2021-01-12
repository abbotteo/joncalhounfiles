package main

import (
	"fmt"
)

func main() {
	/*
	This comment will remain as commit 'one'
	 */
	fmt.Print("The executable\n")
	/*
		this comment will be added to develop, then reverted, then
		this branch will be rebased and check if the change is lost
		The change was lost, so I move the pointer to 0bc89c99 which had
		the comment. But I don't know if I have the changes in develop
	*/
}
