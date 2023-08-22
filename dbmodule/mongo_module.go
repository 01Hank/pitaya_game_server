package dbmodule

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/topfreegames/pitaya/v2/logger"
)

type (
	MongoDBClient struct {
		BaseDB
		conf     *MongoDbConfig
		client   *mongo.Client
		database *mongo.Database
	}

	MongoDbConfig struct {
		DBName string
		MaxNum int
	}
)

// 初始化
func (mClient *MongoDBClient) Init() error {
	err := mClient.Connect()
	if err != nil {
		logger.Log.Warn("init mongo connect fail host: %s, port: %s", mClient.BaseDB.DBHost, mClient.BaseDB.DBPort)
		return err
	}

	mClient.BaseDB.IsConnect = true
	logger.Log.Info("init mongo connect suc host: %s, port: %s", mClient.BaseDB.DBHost, mClient.BaseDB.DBPort)
	return nil
}

// 关服
func (mClient *MongoDBClient) Shutdown() error {
	mClient.DisConnect()
	return nil
}

// 连接mongo
func (mClient *MongoDBClient) Connect() error {
	connectUri := fmt.Sprintf("mongodb://%s:%s/?connect=direct", mClient.BaseDB.DBHost, mClient.BaseDB.DBPort)
	clientOpt := options.Client().ApplyURI(connectUri)
	clientOpt.SetMaxPoolSize(uint64(mClient.conf.MaxNum))

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		logger.Log.Debug("connect mongo init fail!!!")
		return err
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		logger.Log.Debug("connect ping fail!!!")
		return err
	}

	mClient.client = client
	mClient.database = client.Database(mClient.conf.DBName)
	logger.Log.Info("mongo connect is open!!!")
	return nil
}

// 关闭mongo
func (mClient *MongoDBClient) DisConnect() {
	if !mClient.BaseDB.IsConnect {
		return
	}

	mClient.client.Disconnect(context.Background())
	logger.Log.Info("mongo connect is close!!!")
}

// 查询单条
func (mClient *MongoDBClient) FindOne(conllection string, filter *bson.M, result interface{}, opts ...*options.FindOneOptions) {
	collec := mClient.database.Collection(conllection)
	collec.FindOne(context.Background(), filter, opts...).Decode(result)
}

// 查询多条
func (mClient *MongoDBClient) FindMany(ctx context.Context, conllection string, filter *bson.M, results *[]interface{}, opts ...*options.FindOptions) error {
	collec := mClient.database.Collection(conllection)
	cur, err := collec.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	return cur.All(ctx, results)
}

// 保存单条数据
func (mClient *MongoDBClient) SaveOne(conllection string, data interface{}, opts ...*options.InsertOneOptions) error {
	collec := mClient.database.Collection(conllection)
	_, err := collec.InsertOne(context.Background(), data, opts...)
	return err
}

// 保存多条
func (mClient *MongoDBClient) SaveMany(ctx context.Context, conllection string, dataList []interface{}, opts ...*options.InsertManyOptions) error {
	collec := mClient.database.Collection(conllection)
	_, err := collec.InsertMany(ctx, dataList, opts...)
	return err
}

// 更新一条数据
func (mClient *MongoDBClient) UpdateOne(conllection string, filter *bson.M, upData interface{}, opts ...*options.UpdateOptions) error {
	collec := mClient.database.Collection(conllection)
	_, err := collec.UpdateOne(context.Background(), filter, upData, opts...)
	return err
}

// 根据条件删除多条
func (mClient *MongoDBClient) DeleteMany(ctx context.Context, conllection string, filter *bson.M, opts ...*options.DeleteOptions) error {
	collec := mClient.database.Collection(conllection)
	_, err := collec.DeleteMany(ctx, filter, opts...)
	return err
}

// 构建mongo
func BuildMongo(host, port, dbName string, maxNum int) *MongoDBClient {
	base := BaseDB{
		DBHost: host,
		DBPort: port,
	}

	conf := &MongoDbConfig{
		DBName: dbName,
		MaxNum: maxNum,
	}

	return &MongoDBClient{
		BaseDB: base,
		conf:   conf,
	}
}
