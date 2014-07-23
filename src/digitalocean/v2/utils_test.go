package v2

import (
	"log"
	"testing"
)

func TestUtils(t *testing.T) {
	// if Add(5, 6) != 11 {
	// 	log.Println("Addition is not working well")
	// 	t.Fail()
	// }
	var config *Configuration
	var err error
	if config, err = GetConfig(); err != nil {
		log.Println(err)
		t.Fail()
	}
	log.Println(config)

}
