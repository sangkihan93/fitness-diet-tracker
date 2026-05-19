package nutrition

type DietaryRestriction struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type NutritionTarget struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	DailyCalories int     `json:"daily_calories"`
	DailyProteinG float64 `json:"daily_protein_g"`
	DailyCarbsG   float64 `json:"daily_carbs_g"`
	DailyFatG     float64 `json:"daily_fat_g"`
	WaterTargetOz *int    `json:"water_target_oz,omitempty"`
}
