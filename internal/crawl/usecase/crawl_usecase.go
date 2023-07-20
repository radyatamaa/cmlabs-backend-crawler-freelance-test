package usecase

import (
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/domain"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/zaplogger"
	"time"
	"github.com/gocolly/colly"
)

type crawlUseCase struct {
	zapLogger                  zaplogger.Logger
	contextTimeout             time.Duration
}


func NewCrawlUseCase(timeout time.Duration,
	zapLogger zaplogger.Logger) domain.CrawlUseCase {
	return &crawlUseCase{
		contextTimeout:             timeout,
		zapLogger:                  zapLogger,
	}
}

func (c crawlUseCase) GetCrawl(beegoCtx *beegoContext.Context, urlWebsite string) (*domain.CrawlResponse, error) {
	panic("implement me")
}
