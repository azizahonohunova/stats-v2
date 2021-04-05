package stats

import (

	"github.com/azizahonohunova/bank/v2/pkg/types"

)

func Avg(payments []types.Payment) types.Money {
	var sum types.Money = 0
	var cnt int = 0
	for _, x := range payments {
		if x.Status == "FAIL" {
			continue
		}
		cnt++
		sum += x.Amount
	}
	return (sum / types.Money(cnt))
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money = 0
	for _, x := range payments {
		if x.Status == "FAIL" {
			continue
		}
		if x.Category == category {
			sum += x.Amount
		}
	}
	return sum
}
