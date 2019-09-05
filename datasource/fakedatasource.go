package datasource

import (
	"time"

	"github.com/hesahesa/kargorec/model"
)
type FakeDataSource struct {}

func NewFakeDataSource() *FakeDataSource {
	return &FakeDataSource{}
}

// GetAllBidByJobID(int) []model.Bid
// GetAllJob() []model.Job

func (fd *FakeDataSource) GetAllBidByJobID(id int) []model.Bid {
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
		JobID: 1,
	}

	bids := []model.Bid{ bid1, bid2 }

	return bids
} 

func (fd *FakeDataSource) GetAllJob() []model.Job {
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

	return jobs
} 