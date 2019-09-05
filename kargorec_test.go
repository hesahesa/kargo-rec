package kargorec_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/hesahesa/kargorec/model"
	"github.com/hesahesa/kargorec"
)

func TestDummy(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestSortBidByPrice(t *testing.T) {
	bid1 := model.Bid {
		BidID: 1,
		Vehicle: "fuso",
		Price: 3000000,
		TransporterID: 1,
		JobID: 1,
	}

	bid2 := model.Bid {
		BidID: 2,
		Vehicle: "truk",
		Price: 2000000,
		TransporterID: 2,
		JobID: 2,
	}

	bids := []model.Bid{ bid1, bid2 }

	resultBid := kargorec.SortBid(bids, "price")

	assert.Equal(t, 2, len(resultBid))
	assert.Equal(t, 2000000, resultBid[0].Price)
	assert.Equal(t, 3000000, resultBid[1].Price)
}

func TestSortBidByVehicle(t *testing.T) {
	bid1 := model.Bid {
		BidID: 1,
		Vehicle: "fuso",
		Price: 3000000,
		TransporterID: 1,
		JobID: 1,
	}

	bid2 := model.Bid {
		BidID: 2,
		Vehicle: "truk",
		Price: 2000000,
		TransporterID: 2,
		JobID: 2,
	}

	bids := []model.Bid{ bid1, bid2 }

	resultBid := kargorec.SortBid(bids, "vehicle")

	assert.Equal(t, 2, len(resultBid))
	assert.Equal(t, "fuso", resultBid[0].Vehicle)
	assert.Equal(t, "truk", resultBid[1].Vehicle)
}

func TestSortJobByOrigin(t *testing.T) {
	job1 := model.Job {
		JobID: 1,
		Origin: "SMG",
		Destination: "JKT",
		ShipDate: time.Now(),
		Budget: 1000000,
		ShipperID: 1,
	}

	job2 := model.Job {
		JobID: 2,
		Origin: "CGK",
		Destination: "AMS",
		ShipDate: time.Now(),
		Budget: 3000000,
		ShipperID: 2,
	}

	jobs := []model.Job{ job1, job2 }

	resultJob := kargorec.SortJob(jobs, "origin")

	assert.Equal(t, 2, len(resultJob))
	assert.Equal(t, "CGK", resultJob[0].Origin)
	assert.Equal(t, "SMG", resultJob[1].Origin)
}

func TestSortJobByDestination(t *testing.T) {
	job1 := model.Job {
		JobID: 1,
		Origin: "SMG",
		Destination: "JKT",
		ShipDate: time.Now(),
		Budget: 1000000,
		ShipperID: 1,
	}

	job2 := model.Job {
		JobID: 2,
		Origin: "CGK",
		Destination: "AMS",
		ShipDate: time.Now(),
		Budget: 3000000,
		ShipperID: 2,
	}

	jobs := []model.Job{ job1, job2 }

	resultJob := kargorec.SortJob(jobs, "destination")

	assert.Equal(t, 2, len(resultJob))
	assert.Equal(t, "AMS", resultJob[0].Destination)
	assert.Equal(t, "JKT", resultJob[1].Destination)
}