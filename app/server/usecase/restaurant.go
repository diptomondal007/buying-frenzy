package usecase

type RestaurantUseCase interface {
	ListRestaurantsByFilter() error
}

type restaurantUseCase struct{}

func (r restaurantUseCase) ListRestaurantsByFilter() error {
	//TODO implement me
	return nil
}

func NewRestaurantUseCase() RestaurantUseCase {
	return restaurantUseCase{}
}
