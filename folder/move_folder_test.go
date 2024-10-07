package folder

import (
	"testing"

	"github.com/gofrs/uuid" // Ensure you import the uuid package
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	// Create UUIDs for organization IDs
	orgID, _ := uuid.NewV4() // Create a new UUID for orgID

	tests := [...]struct {
	name     string
	folders  []Folder
	moveSrc  string
	moveDst  string
	want     []Folder
	wantErr  string
}{
	{
		name: "Move clear-arclight under topical-micromax",
		folders: []Folder{
			{Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
			{Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight"},
			{Name: "topical-micromax", OrgId: orgID, Paths: "creative-scalphunter.topical-micromax"},
			{Name: "bursting-lionheart", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight.bursting-lionheart"},
		},
		moveSrc: "clear-arclight",
		moveDst: "topical-micromax",
		want: []Folder{
			{Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
			{Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.topical-micromax.clear-arclight"},
			{Name: "topical-micromax", OrgId: orgID, Paths: "creative-scalphunter.topical-micromax"},
			{Name: "bursting-lionheart", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight.bursting-lionheart"},
		},
	},
	{
		name: "Move clear-arclight under itself (error)",
		folders: []Folder{
			{Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
			{Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight"},
		},
		moveSrc: "clear-arclight",
		moveDst: "clear-arclight",
		wantErr: "cannot move a folder to itself",
	},
	{
		name: "Move clear-arclight to a non-existent folder (error)",
		folders: []Folder{
			{Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
			{Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight"},
		},
		moveSrc: "clear-arclight",
		moveDst: "invalid_folder",
		wantErr: "destination folder does not exist",
	},
	{
		name: "Invalid source folder (error)",
		folders: []Folder{
			{Name: "creative-scalphunter", OrgId: orgID, Paths: "creative-scalphunter"},
			{Name: "clear-arclight", OrgId: orgID, Paths: "creative-scalphunter.clear-arclight"},
		},
		moveSrc: "invalid_folder",
		moveDst: "topical-micromax",
		wantErr: "source folder does not exist",
	},
}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := &driver{folders: tt.folders}

			var got []Folder
			var err error
			got, err = driver.MoveFolder(tt.moveSrc, tt.moveDst)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
