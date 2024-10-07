package folder

import (
"fmt"

"github.com/gofrs/uuid"
"strings"
)
func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...

	folders := f.GetFoldersByOrgID(orgID)
	children := []Folder{}
	parentFound := false

	// Iterate through the filtered folders
	for _, folder := range folders {
		// Check if the current folder is the parent
		if folder.Name == name {
			parentFound = true
		}

		// Check if the folder path starts with the parent folder name followed by a dot
		if strings.HasPrefix(folder.Paths, name+".") {
			children = append(children, folder)
		}
	}

	// If the parent folder is not found, print an error message
	if !parentFound {
		fmt.Printf("Error: Folder '%s' does not exist in organization '%s'\n", name, orgID)
	}

	return children
}
