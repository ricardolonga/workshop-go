package mongo

import (
	"errors"
	"fmt"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

func NewMongoSession(mongoURL string) (*mgo.Session, error) {
	if strings.TrimSpace(mongoURL) == "" {
		return nil, errors.New("mongoURL is required")
	}

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, fmt.Errorf("error on MongoDB connection: %q", err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
