package kitchen_test

import (
	"testing"

	"github.com/meauxses/RecipeApp/kitchen"
	"github.com/stretchr/testify/assert"
)

func TestNewIngredient(t *testing.T) {
    assert := assert.New(t)
    ing := kitchen.NewIngredient("chicken", "fridge", 1, "lb", 2023, 1, 23, 2023, 2, 12)
    assert.Equal("chicken", ing.Name(), "name should be chicken")
    assert.Equal(1, ing.Quantity(), "quantity should be 1")
    assert.Equal(1, kitchen.DaysUntilExpiration(ing), "time until expiration should be 1")
  }
