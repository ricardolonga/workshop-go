package mongo_test

import (
	"testing"

	"github.com/ricardolonga/workshop-go/domain"
	"github.com/ricardolonga/workshop-go/internal/storage/mongo"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserStorage_Insert(t *testing.T) {
	t.Run("insert successful", func(t *testing.T) {
		databaseName, err := uuid.NewV4()
		assert.NoError(t, err)

		userStorage, err := mongo.NewUserStorage(mongoURL, databaseName.String())
		assert.NoError(t, err)
		assert.NotNil(t, userStorage)

		user := &domain.User{}

		createdUser, err := userStorage.Insert(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
	})
}
