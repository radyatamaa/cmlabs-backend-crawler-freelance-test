package usecase

import (
	"context"
	"fmt"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/gocolly/colly"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/domain"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/helper"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/zaplogger"
	"os"
	"time"
)

type crawlUseCase struct {
	zapLogger      zaplogger.Logger
	contextTimeout time.Duration
	coll           *colly.Collector
}

func NewCrawlUseCase(coll *colly.Collector, timeout time.Duration,
	zapLogger zaplogger.Logger) domain.CrawlUseCase {
	return &crawlUseCase{
		contextTimeout: timeout,
		zapLogger:      zapLogger,
		coll:           coll,
	}
}

func (c crawlUseCase) generateFile(beegoCtx *beegoContext.Context, content []byte, fileName string) error {
	file, err := os.Create("external/storage/" + fileName + ".html") //create a new file
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", c.zapLogger.SetMessageLog(err))
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", c.zapLogger.SetMessageLog(err))
		return err
	}

	return err
}

func (c crawlUseCase) CrawlWeb(beegoCtx *beegoContext.Context, request *domain.CrawlRequest) (*domain.CrawlResponse, error) {
	_, cancel := context.WithTimeout(beegoCtx.Request.Context(), c.contextTimeout)
	defer cancel()

	result := new(domain.CrawlResponse)

	c.coll.OnRequest(func(r *colly.Request) {

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}

		fmt.Println(r.Method)
	})

	c.coll.OnHTML("title", func(e *colly.HTMLElement) {
		result.MetaContent.MetaTitle = e.Text
	})

	c.coll.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		result.MetaContent.MetaDescription = e.Attr("content")
	})

	result.HTags = make([]domain.HTags, 0)
	var hTags []string
	c.coll.OnHTML(`h1, h2, h3, h4, h5, h6`, func(e *colly.HTMLElement) {
		check, index := helper.ItemExistsIndex(hTags, e.Name)
		if check {
			result.HTags[index].List = append(result.HTags[index].List, e.Text)
		} else {
			result.HTags = append(result.HTags, domain.HTags{
				Tags: e.Name,
				List: []string{e.Text},
			})

			hTags = append(hTags, e.Name)
		}
	})

	fileName := helper.GetFileNameFromUrl(request.WebUrl)
	c.coll.OnResponse(func(r *colly.Response) {
		fmt.Println("-----------------------------")

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
		if r.StatusCode == 200 {
			err := c.generateFile(beegoCtx, r.Body, fileName)
			if err != nil {
				beegoCtx.Input.SetData("stackTrace", c.zapLogger.SetMessageLog(err))
			}
		}
	})

	result.SourceCodeHtmlUrl = fmt.Sprint("http://", beegoCtx.Request.Host, "/external/storage/"+fileName+".html")
	err := c.coll.Visit(request.WebUrl + "?unique=" + helper.RandomString(10))
	if err != nil {
		beegoCtx.Input.SetData("stackTrace", c.zapLogger.SetMessageLog(err))
		return nil, err
	}
	return result, nil
}
