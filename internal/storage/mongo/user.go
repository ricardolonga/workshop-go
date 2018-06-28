package mongo

import (
	"github.com/ricardolonga/workshop-go/domain"
	"github.com/ricardolonga/workshop-go/internal"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userStorage struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
}

const collectionUsers = "users"

func NewUserStorage(mongoURL, databaseName string) (*userStorage, error) {
	session, err := NewMongoSession(mongoURL)
	if err != nil {
		return nil, err
	}

	return &userStorage{
		session:        session,
		databaseName:   databaseName,
		collectionName: collectionUsers,
	}, nil
}

func (es *userStorage) Insert(user *domain.User) (*domain.User, error) {
	session := es.session.Copy()
	defer session.Close()

	user.ID = bson.NewObjectId().Hex()
	if err := session.DB(es.databaseName).C(es.collectionName).Insert(user); err != nil {
		if mgo.IsDup(err) {
			return nil, internal.NewDuplicatedRecordError("user")
		}

		return nil, err
	}

	return user, nil
}

func (es *userStorage) Delete(userID string) error {
	session := es.session.Copy()
	defer session.Close()

	if err := session.DB(es.databaseName).C(es.collectionName).RemoveId(userID); err != nil {
		if err == mgo.ErrNotFound {
			return internal.NewNotFoundError("user")
		}

		return err
	}

	return nil
}

func (es *userStorage) GetAll() ([]*domain.User, error) {
	session := es.session.Copy()
	defer session.Close()

	users := make([]*domain.User, 0)
	if err := session.DB(es.databaseName).C(es.collectionName).Find(nil).All(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (es *userStorage) GetByID(userID string) (*domain.User, error) {
	session := es.session.Copy()
	defer session.Close()

	user := &domain.User{}
	if err := session.DB(es.databaseName).C(es.collectionName).FindId(userID).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return nil, internal.NewNotFoundError("user")
		}

		return nil, err
	}

	return user, nil
}
