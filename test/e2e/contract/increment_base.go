package main

import (
	"github.com/orbs-network/orbs-contract-sdk/go/sdk/v1"
)

var PUBLIC = sdk.Export(inc, value, getOwnAddress)
var SYSTEM = sdk.Export(_init)

var COUNTER_KEY = []byte("counter")

func Inc(value uint64) {}

func _init() {

}
