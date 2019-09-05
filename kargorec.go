package kargorec

import (
	"strings"

	"github.com/hesahesa/kargorec/model"
)

type Kargo struct {
	Dsi DataSourceInterface
}

type DataSourceInterface interface {
	GetAllBidByJobID(int) []model.Bid
	GetAllJob() []model.Job
}

func NewKargo(dsi DataSourceInterface) *Kargo {
	return &Kargo{
		Dsi: dsi,
	}
}

func (k *Kargo) BidbyCriteria(jobID int, criteria string) (bids []model.Bid, err error) {
	allBidByID := k.Dsi.GetAllBidByJobID(jobID)

	sortedBid := SortBid(allBidByID, criteria)
	return sortedBid, nil
}

func (k *Kargo) JobbyCriteria(criteria string) (jobs []model.Job, err error) {
	allJobs := k.Dsi.GetAllJob()

	sortedJobs := SortJob(allJobs, criteria)
	return sortedJobs, nil
}

func SortBid(bids []model.Bid, criteria string) (result []model.Bid) {
	if len(bids) == 1 {
		return bids
	}

	midPoint := len(bids) / 2

	leftPart := bids[:midPoint]
	rightPart := bids[midPoint:]

	sortedLeft := SortBid(leftPart, criteria)
	sortedRight := SortBid(rightPart, criteria)

	l := 0
	r := 0
	for {
		if l >= len(sortedLeft) && r >= len(sortedRight) {
			break
		}
		if l >= len(sortedLeft) {
			result = append(result, sortedRight[r])
			r += 1
			continue
		}
		if r >= len(sortedRight) {
			result = append(result, sortedLeft[l])
			l += 1
			continue
		}

		resComp := compareBid(sortedLeft[l], sortedRight[r], criteria)

		if resComp < 0 {
			result = append(result, sortedLeft[l])
			l += 1
		} else {
			result = append(result, sortedRight[r])
			r += 1
		}
	}
	return result
}

func SortJob(jobs []model.Job, criteria string) (result []model.Job) {
	if len(jobs) == 1 {
		return jobs
	}

	midPoint := len(jobs) / 2

	leftPart := jobs[:midPoint]
	rightPart := jobs[midPoint:]

	sortedLeft := SortJob(leftPart, criteria)
	sortedRight := SortJob(rightPart, criteria)

	l := 0
	r := 0
	for {
		if l >= len(sortedLeft) && r >= len(sortedRight) {
			break
		}
		if l >= len(sortedLeft) {
			result = append(result, sortedRight[r])
			r += 1
			continue
		}
		if r >= len(sortedRight) {
			result = append(result, sortedLeft[l])
			l += 1
			continue
		}

		resComp := compareJob(sortedLeft[l], sortedRight[r], criteria)

		if resComp < 0 {
			result = append(result, sortedLeft[l])
			l += 1
		} else {
			result = append(result, sortedRight[r])
			r += 1
		}
	}
	return result
}

// return negative if b1 < b2, 0 if b1 == b2, positive if b1 > b2
func compareBid(b1 model.Bid, b2 model.Bid, criteria string) (result int) {
	switch criteria {
	case "vehicle":
		return strings.Compare(b1.Vehicle, b2.Vehicle)
	case "price":
		return b1.Price - b2.Price
	}
	return 0
}

func compareJob(j1 model.Job, j2 model.Job, criteria string) (result int) {
	switch criteria {
	case "origin":
		return strings.Compare(j1.Origin, j2.Origin)
	case "destination":
		return strings.Compare(j1.Destination, j2.Destination)
	case "shipdate":
		return int(j1.ShipDate.Unix() - j2.ShipDate.Unix())
	case "budget":
		return int(j1.Budget - j2.Budget)
	}
	return 0
}
