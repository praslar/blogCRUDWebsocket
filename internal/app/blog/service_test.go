package blog_test

import (
	"fmt"
	reflect "reflect"
	"testing"

	"github.com/blogCRUDWebsocket/internal/app/blog"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	Mockdatabase := NewMockdatabase(gomock.NewController(t))
	MockMonitorRoom := NewMockMonitorRoom(gomock.NewController(t))
	service := blog.NewService(Mockdatabase, MockMonitorRoom)
	testCases := []struct {
		desc   string
		expect func()
		input  *blog.CreateBlog
		err    error
	}{
		{
			desc:   "Empty Value",
			expect: func() {},
			input: &blog.CreateBlog{
				Content: "",
				Title:   "",
			},
			err: fmt.Errorf("Empty value! %s", ""),
		},
		{
			desc: "Created",
			expect: func() {
				Mockdatabase.EXPECT().Create(gomock.Any()).Return(nil)
				MockMonitorRoom.EXPECT().Write(gomock.Any())
			},
			input: &blog.CreateBlog{
				Content: "test",
				Title:   "test",
			},
			err: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.expect()
			_, err := service.Create(tC.input)
			if !reflect.DeepEqual(err, tC.err) {
				t.Error("Got", err, " but want", tC.err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	Mockdatabase := NewMockdatabase(gomock.NewController(t))
	MockMonitorRoom := NewMockMonitorRoom(gomock.NewController(t))
	service := blog.NewService(Mockdatabase, MockMonitorRoom)
	testCases := []struct {
		desc   string
		expect func()
		input1 string
		input2 *blog.CreateBlog
		err    error
	}{
		{
			desc: "Not found blog",
			expect: func() {
				Mockdatabase.EXPECT().FindByID("testid").Return(nil, fmt.Errorf("Not found"))
			},
			input1: "testid",
			input2: &blog.CreateBlog{
				Content: "test",
				Title:   "test",
			},
			err: fmt.Errorf("Blog not found"),
		},
		{
			desc: "Update",
			expect: func() {
				Mockdatabase.EXPECT().FindByID(gomock.Any()).Return(&blog.Blog{}, nil)
				Mockdatabase.EXPECT().Update(gomock.Any(), gomock.Any())
				MockMonitorRoom.EXPECT().Write(gomock.Any())
			},
			input1: "testid",
			input2: &blog.CreateBlog{
				Content: "test",
				Title:   "test",
			},
			err: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.expect()
			_, err := service.Update(tC.input1, tC.input2)
			if !reflect.DeepEqual(err, tC.err) {
				t.Error("Got", err, " but want", tC.err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	Mockdatabase := NewMockdatabase(gomock.NewController(t))
	MockMonitorRoom := NewMockMonitorRoom(gomock.NewController(t))
	service := blog.NewService(Mockdatabase, MockMonitorRoom)
	testCases := []struct {
		desc   string
		expect func()
		input  string
		err    error
	}{
		{
			desc: "Not found blog",
			expect: func() {
				Mockdatabase.EXPECT().FindByID("testid").Return(nil, fmt.Errorf("Not found"))
			},
			input: "testid",
			err:   fmt.Errorf("Blog not found"),
		},
		{
			desc: "Deleted",
			expect: func() {
				Mockdatabase.EXPECT().FindByID(gomock.Any()).Return(&blog.Blog{}, nil)
				MockMonitorRoom.EXPECT().Write(gomock.Any())
				Mockdatabase.EXPECT().Delete(gomock.Any())
			},
			input: "testid",
			err:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.expect()
			err := service.Delete(tC.input)
			if !reflect.DeepEqual(err, tC.err) {
				t.Error("Got", err, " but want", tC.err)
			}
		})
	}
}

func TestRead(t *testing.T) {
	Mockdatabase := NewMockdatabase(gomock.NewController(t))
	MockMonitorRoom := NewMockMonitorRoom(gomock.NewController(t))
	service := blog.NewService(Mockdatabase, MockMonitorRoom)
	testCases := []struct {
		desc   string
		expect func()
		err    error
	}{
		{
			desc: "Empty Blog",
			expect: func() {
				Mockdatabase.EXPECT().Read().Return(nil, fmt.Errorf("Empty Blog"))
			},
			err: fmt.Errorf("Empty Blog"),
		},
		{
			desc: "Read",
			expect: func() {
				Mockdatabase.EXPECT().Read().Return([]blog.Blog{}, nil)
			},
			err: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.expect()
			_, err := service.Read()
			if !reflect.DeepEqual(err, tC.err) {
				t.Error("Got", err, " but want", tC.err)
			}
		})
	}
}
