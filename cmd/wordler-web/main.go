package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/forewing/wordler"
	"github.com/gin-gonic/gin"
)

type reqQuery struct {
	Has []string `json:"has" form:"has"`
	No  []string `json:"no" form:"no"`
	At  []string `json:"at" form:"at"`
	Na  []string `json:"na" form:"na"`

	Len int `json:"len" form:"len"`
	Max int `json:"max" form:"max"`
}

type rspError struct {
	Error string `json:"error"`
}

var (
	flagBind = flag.String("bind", "0.0.0.0:8080", "bind address")
	flagMax  = flag.Int("max", 1000, "max response length")
)

func init() {
	flag.Parse()

	if *flagMax <= 0 {
		panic("flag `-max` should be greater than 0")
	}
}

func main() {
	r := gin.Default()
	r.GET("/filter", handleFilter)
	r.POST("/filter", handleFilter)

	r.SetHTMLTemplate(mustLoadTemplate())
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.StaticFS("/statics", http.FS(statics))

	log.Println("Listening on", "http://"+*flagBind)
	log.Println(r.Run(*flagBind))
}

func handleFilter(c *gin.Context) {
	var req reqQuery

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, rspError{Error: err.Error()})
		return
	}

	wl, err := wordler.GetWordList(req.Len)
	if err != nil {
		c.JSON(http.StatusBadRequest, rspError{Error: err.Error()})
		return
	}

	if req.Max <= 0 || req.Max > *flagMax {
		req.Max = *flagMax
	}

	filters := wordler.FilterList{}

	for _, has := range req.Has {
		filters = append(filters, wordler.FilterContain{Target: has})
	}

	for _, no := range req.No {
		filters = append(filters, wordler.FilterNotContain{Target: no})
	}

	for _, at := range req.At {
		index, char, err := wordler.ParseAt(at)
		if err != nil {
			c.JSON(http.StatusBadRequest, rspError{Error: err.Error()})
			return
		}
		filters = append(filters, wordler.FilterAt{Target: char, Index: index})
	}

	for _, na := range req.Na {
		index, char, err := wordler.ParseAt(na)
		if err != nil {
			c.JSON(http.StatusBadRequest, rspError{Error: err.Error()})
			return
		}
		filters = append(filters, wordler.FilterNotAt{Target: char, Index: index})
	}

	result := filters.Run(wl)
	if len(result) > req.Max {
		result = result[0:req.Max]
	}

	c.JSON(http.StatusOK, result)
}
