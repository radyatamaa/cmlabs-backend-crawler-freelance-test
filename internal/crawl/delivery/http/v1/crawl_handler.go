package v1

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/pkg/errors"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/domain"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/response"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/validator"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/zaplogger"
	"gorm.io/gorm"
	"net/http"
)

type CrawlHandler struct {
	ZapLogger zaplogger.Logger
	internal.BaseController
	response.ApiResponse
	Usecase domain.CrawlUseCase
}

func NewCrawlHandler(useCase domain.CrawlUseCase, zapLogger zaplogger.Logger) {
	pHandler := &CrawlHandler{
		ZapLogger: zapLogger,
		Usecase:   useCase,
	}
	beego.Router("/api/v1/crawl/web", pHandler, "post:CrawlWeb")
}

func (h *CrawlHandler) Prepare() {
	// check user access when needed
	h.SetLangVersion()
}

// CrawlWeb
// @Title CrawlWeb
// @Tags Crawl
// @Summary CrawlWeb
// @Produce json
// @Param Accept-Language header string false "lang"
// @Success 200 {object} swagger.BaseResponse{errors=[]object,data=domain.CrawlResponse}
// @Failure 400 {object} swagger.BadRequestErrorValidationResponse{errors=[]swagger.ValidationErrors,data=object}
// @Failure 408 {object} swagger.RequestTimeoutResponse{errors=[]object,data=object}
// @Failure 500 {object} swagger.InternalServerErrorResponse{errors=[]object,data=object}
// @Param body body domain.CrawlRequest true "request payload"
// @Router /v1/crawl/web [post]
func (h *CrawlHandler) CrawlWeb() {
	var request domain.CrawlRequest

	if err := h.BindJSON(&request); err != nil {
		h.Ctx.Input.SetData("stackTrace", h.ZapLogger.SetMessageLog(err))
		h.ResponseError(h.Ctx, http.StatusBadRequest, response.ApiValidationCodeError, response.ErrorCodeText(response.ApiValidationCodeError, h.Locale.Lang), err)
		return
	}
	if err := validator.Validate.ValidateStruct(&request); err != nil {
		h.Ctx.Input.SetData("stackTrace", h.ZapLogger.SetMessageLog(err))
		h.ResponseError(h.Ctx, http.StatusBadRequest, response.ApiValidationCodeError, response.ErrorCodeText(response.ApiValidationCodeError, h.Locale.Lang), err)
		return
	}

	result, err := h.Usecase.CrawlWeb(h.Ctx, &request)
	if err != nil {
		if errors.Is(err, response.ErrInvalidEmailPassword) {
			h.ResponseError(h.Ctx, http.StatusBadRequest, response.InvalidEmailPassword, response.ErrorCodeText(response.InvalidEmailPassword, h.Locale.Lang), err)
			return
		}
		if errors.Is(err, context.DeadlineExceeded) {
			h.ResponseError(h.Ctx, http.StatusRequestTimeout, response.RequestTimeoutCodeError, response.ErrorCodeText(response.RequestTimeoutCodeError, h.Locale.Lang), err)
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.ResponseError(h.Ctx, http.StatusBadRequest, response.DataNotFoundCodeError, response.ErrorCodeText(response.DataNotFoundCodeError, h.Locale.Lang), err)
			return
		}
		h.ResponseError(h.Ctx, http.StatusInternalServerError, response.ServerErrorCode, response.ErrorCodeText(response.ServerErrorCode, h.Locale.Lang), err)
		return
	}
	h.Ok(h.Ctx, h.Tr("message.success"), result)
	return
}
