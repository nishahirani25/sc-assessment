package folder

import (
	"errors"
	"strings"
)

// MoveFolder moves a folder from one parent to another.
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	var sourceFolder, destinationFolder Folder
	var sourceFound, destinationFound bool

	// Find source and destination folders
	for _, folder := range f.folders {
		if folder.Paths == name { // Use Paths instead of paths
			sourceFolder = folder
			sourceFound = true
		}
		if folder.Paths == dst {
			destinationFolder = folder
			destinationFound = true
		}
	}

	// Check if source folder exists
	if !sourceFound {
		return nil, errors.New("source folder does not exist")
	}

	// Check if destination folder exists
	if !destinationFound {
		return nil, errors.New("destination folder does not exist")
	}

	// Check for self-move or child-move
	if strings.HasPrefix(sourceFolder.Paths, destinationFolder.Paths) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}
	if sourceFolder.Paths == destinationFolder.Paths {
		return nil, errors.New("cannot move a folder to itself")
	}
	if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	// Prepare new path for the source folder
	newPath := destinationFolder.Paths + "." + sourceFolder.Name
	sourceFolder.Paths = newPath

	// Update paths of all child folders
	var updatedFolders []Folder
	for _, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, sourceFolder.Paths) {
			newFolderPath := newPath + strings.TrimPrefix(folder.Paths, sourceFolder.Paths)
			folder.Paths = newFolderPath
		}
		updatedFolders = append(updatedFolders, folder)
	}

	// Remove the original source folder and add the updated one
	for i, folder := range updatedFolders {
		if folder.Paths == sourceFolder.Paths {
			updatedFolders = append(updatedFolders[:i], updatedFolders[i+1:]...)
			break
		}
	}
	updatedFolders = append(updatedFolders, sourceFolder)

	return updatedFolders, nil
}
