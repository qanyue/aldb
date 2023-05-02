package database

import (
	"reflect"
	"testing"
)

func TestMgo_InsertTag(t *testing.T) {
	type args struct {
		t *Tag
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"小球藻测试", args{&Tag{Name: "小球藻（淡水）", ResourceName: "Chlorella sp."}}, false},
		{"海带配子体测试", args{&Tag{Name: "海带配子体", ResourceName: "Saccharina japonica"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mgo{}
			if err := m.InsertTag(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("InsertTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMgo_QueryTagByName(t *testing.T) {
	type args struct {
		obj string
	}
	tests := []struct {
		name string
		args args
		want Tag
	}{
		{"小球藻查询测试", args{"小球藻（淡水）"}, Tag{Name: "小球藻（淡水）", ResourceName: "Chlorella sp."}},
		{name: "海带配子体查询测试", want: Tag{Name: "海带配子体", ResourceName: "Saccharina japonica"}, args: args{"海带配子体"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mgo{}
			if got := m.QueryTagByName(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryTagByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMgo_ExistsTag(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"小球藻查询测试", args{"小球藻（淡水）"}, true},
		{name: "海带配子体查询测试", want: true, args: args{"海带配子体"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mgo{}
			if got := m.ExistsTag(tt.args.name); got != tt.want {
				t.Errorf("ExistsTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMgo_DropTag(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"小球藻删除测试", args{"小球藻（淡水）"}, false},
		{name: "海带配子体删除测试", wantErr: false, args: args{"海带配子体"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mgo{}
			if err := m.DropTag(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("DropTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
