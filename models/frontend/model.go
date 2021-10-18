package frontend

import "goProject/models/user"

type GetDataNumResponseBody struct {
	DataNum uint `json:"dataNum"`
}

type DataList [](user.Data)

type GetDataListResponseBody struct {
	State bool     `json:"state"`
	Data  DataList `json:"data"`
}

type GetDownloadRequestBody struct {
	Address string `json:"address"`
	Target  int    `json:"target"`
	Id      string `json:"id"` // 即 md5
	IP      string `json:"ip"`
	Type    int    `json:"type"`
}

type GetDownloadResponseBody struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Md5      string `json:"md5"`
}

type UploadToDfsResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Md5     string `json:"md5"`
}

type DeleteRequestBody struct {
	Md5 string `json:"md5"`
}
