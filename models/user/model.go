package user

type Config struct {
	BlockchainURL string `json:"blockchainURL"`
	DfsURL        string `json:"dfsURL"`
	FilePath      string `json:"filePath"`
}

type Data struct {
	Md5      string `json:"md5"`
	Filename string `json:"filename"`
	Time     string `json:"time"`
	Type     uint   `json:"type"`
}

type User struct {
	Username  string `json:"username"`
	PublicKey string `json:"publicKey"`
	Data      []Data `json:"data"`
}

var DefaultConfig = Config{
	BlockchainURL: "http://localhost:12345",
	DfsURL:        "http://localhost:8080/group1",
	FilePath:      "./files/",
}

var DefaultUserInfo = User{
	Username:  "test",
	PublicKey: "test key",
	Data:      []Data{},
}

const ConfigFile = "./config.json"
const UserFile = "./user.json"
