package log

import "log"

func Event(instance, msg string) {
	log.Printf("%s: %s.\n", instance, msg)
}

func Fatal(instance, msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s. %s", instance, msg, err.Error())
	} else {
		log.Fatalf("%s: %s.", instance, msg)
	}
}
