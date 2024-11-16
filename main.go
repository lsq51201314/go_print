package main

import (
	"fmt"
	"strings"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	//初始化COM组件
	if err := ole.CoInitialize(0); err != nil {
		panic(err)
	}
	defer ole.CoUninitialize()
	//gregn.GridppReport.6
	//gregn.GridppReport
	unknown, err := oleutil.CreateObject("gregn.GridppReport")
	if err != nil {
		panic(err)
	}
	defer unknown.Release()

	gregn, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		panic(err)
	}
	defer gregn.Release()

	//01.载入文件
	//LoadFromStr
	//LoadFromURL
	if _, err := oleutil.CallMethod(gregn, "LoadFromFile", CurrentDir+"\\test.grf"); err != nil {
		panic(err)
	}
	//02.处理标签（按我这边的使用习惯仅仅处理 图像框、综合文字框、条形码或二维码）
	// grctStaticBox 		1 静态文字框。
	// grctShapeBox 		2 图形框。
	// grctSystemVarBox 	3 系统变量框。
	// grctFieldBox 		4 字段框。
	// grctSummaryBox	 	5 统计框。
	// grctRichTextBox 		6 RTF格式文字框
	// grctPictureBox 		7 图像框。
	// grctMemoBox 			8 综合文字框。
	// grctSubReport 		9 子报表。
	// grctLine 			10 线段。
	// grctChart 			11 图表。
	// grctBarcode 			12 条形码。
	// grctFreeGrid 		13 自由表格。
	control, err := oleutil.CallMethod(gregn, "FindFirstControl") //寻找第一个控件，必须
	if err != nil {
		panic(err)
	}
	for control.ToIDispatch() != nil {
		//取标记
		tag, err := control.ToIDispatch().GetProperty("Tag")
		if err != nil {
			panic(err)
		}
		tag_str := tag.Value().(string)
		//替换标签
		if tag_str != "" {
			fmt.Print("标识内容：", tag_str)
			//取类型
			typ, err := control.ToIDispatch().GetProperty("ControlType")
			if err != nil {
				panic(err)
			}
			tid := typ.Value().(int32)
			//替换内容
			for _, v := range TagsData {
				tag_str = strings.ReplaceAll(tag_str, v.Name, v.Value)
			}
			//设置值
			fmt.Println("\t\t" + tag_str)
			switch tid {
			case 7: //图像框
				if _, err := control.ToIDispatch().PutProperty("ImageFile", tag_str); err != nil {
					panic(err)
				}
			case 8: //综合文字框
				if _, err := control.ToIDispatch().PutProperty("Text", tag_str); err != nil {
					panic(err)
				}
			case 12: //条形码或二维码
				if _, err := control.ToIDispatch().PutProperty("Text", tag_str); err != nil {
					panic(err)
				}
			default:
				//不处理其他类型
				fmt.Println("控件类型：", tid)
			}
		}
		//取下一个
		if control, err = oleutil.CallMethod(gregn, "FindNextControl"); err != nil { //寻找下一个控件，必须
			panic(err)
		}
	}
	//03.处理表格
	//Report.DetailGrid.Recordset.Fields.Count;
	detailGrid, err := gregn.GetProperty("DetailGrid")
	if err != nil {
		panic(err)
	}
	recordset, err := detailGrid.ToIDispatch().GetProperty("Recordset")
	if err != nil {
		panic(err)
	}
	fields, err := recordset.ToIDispatch().GetProperty("Fields")
	if err != nil {
		panic(err)
	}
	count, err := fields.ToIDispatch().GetProperty("Count")
	if err != nil {
		panic(err)
	}
	count_i := count.Value().(int32)
	fmt.Println(count_i)
	//构建数据（直接使用载入数据的方法不用监听事件）
	xml := "<xml>"
	for _, arr := range RowsData {
		row := "<row>"
		for i := 0; i < int(count_i); i++ {
			//单元格
			item, err := fields.ToIDispatch().CallMethod("ItemAt", i+1)
			if err != nil {
				panic(err)
			}
			//名称
			name, err := item.ToIDispatch().GetProperty("Name")
			if err != nil {
				panic(err)
			}
			name_str := name.Value().(string)
			//标识
			tag, err := item.ToIDispatch().GetProperty("Tag")
			if err != nil {
				panic(err)
			}
			tag_str := tag.Value().(string)
			//类型
			// grftString 		1 字符字段。
			// grftInteger 		2 整数字段。
			// grftFloat 		3 浮点数字段。
			// grftCurrency 	4 货币字段。
			// grftBoolean 		5 布尔字段。
			// grftDateTime 	6 日期时间字段。
			// grftBinary 		7 二进制字段。
			// fty, err := item.ToIDispatch().GetProperty("FieldType")
			// if err != nil {
			// 	panic(err)
			// }
			// fty_str := fty.Value().(int32)
			// fmt.Println(fty_str)
			//替换内容
			for _, v := range arr {
				tag_str = strings.ReplaceAll(tag_str, v.Name, v.Value)
			}
			row += fmt.Sprintf("<%s>%s</%s>", name_str, tag_str, name_str)
		}
		row += "</row>"
		xml += row
	}
	xml += "</xml>"
	load, err := recordset.ToIDispatch().CallMethod("LoadDataFromXML", xml)
	if err != nil {
		panic(err)
	}
	fmt.Println(load.Value())
	//打印文档
	//Report.Printer.PrinterName
	printer, err := gregn.GetProperty("Printer")
	if err != nil {
		panic(err)
	}
	//指定打印机
	if _, err := printer.ToIDispatch().PutProperty("PrinterName", "Microsoft Print to PDF"); err != nil {
		panic(err)
	}
	if _, err := oleutil.CallMethod(gregn, "Print", false); err != nil {
		panic(err)
	}
}
