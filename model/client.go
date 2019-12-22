package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Client struct {
	Name string `json:"name"`
	gorm.Model
}

func (c *Client) Save(db *gorm.DB) (*Client, error) {
	err := db.Debug().Create(&c).Error

	if err != nil {
		return &Client{}, err
	}

	db.Save(&c)
	return c, nil
}

func (c *Client) Update(db *gorm.DB, id uint64) (*Client, error) {
	db = db.Debug().Model(&Client{}).Where("id=?", id).Take(&Client{}).UpdateColumns(
		map[string]interface{}{
			"name": c.Name,
		},
	)

	if db.Error != nil {
		return &Client{}, db.Error
	}

	err := db.Debug().Model(&Client{}).Where("id=?", id).Take(&c).Error
	if err != nil {
		return &Client{}, err
	}

	return c, nil
}

func (C *Client) FindAll(db *gorm.DB) (*[]Client, error) {
	var err error
	clients := []Client{}
	err = db.Debug().Model(&Client{}).Limit(100).Find(&clients).Error

	if err != nil {
		return &[]Client{}, err
	}

	return &clients, err
}
