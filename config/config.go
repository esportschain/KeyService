package config

type Conf struct  {
	BindIp string `json:"ip"`	// 简体的IP
	Port int `json:"port"`       // 监听的端口
	Version string `json:"version"` // 应用版本
}


