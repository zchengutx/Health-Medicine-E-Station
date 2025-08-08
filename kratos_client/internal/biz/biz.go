package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUsecase, NewCityUsecase, NewDoctorsUsecase, NewDrugService, NewEstimateService, NewChatService, NewCartService, NewOrderUsecase, NewPaymentUsecase, NewCouponUsecase, NewPrescriptionUsecase)
