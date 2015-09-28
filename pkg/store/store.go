package store

import (
	"github.com/Sirupsen/logrus"
)

var drivers = make(map[string]DriverFunc)

func Register(name string, driver DriverFunc) {
	drivers[name] = driver
}

type DriverFunc func(driver, datasource string) (Store, error)

func New(driver, datasource string) (Store, error) {
	fn, ok := drivers[driver]

	if !ok {
		logrus.Fatalf("Store: Unknown driver %q", driver)
	}

	logrus.Infof("Store: Loading driver %s", driver)
	logrus.Infof("Store: Loading config %s", datasource)

	return fn(driver, datasource)
}

type Store interface {
}
