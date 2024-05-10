package smtp

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestSMTP(t *testing.T) {
	pass := rand.Intn(8999) + 1000
	SendMail("mirodilpychekio@gmail.com", "your password: "+strconv.Itoa(pass))
}
