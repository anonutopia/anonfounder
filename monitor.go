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
		// todo - make sure that everything is ok with 100 here
		pages, err := wnc.TransactionsAddressLimit(conf.NodeAddress, 100)
		if err != nil {
			log.Println(err)
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				wm.checkTransaction(&t)
			}
		}

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
	if t.Type == 4 &&
		t.Timestamp >= wm.StartedTime &&
		t.Sender != conf.NodeAddress &&
		t.Recipient == conf.NodeAddress &&
		len(t.AssetID) == 0 {
		amount := token.issueAmount(t.Amount, t.AssetID)

		atr := &gowaves.AssetsTransferRequest{
			Amount:    int(float64(t.Amount) * conf.FounderFactor),
			AssetID:   "",
			Fee:       100000,
			Recipient: conf.FounderAddress,
			Sender:    conf.NodeAddress,
		}

		_, err := wnc.AssetsTransfer(atr)
		if err != nil {
			log.Printf("[WavesMonitor.processTransation] error assets transfer: %s", err)
		} else {
			log.Printf("Sent token: %s => %d", t.Sender, amount)
		}

		atr = &gowaves.AssetsTransferRequest{
			Amount:    amount,
			AssetID:   conf.TokenID,
			Fee:       100000,
			Recipient: t.Sender,
			Sender:    conf.NodeAddress,
		}

		_, err = wnc.AssetsTransfer(atr)
		if err != nil {
			log.Printf("[WavesMonitor.processTransation] error assets transfer: %s", err)
		} else {
			log.Printf("Sent token: %s => %d", t.Sender, amount)
		}
	}

	tr.Processed = true
	db.Save(tr)
}

func initMonitor() {
	wm := &WavesMonitor{}
	wm.start()
}
