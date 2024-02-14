package lib

import (
	"log"
	"merlin-newpool/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckTKNInserted(contractAddress common.Address) bool {
	var results []*models.LogToken
	err := mgm.Coll(&models.LogToken{}).SimpleFind(&results, bson.M{"tokenAddress": contractAddress.Hex()})
	if err != nil {
		log.Fatal(err) // Atau handle error sesuai kebutuhan
	}

	// Cek apakah ada data yang ditemukan
	if len(results) > 0 {
		return false
	}
	return true
}
