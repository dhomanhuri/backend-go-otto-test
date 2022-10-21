package presentation

import (
	"otto/test/features/brands"
	"otto/test/features/brands/presentation/dto"
	"otto/test/helper"

	"github.com/gin-gonic/gin"
)

type BrandPres struct {
	Pres brands.Bussines
}

func BrandHandler(pres brands.Bussines) BrandPres {
	return BrandPres{
		Pres: pres,
	}
}

func (Bb BrandPres) BrandInsert(c *gin.Context) {
	brandReq := dto.BrandRequest{}
	err := c.Bind(&brandReq)
	if err != nil {
		c.JSON(helper.BadRequest("error bidn"))
		c.Abort()
		return
	}
	core := brands.Core{
		Name: brandReq.Name,
	}

	result := Bb.Pres.BussInsert(core)
	if result != nil {
		c.JSON(helper.BadRequest(result.Error()))
		c.Abort()
		return
	}
	c.JSON(helper.SuccessInsert())
}

func (Bb BrandPres) GetBrand(c *gin.Context) {
	id := c.Query("id")
	result, err := Bb.Pres.BussGet(id)
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG(err.Error()))
		return
	}

	c.JSON(helper.SuccessGetData(dto.FromCore(result)))
}
