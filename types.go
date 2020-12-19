package main

const (
	// AppName 应用名称.
	AppName = "gitee"
	// Domain 域.
	Domain = "https://gitee.com"
)

const (
	DeleteCommand = "delete"
)

const (
	ConfigFlag = "config"
)

// AccountInfo 账号信息.
type AccountInfo struct {
	// Token X-CSRF-Token.
	Token string
	// Repositories 仓库地址，英文逗号分隔
	Repositories string
	// 账号名称
	UserName string
	// 账号密码
	Password string
	// Cookie
	Cookie string
}

// Account 账号信息.
var Account AccountInfo

const (
	// HttpTypePost3W 请求
	HttpTypePost3W = "application/x-www-form-urlencoded"
)
