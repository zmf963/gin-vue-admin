/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-10 12:41:39
 * @LastEditors: zmf96
 * @LastEditTime: 2022-02-10 12:46:14
 * @FilePath: /server/config/captcha_key.go
 * @Description:
 */
package config

type CaptchaKey struct {
	CaptchaId string `mapstructure:"captcha-id" json:"captchaId" yaml:"captcha-id"` // 固定的验证码ID
	Captcha  string `mapstructure:"captcha" json:"captcha" yaml:"captcha"`           // 固定的验证码
}
