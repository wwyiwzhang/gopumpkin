package yelper

import (
	"fmt"
)

type Businesses struct {
	Item   []BusinessItem `json:"businesses"`
	Total  int
	Region map[string]interface{}
}

type BusinessItem struct {
	Id          string
	Name        string
	ReviewCount int `json:"review_count"`
	Rating      float32
	Price       string
	Categories  []Category `json:"categories"`
}

type Category struct {
	Title string
}

func (b *Businesses) ConvertItem() [][]string {
	result := [][]string{}
	for _, i := range b.Item {
		sub := []string{
			i.Id,
			i.Name,
			fmt.Sprintf("%d", i.ReviewCount),
			fmt.Sprintf("%f", i.Rating),
			i.Price,
		}
		sub = append(sub, i.getCategory()...)
		result = append(result, sub)
	}
	return result
}

func (bi *BusinessItem) getCategory() []string {
	c := []string{}
	for _, i := range bi.Categories {
		c = append(c, i.Title)
	}
	return c
}
