package mongo_test

import (
	"testing"

	"github.com/NeowayLabs/logger"
	"github.com/ricardolonga/workshop-go/internal/storage/mongo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

var (
	mongoURL string
	session  *mgo.Session
)

func init() {
	mongoURL = "172.26.0.2" // os.Getenv("MONGO_URL")
	if mongoURL != "" {
		if mongoSession, err := mgo.Dial(mongoURL); err != nil {
			logger.Fatal("mongo connection failed: %q", err)
		} else {
			session = mongoSession
		}
	}
}

func TestNewMongoSession(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mongoSession, err := mongo.NewMongoSession(mongoURL)
		assert.NoError(t, err)
		assert.NotNil(t, mongoSession)
	})

	t.Run("empty_url", func(t *testing.T) {
		mongoSession, err := mongo.NewMongoSession("")
		assert.Error(t, err)
		assert.Nil(t, mongoSession)
	})

	t.Run("invalid_url", func(t *testing.T) {
		mongoSession, err := mongo.NewMongoSession("192.168.123.123")
		assert.Error(t, err)
		assert.Nil(t, mongoSession)
	})
}
