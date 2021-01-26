package app

import "github.com/SpaiR/strongdmm/internal/app/byond/dmm"

func (a *app) openMap(path string) {
	dmmData, _ := dmm.NewData(path)
	a.data.AddRecentMap(a.loadedEnvironment.RootFilePath, path)
	println(dmmData)
}
