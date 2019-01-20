package logger

import (
	"log"
)

func Info(msg string) {
	log.Println("[INFO]",msg)
}

func Warn(msg string) {
	log.Println("[WARN]", msg)
}

func Err(msg string)  {
	log.Fatalln("[ERRO]", msg)
}

func ErrNotNil(err error, msg string)  {
	if err != nil {
		log.Fatalln("[ERRO]", msg, " | ", err)
	}
}