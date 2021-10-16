package goFastdfs

import (
	"fmt"
	"testing"
)

func TestUploadFile(t *testing.T) {
	const filename = "../../user.json"
	const hostURL = "http://localhost:8085/group1"
	var response = UploadFile(filename, hostURL)
	fmt.Printf("%#v", response)
}

func TestGetStatus(t *testing.T) {
	const hostURL = "http://localhost:8085/group1/stat"
	var response = GetStatus(hostURL)
	fmt.Printf("%+v", response)
}

func TestDeleteOneFile(t *testing.T) {
	const filename = "../login.html"
	const uploadURL = "http://localhost:8085/group1/upload"
	var response = UploadFile(filename, uploadURL)
	fmt.Printf("%+v\n", response)
	var md5 = response.Md5
	const hostURL = "http://localhost:8085/group1/delete"
	var deleteResult = DeleteOneFile(hostURL, md5)
	fmt.Printf("%+v", deleteResult)
}

func TestGetFileInfo(t *testing.T) {
	const md5 = "6054b52e6981f9960fcf334b0ddb72e91"
	const hostURL = "http://localhost:8085/group1/get_file_info"
	var response = GetFileInfo(hostURL, md5)

	fmt.Printf("%+v\n", response)
}

func TestGetDirInfo(t *testing.T) {
	const dir = ""
	const hostURL = "http://localhost:8085/group1/list_dir"
	var response = GetDirInfo(hostURL, dir)

	fmt.Printf("%+v\n", response)
}

func TestGetDownloadLink(t *testing.T) {
	const md5 = "6054b52e6981f9960fcf334b0ddb72e9"
	const hostURL = "http://localhost:8085/group1/"

	// fmt.Println(hostURL[6:] + "1")
	fmt.Println(GetDownloadURL(hostURL, md5))
}

func TestDownloadFile(t *testing.T) {
	const md5 = "6054b52e6981f9960fcf334b0ddb72e9"
	const hostURL = "http://localhost:8085/group1/"
	const savePath = "../../"

	DownloadFile(hostURL, md5, savePath)
}
