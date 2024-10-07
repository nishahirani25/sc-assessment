package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	var sourceFolder, destinationFolder Folder
	var sourceFound, destinationFound bool

	for _, folder := range f.folders {
		if folder.Paths == name { 
			sourceFolder = folder
			sourceFound = true
		}
		if folder.Paths == dst {
			destinationFolder = folder
			destinationFound = true
		}
	}

	if !sourceFound {
		return nil, errors.New("source folder does not exist")
	}

	if !destinationFound {
		return nil, errors.New("destination folder does not exist")
	}

	if strings.HasPrefix(sourceFolder.Paths, destinationFolder.Paths) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}
	if sourceFolder.Paths == destinationFolder.Paths {
		return nil, errors.New("cannot move a folder to itself")
	}
	if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	newPath := destinationFolder.Paths + "." + sourceFolder.Name
	sourceFolder.Paths = newPath

	var updatedFolders []Folder
	for _, folder := range f.folders {
		if strings.HasPrefix(folder.Paths, sourceFolder.Paths) {
			newFolderPath := newPath + strings.TrimPrefix(folder.Paths, sourceFolder.Paths)
			folder.Paths = newFolderPath
		}
		updatedFolders = append(updatedFolders, folder)
	}

	for i, folder := range updatedFolders {
		if folder.Paths == sourceFolder.Paths {
			updatedFolders = append(updatedFolders[:i], updatedFolders[i+1:]...)
			break
		}
	}
	updatedFolders = append(updatedFolders, sourceFolder)

	return updatedFolders, nil
}
