package main

import (
	"context"

	"github.com/sjxiang/biz-demo/easy-note/cmd/note/pack"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/service"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

// 请求参数校验，错误统一处理
type NoteServiceImpl struct {
	pb.UnimplementedNoteServiceServer
}


func (impl *NoteServiceImpl) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (resp *pb.CreateNoteResponse, err error) {
	resp = new(pb.CreateNoteResponse)
	if req.UserId <= 0 || len(req.Title) == 0 || len(req.Content) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil 
	}

	err = service.NewCreateNoteService(ctx).CreateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
	

func (impl *NoteServiceImpl) MGetNote(ctx context.Context, req *pb.MGetNoteRequest) (resp *pb.MGetNoteResponse, err error) {
	resp = new(pb.MGetNoteResponse)
	
	if len(req.NoteIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil 
	}

	notes, err := service.NewMGetNoteService(ctx).MGetNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	return resp, nil 
}


func (impl *NoteServiceImpl) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (resp *pb.DeleteNoteResponse, err error) {
	resp = new(pb.DeleteNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDelNoteService(ctx).DelNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)	
	return resp, nil
}
	

func (impl *NoteServiceImpl) QueryNote(ctx context.Context, req *pb.QueryNoteRequest) (resp *pb.QueryNoteResponse, err error) {
	resp = new(pb.QueryNoteResponse)
	
	if req.UserId <= 0 || req.Limit < 0 || req.Offset < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.Limit == 0 {
		req.Limit = consts.DefaultLimit
	}

	notes, total, err := service.NewQueryNoteService(ctx).QueryNoteService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	resp.Total = total
	return resp, nil
}


func (impl *NoteServiceImpl) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (resp *pb.UpdateNoteResponse, err error) {
	resp = new(pb.UpdateNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateNoteService(ctx).UpdateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)	
	return resp, nil
}
