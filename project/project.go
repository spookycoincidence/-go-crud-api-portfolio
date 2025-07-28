package project

import "gorm.io/gorm"

type Project struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
