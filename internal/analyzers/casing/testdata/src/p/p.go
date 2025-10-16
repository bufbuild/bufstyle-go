package p

func foo() {
	//fileName := ""
	//_ = fileName
	//filename := "" // `Use "fileName" instead of "filename"`
	//_ = filename   // `Use "fileName" instead of "filename"`
	filePath := ""
	_ = filePath
	filepath := "" // `Use "filePath" instead of "filepath"`
	_ = filepath   // `Use "filePath" instead of "filepath"`
	dirName := ""
	_ = dirName
	dirname := "" // want `Use "dirName" instead of "dirname"`
	_ = dirname   // want `Use "dirName" instead of "dirname"`
	dirPath := ""
	_ = dirPath
	dirpath := "" // want `Use "dirPath" instead of "dirpath"`
	_ = dirpath   // want `Use "dirPath" instead of "dirpath"`
	FileName := ""
	_ = FileName
	Filename := "" // `Use "FileName" instead of "Filename"`
	_ = Filename   // `Use "FileName" instead of "Filename"`
	FilePath := ""
	_ = FilePath
	Filepath := "" // `Use "FilePath" instead of "Filepath"`
	_ = Filepath   // `Use "FilePath" instead of "Filepath"`
	DirName := ""
	_ = DirName
	Dirname := "" // want `Use "DirName" instead of "Dirname"`
	_ = Dirname   // want `Use "DirName" instead of "Dirname"`
	DirPath := ""
	_ = DirPath
	Dirpath := "" // want `Use "DirPath" instead of "Dirpath"`
	_ = Dirpath   // want `Use "DirPath" instead of "Dirpath"`
}
