package kitchen

//recipe has a list of ingredients used, list of utensils used, recipe steps, each recipe step has a time associated with which you can call a timer for that time
//each step also should include the ingredients listed in order used for that step?
type Recipe struct {
    title string
    ingredientMap map[string]*Ingredient
    utensilMap map[string]int
    stepList []*RecipeStep
}

type RecipeStep struct {
    text string
    minMinutes int
    maxMinutes int
}

func NewRecipeStep(text string, minMinutes int, maxMinutes int) *RecipeStep {
    recipeStep := RecipeStep{}
    recipeStep.SetText(text)
    recipeStep.SetMinMinutes(minMinutes)
    recipeStep.SetMaxMinutes(maxMinutes)
    return &recipeStep
}

func(recipeStep *RecipeStep) Text() string {
    return recipeStep.text
}
func(recipeStep *RecipeStep) SetText(text string) {
    recipeStep.text = text
}
func(recipeStep *RecipeStep) MinMinutes() int {
    return recipeStep.minMinutes
}
func(recipeStep *RecipeStep) SetMinMinutes(min int) {
    recipeStep.minMinutes = min
}
func(recipeStep *RecipeStep) MaxMinutes() int {
    return recipeStep.maxMinutes
}
func(recipeStep *RecipeStep) SetMaxMinutes(max int) {
    recipeStep.maxMinutes = max
}



func NewRecipe(title string, ingredientMap map[string]*Ingredient, utensilMap map[string]int, stepList []*RecipeStep) *Recipe {
    recipe := &Recipe{}
    recipe.SetTitle(title)
    recipe.SetIngredientMap(ingredientMap)
    recipe.SetUtensilMap(utensilMap)
    recipe.SetStepList(stepList)
    return recipe
}

func (recipe *Recipe) Title() string {
    return recipe.title
}

func (recipe *Recipe) SetTitle(title string) {
    recipe.title = title
}

func (recipe *Recipe) IngredientMap() map[string]*Ingredient {
    return recipe.ingredientMap
}

func (recipe *Recipe) SetIngredientMap( ingredientMap map[string]*Ingredient) {
    recipe.ingredientMap = ingredientMap
}
func (recipe *Recipe) UtensilMap() map[string]int{
    return recipe.utensilMap
}
func (recipe *Recipe) SetUtensilMap(utensilMap map[string]int){
    recipe.utensilMap = utensilMap
}
func (recipe *Recipe) StepList() []*RecipeStep {
    return recipe.stepList
}
func (recipe *Recipe) SetStepList(stepList []*RecipeStep) {
    recipe.stepList = stepList
}

