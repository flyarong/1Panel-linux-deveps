package dto

import "time"

type SSHInfo struct {
	Status                 string `json:"status"`
	Message                string `json:"message"`
	Port                   string `json:"port"`
	ListenAddress          string `json:"listenAddress"`
	PasswordAuthentication string `json:"passwordAuthentication"`
	PubkeyAuthentication   string `json:"pubkeyAuthentication"`
	PermitRootLogin        string `json:"permitRootLogin"`
	UseDNS                 string `json:"useDNS"`
}

type GenerateSSH struct {
	EncryptionMode string `json:"encryptionMode" validate:"required,oneof=rsa ed25519 ecdsa dsa"`
	Password       string `json:"password"`
}

type GenerateLoad struct {
	EncryptionMode string `json:"encryptionMode" validate:"required,oneof=rsa ed25519 ecdsa dsa"`
}

type SSHConf struct {
	File string `json:"file"`
}
type SearchSSHLog struct {
	PageInfo
	Info   string `json:"info"`
	Status string `json:"Status" validate:"required,oneof=Success Failed All"`
}
type SSHLog struct {
	Logs            []SSHHistory `json:"logs"`
	TotalCount      int          `json:"totalCount"`
	SuccessfulCount int          `json:"successfulCount"`
	FailedCount     int          `json:"failedCount"`
}

type SearchForAnalysis struct {
	OrderBy string `json:"orderBy" validate:"required,oneof=Success Failed"`
}

type SSHLogAnalysis struct {
	Address         string `json:"address"`
	Area            string `json:"area"`
	SuccessfulCount int    `json:"successfulCount"`
	FailedCount     int    `json:"failedCount"`
	Status          string `json:"status"`
}

type SSHHistory struct {
	Date     time.Time `json:"date"`
	DateStr  string    `json:"dateStr"`
	Area     string    `json:"area"`
	User     string    `json:"user"`
	AuthMode string    `json:"authMode"`
	Address  string    `json:"address"`
	Port     string    `json:"port"`
	Status   string    `json:"status"`
	Message  string    `json:"message"`
}
