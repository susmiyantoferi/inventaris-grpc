package log

import "log"

type ConsoleLogging struct{}

func (c *ConsoleLogging) MsgInfo(msg string){
	log.Printf("[INFO] %s", msg)
}

func (c *ConsoleLogging) ErrInfo(err error, str string){
	log.Printf("[ERROR] %s | %v", err, str)
}