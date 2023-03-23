package log

import log "github.com/sirupsen/logrus"

const key = "Instance"

func Event(instance, msg string) {
	if instance != "" {
		log.WithField(key, instance).Info(msg)
	} else {
		log.Info(msg)
	}

}

func Fatal(instance, msg string, err error) {
	log.WithField(key, instance).Fatal(msg, ": ", err.Error())
}

func Error(instance, msg string, err error) {
	log.WithField(key, instance).Errorln(msg, ": ", err.Error())
}
