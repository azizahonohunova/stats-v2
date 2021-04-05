package stats
import (
	"fmt"

	"github.com/azizahonohunova/bank/pkg/types"
)


func ExampleAvg() {
	payment := []types.Payment{
		{
			ID:       1000,
			Amount:   100,
			Category: "gold",
		},
		{
			ID:       2000,
			Amount:   500,
			Category: "gold",
		},
	}
	max := Avg(payment)
	fmt.Println(max)
	//Output: 300

}
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "Auto",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   1000,
			Category: "Cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   30,
			Category: "Restaurang",
			Status:   types.StatusFail,
		},
	}
	result := types.Category("Cafe")
	totalInCategory := TotalInCategory(payments, result)
	fmt.Println(totalInCategory)
	//Output: 1000
}
