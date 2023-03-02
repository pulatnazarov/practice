package recursive

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	title  = "title "
	parent = "parent"
)

var (
//repeatedIDs = make(map[string]bool, 0)
//mapIDs      = make(map[string]bool, 0)
//ids         = make([]string, 0)
)

type Folder struct {
	Id             string
	title          string
	parentFolderID string
	childFolders   []Folder
	childFolderIDs []string
}

func RandomUUID() string {
	return uuid.New().String()
}

// Recursive function
func Recursive(n int) []Folder {
	x := 0
	if n == 1 {
		x = 50
	} else if n == 2 {
		x = 30
	} else if n == 3 {
		x = 20
	} else if n == 4 {
		x = 20
	} else {
		return nil
	}
	folders := make([]Folder, 0)
	for i := 0; i < x/10; i++ {
		childFolders1 := make([]Folder, 0)
		folders = append(folders, Folder{
			Id:             RandomUUID(),
			title:          fmt.Sprintf("%s%d", title, i+x),
			parentFolderID: parent,
			childFolders:   childFolders1,
			childFolderIDs: []string{},
		})

		//if _, ok := mapIDs[folders[i].id]; !ok {
		//	mapIDs[folders[i].id] = true
		//}
		childFolders := Recursive(n + 1)

		//for key := range mapIDs {
		//	if key != folders[i].id {
		//		ids = append(ids, key)
		//	} else {
		//		repeatedIDs[key] = true
		//	}
		//}
		//
		//for j, id := range ids {
		//	_, ok := repeatedIDs[id]
		//	if ok {
		//		if j < len(ids) {
		//			ids = append(ids[:j], ids[j+1:]...)
		//		}
		//	}
		//}

		folders[i].childFolders = childFolders
		//folders[i].childFolderIDs = ids
	}
	//ids = make([]string, 0)
	//repeatedIDs = make(map[string]bool, 0)
	//mapIDs = make(map[string]bool, 0)

	return folders
}

func New(id string) {
	for {

	}
}

func Print(folders []Folder) {
	for i, folder := range folders {
		fmt.Println()
		fmt.Println((i + 1) * 100)
		fmt.Println()
		fmt.Println("folder id: ", folder.Id)
		fmt.Println("folder title: ", folder.title)
		fmt.Println("folder parentFolderID: ", folder.parentFolderID)
		fmt.Println("folder childFolderIDs: ", folder.childFolderIDs, len(folder.childFolderIDs))
		for _, childFolder := range folder.childFolders {
			fmt.Println()
			fmt.Println("			id: ", childFolder.Id)
			fmt.Println("			title: ", childFolder.title)
			fmt.Println("			parentFolderID: ", childFolder.parentFolderID)
			fmt.Println("			childFolderIDs: ", childFolder.childFolderIDs, len(childFolder.childFolderIDs))
			fmt.Println()
			for _, childFolder2 := range childFolder.childFolders {
				fmt.Println()
				fmt.Println("						id: ", childFolder2.Id)
				fmt.Println("						title: ", childFolder2.title)
				fmt.Println("						parentFolderID: ", childFolder2.parentFolderID)
				fmt.Println("						childFolderIDs: ", childFolder2.childFolderIDs, len(childFolder2.childFolderIDs))
				fmt.Println()
				for _, childFolder3 := range childFolder2.childFolders {
					fmt.Println()
					fmt.Println("									id: ", childFolder3.Id)
					fmt.Println("									title: ", childFolder3.title)
					fmt.Println("									parentFolderID: ", childFolder3.parentFolderID)
					fmt.Println("									childFolderIDs: ", childFolder3.childFolderIDs, len(childFolder3.childFolderIDs))
					fmt.Println()
				}
			}
		}
		fmt.Println()
	}
}

/*folders := make([]Folder, 0)
for i := 0; i < x; i++ {
	childFolders1 := make([]Folder, 0)
	folders = append(folders, Folder{
		id:             RandomUUID(),
		title:          fmt.Sprintf("%s%d", title, i+11),
		parentFolderID: parent,
		childFolders:   childFolders1,
		childFolderIDs: []string{},
	})

	childFolders2 := make([]Folder, 0)
	for j := 0; j < 3; j++ {
		childFolders1 = append(childFolders1, Folder{
			id:             RandomUUID(),
			title:          fmt.Sprintf("%s%d", title, j+21),
			parentFolderID: folders[i].id,
			childFolders:   []Folder{},
			childFolderIDs: []string{},
		})

		for k := 0; k < 2; k++ {
			childFolders2 = append(childFolders2, Folder{
				id:             RandomUUID(),
				title:          fmt.Sprintf("%s%d", title, k+31),
				parentFolderID: childFolders1[j].id,
				childFolders:   []Folder{},
			})
		}
		childFolders1[j].childFolders = childFolders2

	}
	folders[i].childFolders = childFolders1
}*/
