package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {

	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())

	strategy := GetInstance().GetStrategy(PROMOTION_COUPON_STRATEGY)

	request := new(PromotionRequest)
	strategy.DoPromotion(*request)

}
