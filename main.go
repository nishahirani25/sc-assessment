// package main
//
// import (
// 	"fmt"
//
// 	"github.com/georgechieng-sc/interns-2022/folder"
// 	"github.com/gofrs/uuid"
// )
//
// func main() {
// 	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
//
// 	res := folder.GetAllFolders()
//
// 	// example usage
// 	folderDriver := folder.NewDriver(res)
// 	orgFolder := folderDriver.GetFoldersByOrgID(orgID)
//
// 	folder.PrettyPrint(res)
// 	fmt.Printf("\n Folders for orgID: %s", orgID)
// 	folder.PrettyPrint(orgFolder)
// }


package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID) // Assuming DefaultOrgID is a const

	res := folder.GetAllFolders() // Fetch all folders

	// Create a folder driver
	folderDriver := folder.NewDriver(res)

	// Get folders by orgID
	orgFolders := folderDriver.GetFoldersByOrgID(orgID)

	// Pretty print all folders
	folder.PrettyPrint(res)

	fmt.Printf("\nFolders for orgID: %s\n", orgID)
	folder.PrettyPrint(orgFolders)

	// Create a reader to take user input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the folder name you want to query (or type 'exit' to quit): ")
		folderName, _ := reader.ReadString('\n')
		folderName = folderName[:len(folderName)-1] // Remove newline character

		if folderName == "exit" {
			break
		}

		// Get child folders for the provided folder name
		childFolders := folderDriver.GetAllChildFolders(orgID, folderName)
		fmt.Printf("\nChild Folders for '%s':\n", folderName)
		folder.PrettyPrint(childFolders)
	}
}


