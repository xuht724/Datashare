package blockchain

type VerifyResult struct {
	State bool `json:"state"`
}

type VerifyRequestBody struct {
	Address string `json:"address"`
	Target  int    `json:"target"`
	Ip      string `json:"ip"`
	Type    int    `json:"type"`
}

type AddLogRequestBody struct {
	Address string `json:"address"`
	Target  int    `json:"target"`
	Ip      string `json:"ip"`
	Time    string `json:"time"`
}

type AddLogResult struct {
	State   bool   `json:"state"`
	Message string `json:"message"`
}

type Log struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Ip   string `json:"ip"`
}

type GetLogsRequestBody struct {
	SerialNum int `json:"serialNum"`
}

type GetLogsResult struct {
	Num  string `json:"num"`
	Logs []Log  `json:"Logs"`
}
