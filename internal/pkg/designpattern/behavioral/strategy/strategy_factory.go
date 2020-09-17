package strategy

import (
	"sync"
)

type PromotionFactory struct {
	strategyMap map[string]PromotionStrategy
}

var single *PromotionFactory
var once sync.Once

/**
单例工厂
*/
func GetInstance() *PromotionFactory {
	once.Do(func() {
		single = new(PromotionFactory)
		single.strategyMap = make(map[string]PromotionStrategy)
		single.strategyMap[PROMOTION_COUPON_STRATEGY] = new(CouponStrategy)
		single.strategyMap[PROMOTION_CASHBACK_STRATEGY] = new(CashbackStrategy)
	})

	return single
}

/**
获取指定策略
*/
func (factory *PromotionFactory) GetStrategy(name string) PromotionStrategy {
	return factory.strategyMap[name]
}
