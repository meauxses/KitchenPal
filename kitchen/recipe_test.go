package kitchen_test

import (
	"testing"

	"github.com/meauxses/RecipeApp/kitchen"
	"github.com/stretchr/testify/assert"
)

func TestNewRecipe(t *testing.T) {
    assert := assert.New(t)
    ing := kitchen.NewIngredient("chicken breast","fridge", 1, "lb", 2022, 1, 1, 2022, 12, 18)
    ing2 := kitchen.NewIngredient("beef","fridge", 1, "lb", 2022, 1, 1, 2022, 12, 18)
    ing3 := kitchen.NewIngredient("cheese","fridge", 1, "lb", 2022, 1, 1, 2022, 12, 18)
    ingredientMap := map[string]*kitchen.Ingredient{"chicken breast": ing, "beef":ing2, "cheese":ing3}
    utensilMap := map[string]int{"frying pan": 1, "dutch oven": 1}
    step1 := kitchen.NewRecipeStep("Fry the beef", 1, 5)
    step2 := kitchen.NewRecipeStep("Wash the cheese", 1, 5)
    stepList := []*kitchen.RecipeStep{}
    stepList = append(stepList, step1, step2)
    recipe1 := kitchen.NewRecipe("Beef and Cheese", ingredientMap, utensilMap, stepList)

    assert.Equal("Beef and Cheese", recipe1.Title(), "beef and cheese")
    
    
  }