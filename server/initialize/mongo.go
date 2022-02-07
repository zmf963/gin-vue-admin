package initialize

import (
	"context"
	"os"
	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// 初始化数据库并产生数据库全局变量
func MongoInitConnect() *mongo.Database {
	m := global.GVA_CONFIG.Mongo
	if m.MongoDB == "" {
		m.MongoDB = os.Getenv("MONGO_DB")
	}
	if m.MongoURL == "" {
		m.MongoURL = os.Getenv("MONGO_URL")
	}
	global.GVA_LOG.Info("Connecting to MongoDB...", zap.String("MongoDB", m.MongoURL))
	client, err := mongo.NewClient(options.Client().ApplyURI(m.MongoURL))
	if err != nil {
		global.GVA_LOG.Panic("mongo connect error", zap.Error(err))
	}
	err = client.Connect(context.TODO())
	if err != nil {
		global.GVA_LOG.Panic("mongo connect error", zap.Error(err))
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		global.GVA_LOG.Panic("mongo connect error", zap.Error(err))
	}
	return client.Database(m.MongoDB)
}

