package entities

type Subscriptions = []Subscription

type Subscription struct {
	Id     int `json:"id"`
	CityId int `json:"cityId"`
}

type UserSubscription struct {
	UserId         int `json:"userId"`
	SubscriptionId int `json:"subsciptionId"`
}
