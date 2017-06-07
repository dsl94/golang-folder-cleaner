# golang-folder-cleaner
This is simple folder cleaner written in Go lnaguage
Usage instructions:
1. Configure folders you want to clear here:
    var folders = [...]string{"/Users/Nemanja/Downloads", "/Users/Nemanja/Documents/ViberDownloads"}
2. Configure files/folders you want to skip, this can be part of file/folder name or extension
    var skip = [...]string{"Udemy"}
3. Run script
