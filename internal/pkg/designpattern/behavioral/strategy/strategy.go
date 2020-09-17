package strategy

import (
	"fmt"
)

type PromotionRequest struct {
	promotionId string
}

type PromotionStrategy interface {
	DoPromotion(request PromotionRequest)
}

/**

 */
type CouponStrategy struct{}

func (strategy *CouponStrategy) DoPromotion(request PromotionRequest) {
	fmt.Println("CouponStrategy", request.promotionId)
}

type CashbackStrategy struct{}

func (CashbackStrategy) DoPromotion(request PromotionRequest) {
	fmt.Println("CashbackStrategy", request.promotionId)
}
