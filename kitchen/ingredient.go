package kitchen

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	name string `bson:"name,omitempty"`
	storage string `bson:"storage,omitempty"`
	quantity int `bson:"quantity,omitempty"`
	baseUnit string `bson:"baseUnit,omitempty"`
	purchaseDate time.Time `bson:"purchaseDate,omitempty"`
	expirationDate time.Time `bson:"expirationDate,omitempty"`
}

func NewIngredient(name string, storage string, quantity int, baseUnit string, pyear int, pmonth time.Month, pday int, eyear int, emonth time.Month, eday int ) *Ingredient {
	ing := Ingredient{}
	ing.SetName(strings.ToLower(name))
	ing.SetStorage(strings.ToLower(storage))
	ing.SetQuantity(quantity)
	ing.SetBaseUnit(baseUnit)
	ing.SetPurchaseDate(pyear, pmonth, pday)
	ing.SetExpirationDate(eyear, emonth, eday)
	return &ing
}
func (ing *Ingredient) Name() string {
	return ing.name
}
func (ing *Ingredient) SetName(name string) {
	ing.name = name
}

func (ing *Ingredient) Storage() string {
	return ing.storage
}
func (ing *Ingredient) SetStorage(storage string) {
	ing.storage = storage
}
func (ing *Ingredient) Quantity()int {
	return ing.quantity
}
func (ing *Ingredient) SetQuantity(quantity int) {
	ing.quantity = quantity
}
func (ing *Ingredient) BaseUnit() string{
	return ing.baseUnit
}
func (ing *Ingredient) SetBaseUnit(baseUnit string) {
	ing.baseUnit = baseUnit
}
func (ing *Ingredient) PurchaseDate() time.Time {
	return ing.purchaseDate
}
func (ing *Ingredient) SetPurchaseDate(year int, month time.Month, day int) {
	ing.purchaseDate = time.Date(year, month, day, 0, 0, 0, 0,  time.Local)
}
func (ing *Ingredient) ExpirationDate() time.Time {
	return ing.expirationDate
}
func (ing *Ingredient) SetExpirationDate(year int, month time.Month, day int) {
	ing.expirationDate = time.Date(year, month, day, 0, 0, 0, 0,  time.Local)
}

func DaysUntilExpiration(ing *Ingredient) int {
	t := time.Now()
	year, month, day := t.Date()
	t = time.Date(year, month, day, 0, 0, 0, 0,  time.Local)
	duration := ing.expirationDate.Sub(t)
	duration = duration.Round(time.Hour)
	days := int(duration.Hours() / 24)
	return days
}


