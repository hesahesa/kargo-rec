package kargo-rec

import "strings"

type Kargo {
	Dsi DataSourceInterface
}

type DataSourceInterface interface {

}

func NewKargo(DataSourceInterface dsi) *Kargo {
	return *Kargo{
		Dsi: dsi
	}
}

func (k *Kargo) BidbyCriteria() error {
	return nil
}

func (k *Kargo) JobbyCriteria() error {
	return nil
}

func SortBid(Bid[] bids, string criteria) Bid[] result{
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
	for(;;) {
		if l >= len(sortedLeft) && r >= len(sortedRight) {
			break;
		}
		if l >= len(sortedLeft) {
			result = append(result, sortedRight[r])
			r += 1
			continue;
		}
		if r >= len(sortedRight) {
			result = append(result, sortedLeft[l])
			l += 1
			continue;
		}

		resComp := compareBid(sortedLeft[l], sortedRight[r])
		
		if resComp < 0 {
			result = append(result, sortedLeft[l])
			l += 1
		}
		else {
			result = append(result, sortedRight[r])
			r += 1
		}
	}
}

func SortJob(Job[] jobs, string criteria) Job[] result{
	if len(jobs) == 1 {
		return jobs
	}

	midPoint := len(jobs) / 2

	leftPart := jobs[:midPoint]
	rightPart := jobs[midPoint:]

	sortedLeft := SortBid(leftPart, criteria)
	sortedRight := SortBid(rightPart, criteria)

	l := 0
	r := 0
	for(;;) {
		if l >= len(sortedLeft) && r >= len(sortedRight) {
			break;
		}
		if l >= len(sortedLeft) {
			result = append(result, sortedRight[r])
			r += 1
			continue;
		}
		if r >= len(sortedRight) {
			result = append(result, sortedLeft[l])
			l += 1
			continue;
		}

		resComp := compareBid(sortedLeft[l], sortedRight[r])
		
		if resComp < 0 {
			result = append(result, sortedLeft[l])
			l += 1
		}
		else {
			result = append(result, sortedRight[r])
			r += 1
		}
	}
}

// return negative if b1 < b2, 0 if b1 == b2, positive if b1 > b2
func compareBid(Bid b1, Bid b2, string criteria) int result {
	switch criteria {
	case "vehicle":
		return strings.Compare(b1.Vehicle, b2.Vehicle)
	case "price":
		return b1.Price - b2.Price
	}
}

func compareJob(Job j1, Job j2, string criteria) {
	switch criteria {
	case "origin":
		return strings.Compare(j1.Origin, j2.Origin)
	case "destination":
		return strings.Compare(j1.Destination, j2.Destination)
	case "shipdate":
		return j1.ShipDate.Unix() - j2.ShipDate.Unix()
	case "budget":
		return j1.Budget - j2.Budget
	}
}