package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	orgID := uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"))

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
            name:  "Case with multiple nested folders for orgID",
            orgID: orgID,
            folders: []folder.Folder{
                {Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
                {Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight"},
                {Name: "topical-micromax", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight.topical-micromax"},
                {Name: "bursting-lionheart", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart"},
            },
            want: []folder.Folder{
                {Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
                {Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight"},
                {Name: "topical-micromax", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight.topical-micromax"},
                {Name: "bursting-lionheart", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight.topical-micromax.bursting-lionheart"},
            },
        },
        {
            name:  "Case with no folders for different orgID",
            orgID: uuid.Must(uuid.NewV4()), // Different orgID
            folders: []folder.Folder{
                {Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
            },
            want: []folder.Folder{},
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}

