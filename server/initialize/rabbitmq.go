/*
 * @Version: 0.1
 * @Autor: zmf96
 * @Email: zmf96@qq.com
 * @Date: 2022-02-11 17:10:24
 * @LastEditors: zmf96
 * @LastEditTime: 2022-02-11 17:38:10
 * @FilePath: /server/initialize/rabbitmq.go
 * @Description:
 */
package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/streadway/amqp"
)

func RabbitmqInit() *amqp.Channel {
	// 初始化rabbitmq
	conn, err := amqp.Dial(global.GVA_CONFIG.Rabbitmq.RabbitmqUri)
	if err != nil {
		panic(err)
	}
	// 初始化rabbitmq的channel
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	_, err = channel.QueueDeclare("server:default", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	return channel
}
