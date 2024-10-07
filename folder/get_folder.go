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

	folders := f.GetFoldersByOrgID(orgID)
	children := []Folder{}
	parentFound := false

	for _, folder := range folders {
		if folder.Name == name {
			parentFound = true
		}

		if strings.HasPrefix(folder.Paths, name+".") {
			children = append(children, folder)
		}
	}

	if !parentFound {
		fmt.Printf("Error: Folder '%s' does not exist in organization '%s'\n", name, orgID)
	}

	return children
}
