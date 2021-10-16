package goFastdfs

type Status struct {
	Message string     `json:"message"`
	Status  string     `json:"status"`
	Data    []DateInfo `json:"data"`
}

type DateInfo struct {
	Date      string `json:"date"`
	FileCount int64  `json:"fileCount"`
	TotalSize int64  `json:"totalSize"`
}

type UploadResult struct {
	Domain     string  `json:"domain"`
	Mtime      float64 `json:"mtime"`
	ReturnMsg  string  `json:"retmsg"`
	ReturnCode float64 `json:"retcode"`
	Src        string  `json:"src"`
	Url        string  `json:"url"`
	Md5        string  `json:"md5"`
	Path       string  `json:"path"`
	Scene      string  `json:"scene"`
	Size       int64   `json:"size"`
	Scenes     string  `json:"scenes"`
}

type FileInfo struct {
	Name      string   `json:"name"`
	Rename    string   `json:"rename"`
	Path      string   `json:"path"`
	Md5       string   `json:"md5"`
	Size      int64    `json:"size"`
	Peers     []string `json:"peers"`
	Scene     string   `json:"scene"`
	TimeStamp int64    `json:"timeStamp"`
	Offset    int64    `json:"offset"`
}

type FileInfoResult struct {
	Data    FileInfo `json:"data"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
}

type DeleteResult struct {
	Data    FileInfo `json:"data"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
}

type DirInfo struct {
	Name  string `json:"name"`
	Md5   string `json:"md5"`
	Path  string `json:"path"`
	Size  int64  `json:"size"`
	Mtime int64  `json:"mtime"`
	IsDir bool   `json:"is_dir"`
}

type DirInfoResult struct {
	Data    []DirInfo `json:"data"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
}
