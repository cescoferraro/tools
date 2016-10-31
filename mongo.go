package tools

import (
	"time"

	"github.com/fatih/color"
	"gopkg.in/mgo.v2"
	"strconv"
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

var mongoLogger = Logger{Title:"MONGO", Color:color.FgGreen}



//Cloner clones the local MONGO object
func Cloner() *MongoStore {
	return &MongoStore{
		Database: StoreObject.Database,
		Session: StoreObject.Session.Copy(),
	}
}

//InitMongo set MongoDB connection
func (store MongoStore) Init(done chan bool) {
	mongoLogger.Print("Attempting to get MongoDB session")
	err := store.connect()
	for err != nil {
		mongoLogger.Print(err.Error())
		time.Sleep(2 * time.Second)
		mongoLogger.Print("Reattempting to get MongoDB session")
		err = store.connect()
	}
	store.Session.SetMode(mgo.Monotonic, true)
	mongoLogger.Print("MongoDB session connected")
	StoreObject = store
	done <- true
}

//Connect is a function that connects to MOngoDB
func (store MongoStore) connect() (error) {

	inf := &mgo.DialInfo{
		Addrs:    []string{store.Host + ":" + strconv.Itoa(store.Port)},
		Database: store.AuthDatabase,
		Username: store.User,
		Password: store.Password,
	}

	var err error
	store.Session, err = mgo.DialWithInfo(inf)
	if err != nil {
		return err
	}
	if err = store.Session.Ping(); err != nil {
		return err
	}
	return nil
}
