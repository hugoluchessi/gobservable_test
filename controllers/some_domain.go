package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/hugoluchessi/gobservable_test/config"
)

func HandleSomeDomainGetSomething(ms *config.MonitorServices, rw http.ResponseWriter, req *http.Request) {
	sleepTime := rand.Uint32() % 700
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	if sleepTime > 500 {
		ms.IncrCounter([]string{"req", "slow"}, 1)
	}

	fmt.Fprint(rw, "DONE!")
}
