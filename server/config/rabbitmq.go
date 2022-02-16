/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-11 17:13:33
 * @LastEditors: zmf96
 * @LastEditTime: 2022-02-11 17:14:34
 * @FilePath: /server/config/rabbitmq.go
 * @Description:
 */
package config

type Rabbitmq struct {
	RabbitmqUri string `mapstructure:"rabbitmq-uri" json:"rabbitmqUri" yaml:"rabbitmq-uri"` // rabbitmq的uri
}
