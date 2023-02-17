package kitchen

import (
	"context"
	"fmt"
	"log"
	"sort"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type Manager struct {
    fridge *Storage
    freezer *Storage
    pantry *Storage
    wastebin *Storage
    selection *Storage
}

var kitch *Manager

func GetKitchen() *Manager {
    if kitch == nil {
            kitch = &Manager{fridge: NewStorage(), freezer: NewStorage(), pantry: NewStorage(), wastebin: NewStorage(), selection: NewStorage()}
    } else {
        fmt.Println("Kitchen instance already created.")
    }
    return kitch
}
func (manager *Manager) GetFridge() *Storage {
    return manager.fridge
}
func (manager *Manager) GetFreezer() *Storage {
    return manager.freezer
}
func (manager *Manager) GetPantry() *Storage {
    return manager.pantry
}
func (manager *Manager) GetWastebin() *Storage {
    return manager.wastebin
}
func (manager *Manager) GetSelection() *Storage {
    return manager.selection
}
func (manager *Manager) GetStorage(storage string) *Storage{
    if storage == "fridge" {
        return manager.fridge
    }
    if storage == "freezer" {
        return manager.freezer
    }
    if storage == "pantry" {
        return manager.pantry
    }
    if storage == "selection" {
        return manager.selection
    }
    if storage == "wastebin" {
        return manager.wastebin
    }
    return nil
}

func LoadStorage(client *mongo.Client, container *Storage, storage string) {
    ingredientCol := client.Database("kitchen").Collection("ingredients")
    var results []Ingredient
    cursor, err := ingredientCol.Find(context.TODO(), bson.D{{Key: "storage", Value: storage}})
	if err != nil {
		panic(err)
	}
	if err := cursor.All(context.TODO(), results); err != nil {
		panic(err)
	}
	for _, result := range results {
        container.ingredientMap[result.name] = append(container.ingredientMap[result.name], &result)
	}
}
func SaveStorage(client *mongo.Client, container *Storage, storage string)      {
    ingredientCol := client.Database("kitchen").Collection("ingredients")
    for _,arr := range container.ingredientMap {
        for _,ingredient := range arr {
            result,err := ingredientCol.InsertOne(context.TODO(), *ingredient)
            if err != nil {
                panic(err)
            } else {
                log.Printf("Inserted ingredient with id: %v\n", result.InsertedID)
            }
        }
    }
}

func SelectIngredient(ing *Ingredient) {
    kitch.GetSelection().AddToStorage(ing)
}
func DeselectIngredient(ing *Ingredient) {
    kitch.GetSelection().RemoveFromStorage(ing)
}
func MoveSelectionToWastebin() {
    for _,arr := range kitch.selection.ingredientMap {
        for _,ingredient := range arr {
            kitch.wastebin.ingredientMap[ingredient.name] = append(kitch.wastebin.ingredientMap[ingredient.name], ingredient)
        }
    }
}
func ClearWastebin() {
    for _,arr := range kitch.wastebin.ingredientMap {
        for _,ing := range arr {
            storageToCheck := ing.storage
            kitch.GetStorage(storageToCheck).RemoveFromStorage(ing)
            kitch.selection.RemoveFromStorage(ing)
            ing = nil 
        }
    }
    kitch.wastebin.ingredientMap = nil
    kitch.wastebin.size = 0
}

func ViewStorage(storageSlice []*Ingredient)   {
    //use string builder
}
func StorageToSlice(storage *Storage) []*Ingredient {
    sliceForm := []*Ingredient{}
    for _,value := range storage.GetIngredients() {
        sliceForm = append(sliceForm, value...)
    }
    return sliceForm
}

func SortStorageName(sliceForm []*Ingredient)  []*Ingredient {
    sort.SliceStable(sliceForm, func(i, j int) bool { return sliceForm[i].name < sliceForm[j].name })
    return sliceForm
}
func SortStorageExpiration(sliceForm []*Ingredient) []*Ingredient {
    sort.SliceStable(sliceForm, func(i, j int) bool { return sliceForm[i].expirationDate.Before(sliceForm[j].expirationDate)})
    return sliceForm
}
func SortStoragePurchase(sliceForm []*Ingredient)  []*Ingredient {
    sort.SliceStable(sliceForm, func(i, j int) bool { return sliceForm[i].purchaseDate.Before(sliceForm[j].purchaseDate)})
    return sliceForm
}


func ViewRecipes()   {}
func FindRecipeAccordingToSelection() {}
func LoadRecipes(client *mongo.Client) {

}
