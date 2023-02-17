package kitchen_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/meauxses/RecipeApp/kitchen"
)

func TestNewStorage(t *testing.T) {
	assert := assert.New(t)
	ing := kitchen.NewIngredient("chicken", "fridge", 1, "lb", 2022, 1, 1, 2022, 12, 11)
	fridge := kitchen.NewStorage()

	assert.Equal(0, fridge.Size(), "New storage is empty")
	fridge.AddToStorage(ing)
	assert.Equal(ing, (*fridge.GetIngredients())["chicken"][0], "should be the chicken object" )
	assert.Equal(1, fridge.Size(), "should be one")
	var foundIng *kitchen.Ingredient
	for _,value := range *fridge.GetIngredients() {
		foundIng = value[0]
	}
	assert.Equal("chicken", foundIng.Name(), "should be chicken")
	ing2 := kitchen.NewIngredient("beef", "fridge",1, "lb", 2022, 1, 1, 2022, 12, 11)
	ing3 := kitchen.NewIngredient("Cheddar cheese","fridge", 1, "lb", 2022, 1, 1, 2022, 12, 11)
	fridge.AddToStorage(ing2)
	fridge.AddToStorage(ing3)
	sortedNames := fridge.GetSortedIngredientNames()
	assert.Equal("beef", sortedNames[0], "should be beef")
	assert.Equal("cheddar cheese", sortedNames[1], "should be cheddar")
	assert.Equal("chicken", sortedNames[2], "should be chicken")

  }