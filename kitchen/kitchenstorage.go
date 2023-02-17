package kitchen

import (
	"sort"
	"strings"
)

type Storage struct {
	ingredientMap map[string][]*Ingredient
	size int
} 


func NewStorage() *Storage {
	ingredients := map[string][]*Ingredient{}
	store := Storage{ingredientMap: ingredients, size: 0}
	return &store
}

func (store *Storage) GetIngredients() map[string][]*Ingredient {
	return store.ingredientMap
}

func (store *Storage) GetSortedIngredientNames() []string {
	sortList := []string{}

	for key := range store.ingredientMap  {
		sortList = append(sortList, strings.ToLower(key))
	}
	sort.Strings(sortList)
	return sortList
}

func (store *Storage) AddToStorage(ingredient *Ingredient) {
	store.ingredientMap[ingredient.Name()] = append(store.ingredientMap[ingredient.Name()], ingredient)
	store.size += 1
}
//CHANGE TO MOVE FOOD TO WASTEBIN INSTEAD
func (store *Storage) RemoveFromStorage(ing *Ingredient) {
    
    arr := store.ingredientMap[ing.name]
	arrLen := len(arr)
    for i,ingredient := range arr{
        if ingredient == ing {
            arr = append( arr[:i], arr[i+1:arrLen]...)

        }
        if len(arr) == 0 {
            delete(kitch.GetSelection().ingredientMap, ing.name)
        }
    }
}

func (store *Storage) Size() int {
	return store.size
}




