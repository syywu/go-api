package mocks

import "github.com/syywu/go-api.git/pkg/models"

var Wines = []models.Wine{
	{
		Id:          1,
		Producer:    "Olivier Leflaive",
		Name:        "St Aubin 1er Cru Dents de Chien",
		Region:      "France",
		Rating:      92,
		Description: "Gorgeous nose of buttery lemon, orange zest, honeysuckle, vanilla and melon. The texture is crisp and refreshing, with a vanilla caramel density shot through with lime. Creaminess and oak underpin the fruit.",
	},
}
