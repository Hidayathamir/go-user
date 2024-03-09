package integration_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/Eun/go-hit"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	err := ping()
	if err != nil {
		logrus.Fatalf("ping: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func ping() error {
	var err error
	for i := 0; i < 10; i++ {
		err = Do(Get("http://localhost:8080/ping"), Expect().Status().Equal(http.StatusOK))
		if err == nil {
			logrus.Infof("ping success 🟢")
			return nil
		}

		logrus.
			WithField("attempt count", i+1).
			Warnf("err ping: %v", err)

		time.Sleep(time.Second)
	}

	return err
}
