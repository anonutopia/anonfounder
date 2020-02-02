package main

import (
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

type WavesMonitor struct {
	StartedTime int64
}

func (wm *WavesMonitor) start() {
	wm.StartedTime = time.Now().Unix() * 1000
	for {
		// todo - make sure that everything is ok with 10 here
		pages, err := wnc.TransactionsAddressLimit(conf.NodeAddress, 100)
		if err != nil {
			log.Println(err)
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				wm.checkTransaction(&t)
			}
		}
		// p, _ := pc.DoRequest()
		// log.Println(p)

		// ab, _ := wnc.AddressesBalance("3PLJQASFXtiohqbirYwSswjjbYGLfaGDEQy")
		// log.Println(ab)

		time.Sleep(time.Second)
	}
}

func (wm *WavesMonitor) checkTransaction(t *gowaves.TransactionsAddressLimitResponse) {
	tr := Transaction{TxID: t.ID}
	db.FirstOrCreate(&tr, &tr)
	if tr.Processed != true {
		wm.processTransaction(&tr, t)
	}
}

func (wm *WavesMonitor) processTransaction(tr *Transaction, t *gowaves.TransactionsAddressLimitResponse) {
	if t.Type == 4 && t.Timestamp >= wm.StartedTime && t.Sender != conf.NodeAddress && t.Recipient == conf.NodeAddress && len(t.AssetID) == 0 {
		log.Println(tr)
		log.Println(t)
	}

	tr.Processed = true
	db.Save(tr)
}

func initMonitor() {
	wm := &WavesMonitor{}
	wm.start()
}
