package main

import (
	"log"
	"sort"

	"github.com/go-numb/go-ftx/rest/public/futures"
	"github.com/schigh/slice"
)

var (
	rc = RestClient
)

type TradableFutures struct {
	Name                              string
	Change1H, Change24H, VolumeUsd24H float64
}

type TradableFuturesSlice []TradableFutures

func (in *TradableFuturesSlice) Sort() (ret []TradableFutures) {
	inn := *in
	 sort.SliceStable(*in,
		func(i int, j int) bool { return inn[i].VolumeUsd24H < inn[j].VolumeUsd24H })
	ret = *in
	return
}

func CurrentPotentFutures(perpetual bool) (list []*TradableFutures) {
	getFutures, err := rc.Futures(&futures.RequestForFutures{})
	rangeFutures := *getFutures

	if err != nil {
		return
	}
	ret := getFutureList(perpetual, rangeFutures, make([]futures.FutureForList, 0))

	// work todo???
	for i := range ret {

	}

	log.Println(len(*getFutures))
	log.Println(len(list))
	return list
}


func getFutureList(perpetual bool, rangeFutures futures.ResponseForFutures, list []futures.FutureForList) []futures.FutureForList {
	if perpetual {
		for i := range rangeFutures {
			if rangeFutures[i].Perpetual {
				list = append(list, rangeFutures[i])
			}

		}
	} else {
		for v := range rangeFutures {
			list = append(list, rangeFutures[v])
		}
	}
	return list
}
