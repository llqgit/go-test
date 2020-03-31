package main

import (
	"bytes"
	"fmt"

	//"encoding/csv"
	"github.com/artonge/go-csv-tag/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
	//"strings"
)

type Demo struct { // A structure with tags
	Name string  `csv:"name"`
	ID   int     `csv:"ID"`
	Num  float64 `csv:"number"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {

		fileName := "test.csv"

		tab := []Demo{ // Create the slice where to put the content
			{
				Name: "some name",
				ID:   1,
				Num:  42.5,
			},
		}

		writer := new(bytes.Buffer)
		_ = csvtag.DumpToWriter(tab, writer)

		extraHeaders := map[string]string{
			"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, fileName),
		}
		//ctx.Header("Content-Disposition", `attachment; filename="test.csv"`)
		ctx.DataFromReader(http.StatusOK, int64(writer.Len()), "text/csv", writer, extraHeaders)

		//ctx.DataFromReader(http.StatusOK, int64(length), "text/csv", file, extraHeaders)
		//ctx.DataFromReader()
		//ctx.RespFile(data, "csv")
		//return ctx.FileCsv(data)
		//ctx.JSON(http.StatusOK, "ok")
	})

	_ = router.Run(":8088")
}
