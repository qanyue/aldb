package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/model"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	r := SetRouter()
	w := httptest.NewRecorder()
	f, err := excelize.OpenFile("D:/WorkSpace/aldb/server/Tag.xlsx")
	if err != nil {
		t.Error("Open xlsx file error")
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("tag")
	if err != nil {
		fmt.Println(err)
		return
	}
	var tag model.Tag
	for _, row := range rows {
		tag.ResourceName = strings.Trim(row[0], " ")
		tag.Name = row[1]
		marshalled, err := json.Marshal(tag)
		if err != nil {
			log.Fatalf("impossible to marshall teacher: %s", err)
		}
		req, err := http.NewRequest("POST", "/api/tag/add", bytes.NewReader(marshalled))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatalf("impossible to build request: %s", err)
		}
		r.ServeHTTP(w, req)
		fmt.Println("good---------------------")
		resBody, err := io.ReadAll(w.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}
		//log.Printf("res body: %s", string(resBody))

		fmt.Printf("%#v, body: %#v \n", tag, string(resBody))
		assert.Equal(t, 200, w.Code)
	}
	//
	//r.ServeHTTP(w, req)
	//assert.Equal(t, "pong", w.Body.String())
}
func TestSetRouter(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
