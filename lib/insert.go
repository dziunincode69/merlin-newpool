package lib

import (
	"log"
	"merlin-newpool/models"

	"github.com/kamva/mgm/v3"
)

func InsertTknLog(tknAddress, name, symbol, Owner, poolAddress string, Fee, Poolid int64) {
	tkn := models.NewModelLogTkn(tknAddress, name, symbol, Owner, poolAddress, Fee, Poolid)
	err := mgm.Coll(tkn).Create(tkn)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertNewTkn(tknAddress, name, symbol, Owner string) {
	tkn := models.NewModelNewToken(tknAddress, name, symbol, Owner)
	err := mgm.Coll(tkn).Create(tkn)
	if err != nil {
		log.Fatal(err)
	}
}
