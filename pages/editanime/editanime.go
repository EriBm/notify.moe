package editanime

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Main anime edit page.
func Main(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this anime", nil)
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.EditAnimeTabs(anime) + editform.Render(anime, "Edit anime", user))
}

// Characters anime characters edit page.
func Characters(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit", nil)
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	animeCharacters, err := arn.GetAnimeCharacters(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime characters not found", err)
	}

	return ctx.HTML(components.EditAnimeTabs(anime) + editform.Render(animeCharacters, "Edit anime characters", user))
}

// Relations anime relations edit page.
func Relations(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit", nil)
	}

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	animeRelations, err := arn.GetAnimeRelations(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime relations not found", err)
	}

	return ctx.HTML(components.EditAnimeTabs(anime) + editform.Render(animeRelations, "Edit anime relations", user))
}
