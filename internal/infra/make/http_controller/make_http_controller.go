package make_http_controller

import (
	card_services "github.com/ggarber42/payme/internal/domain/services/card"
	"github.com/ggarber42/payme/internal/infra/input/http_controller"
	postegres_repo "github.com/ggarber42/payme/internal/infra/output/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeHttpController (pgpool *pgxpool.Pool) *http_controller.HttpController {
	repo := postegres_repo.NewRepo(pgpool)
	cs := card_services.NewCardService(repo)

	hc := http_controller.NewHttpController(cs)
	return hc
}
