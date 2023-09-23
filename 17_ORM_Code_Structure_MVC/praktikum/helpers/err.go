package helpers

import "github.com/sirupsen/logrus"

func LogIfError(message string, err error) error {
  if err != nil {
    logrus.Error(message, err.Error())
  }
  return nil
}