package prefs

type Prefs struct {
	Interface Interface
	Controls  Controls
	Editor    Editor
}

type Interface struct {
	Scale int
	Fps   int
}

type Controls struct {
	AltScrollBehaviour   bool
	QuickEditContextMenu bool
	QuickEditMapPane     bool
}

type Editor struct {
	SaveFormat        string
	NudgeMode         string
	SanitizeVariables bool
}
