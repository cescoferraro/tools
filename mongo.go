package tools

import (
	"time"
	"gopkg.in/mgo.v2"
	"strconv"
	"github.com/cescoferraro/tools/logger"
)



//MONGO local MongoObject
var StoreObject MongoStore

type MongoStore struct {
	Database     string
	AuthDatabase string
	Session      *mgo.Session
	User         string
	Password     string
	Host         string
	Port         int
}

var mongoLogger = logger.New("MONGO")

//Cloner clones the local MONGO object
func Cloner() *MongoStore {
	return &MongoStore{
		User:StoreObject.User,
		Password:StoreObject.Password,
		Host:StoreObject.Host,
		Port:StoreObject.Port,
		AuthDatabase: StoreObject.AuthDatabase,
		Database: StoreObject.Database,
		Session: StoreObject.Session.Copy(),
	}
}

//InitMongo set MongoDB connection
func (store MongoStore) Init(done chan bool) {
	mongoLogger.Print("Attempting to get MongoDB session")
	var err error
	store , err = store.connect()
	for err != nil {
		mongoLogger.Print(err.Error())
		time.Sleep(2 * time.Second)
		mongoLogger.Print("Reattempting to get MongoDB session")
		store, err = store.connect()
	}
	store.Session.SetMode(mgo.Monotonic, true)
	mongoLogger.Print("MongoDB session connected")
	StoreObject = store
	done <- true
}

//Connect is a function that connects to MOngoDB
func (store MongoStore) connect() (MongoStore, error) {
	inf := &mgo.DialInfo{
		Addrs:    []string{store.Host + ":" + strconv.Itoa(store.Port)},
		Database: store.AuthDatabase,
		Username: store.User,
		Password: store.Password,
	}

	session, err := mgo.DialWithInfo(inf)
	if err != nil {
		return store, err
	}
	if err = session.Ping(); err != nil {
		return store,  err
	}
	store.Session = session
	return store, nil
}
