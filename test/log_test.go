package test

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusion/log"
	"testing"
)

func TestLog(t *testing.T) {
	illusion.Start()
	log.Info("asdfasdf:{}, ff:{}", "ffff", 2)
}
