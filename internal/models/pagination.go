package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func NewPagination(c *gin.Context) *Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	logrus.WithFields(logrus.Fields{
		"limit": limit,
		"page":  page,
	}).Info("NewPagination")
	return &Pagination{
		Limit: limit,
		Page:  page,
	}
}

func DefaultPagination() *Pagination {
	return &Pagination{
		Limit: 10,
		Page:  1,
	}
}

func (p *Pagination) GetScope(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	logrus.WithFields(logrus.Fields{
		"limit": p.Limit,
		"page":  p.Page,
	}).Info("GetScope")
	return db.Offset(offset).
		Limit(p.Limit)
}
