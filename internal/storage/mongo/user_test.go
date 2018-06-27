package mongo_test

import (
	"testing"

	"github.com/ricardolonga/workshop-go/domain"
	"github.com/ricardolonga/workshop-go/internal/storage/mongo"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestExtraStorage(t *testing.T) {
	t.Run("insert", func(t *testing.T) {
		databaseName, err := uuid.NewV4()
		assert.NoError(t, err)

		extraStorage, err := mongo.NewUserStorage(mongoURL, databaseName.String())
		assert.NoError(t, err)
		assert.NotNil(t, extraStorage)

		user := &domain.User{}

		createdExtra, err := extraStorage.Insert(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdExtra)
	})
}
