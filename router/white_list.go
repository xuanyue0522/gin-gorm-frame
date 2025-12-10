package router

// AdminAuthWhiteList 管理后台认证白名单
var AdminAuthWhiteList = map[string]bool{
	"/ping":                                  true,
	"/metrics":                               true,
	"/admindo/v1/verify/captcha/check":       true,
	"/admindo/v1/verify/captcha":             true,
	"/admindo/v1/user/verify/smscode":        true,
	"/admindo/v1/user/mobile/verify_login":   true,
	"/admindo/v1/user/mobile/password_login": true,
	"/admindo/v1/user/password/reset":        true,
}
