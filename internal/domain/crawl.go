package domain

import beegoContext "github.com/beego/beego/v2/server/web/context"

type CrawlResponse struct {
	MetaContent       MetaContent `json:"meta_content"`
	HTags             []HTags     `json:"h_tags"`
	SourceCodeHtmlUrl string      `json:"source_code_html_url"`
}

type CrawlRequest struct {
	WebUrl string `json:"web_url" validate:"required,URL"`
}

type MetaContent struct {
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
}

type HTags struct {
	Tags string   `json:"tags"`
	List []string `json:"list"`
}

type IndexableLink struct {
	No       int    `json:"no"`
	LinksURL string `json:"links_url"`
}

// CrawlUseCase UseCase Interface
type CrawlUseCase interface {
	CrawlWeb(beegoCtx *beegoContext.Context, request *CrawlRequest) (*CrawlResponse, error)
}
