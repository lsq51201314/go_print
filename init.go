package main

import (
	"encoding/base64"
	"os"
)

var CurrentDir string

type NameValue struct {
	Name  string
	Value string
}

var TagsData []NameValue = make([]NameValue, 0)
var RowsData [][]NameValue = make([][]NameValue, 0)

func init() {
	var err error
	if CurrentDir, err = os.Getwd(); err != nil {
		panic(err)
	}
}

func init() {
	TagsData = append(TagsData, NameValue{
		Name:  "[报告类型]",
		Value: "售 后 服 务 收 费 报 告",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[呈报日期]",
		Value: "2024-11-11",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[呈报类型]",
		Value: "售后服务",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[所属店铺]",
		Value: "天猫旗舰店",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[总的金额]",
		Value: "1688.00",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[单据编号]",
		Value: "147258369123",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[呈报人]",
		Value: "奔跑的蜗牛",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[事件描述]",
		Value: "没什么，我就想提交一个报告。",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[货品图片0]",
		Value: CurrentDir + "\\001.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[货品图片1]",
		Value: CurrentDir + "\\002.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[货品图片2]",
		Value: CurrentDir + "\\003.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[货品图片3]",
		Value: CurrentDir + "\\004.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[经理签字]",
		Value: CurrentDir + "\\005.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[联络签字]",
		Value: CurrentDir + "\\006.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[主管签字]",
		Value: CurrentDir + "\\007.jpg",
	})
	TagsData = append(TagsData, NameValue{
		Name:  "[货主签字]",
		Value: CurrentDir + "\\008.jpg",
	})
}

func init() {
	img1, _ := os.ReadFile("001.jpg")
	img2, _ := os.ReadFile("002.jpg")
	RowsData = append(RowsData, []NameValue{
		{
			Name:  "[货品编号]",
			Value: "654121245222",
		},
		{
			Name:  "[加金费用]",
			Value: "256",
		},
		{
			Name:  "[扣点]",
			Value: "10%",
		},
		{
			Name:  "[结算金额]",
			Value: "188.00",
		},
		{
			Name:  "[收款日期]",
			Value: "2024-11-18",
		},
		{
			Name:  "[货主名称]",
			Value: "奔跑的蜗牛",
		},
		{
			Name:  "[截图图片]",
			Value: base64.StdEncoding.EncodeToString(img1),
		},
	})
	RowsData = append(RowsData, []NameValue{
		{
			Name:  "[货品编号]",
			Value: "669851122121",
		},
		{
			Name:  "[加金费用]",
			Value: "128",
		},
		{
			Name:  "[扣点]",
			Value: "5%",
		},
		{
			Name:  "[结算金额]",
			Value: "1688.00",
		},
		{
			Name:  "[收款日期]",
			Value: "2024-11-11",
		},
		{
			Name:  "[货主名称]",
			Value: "老污龟",
		},
		{
			Name:  "[截图图片]",
			Value: base64.StdEncoding.EncodeToString(img2),
		},
	})
}
