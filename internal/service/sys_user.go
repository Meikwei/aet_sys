// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"

	"google.golang.org/grpc"

	//"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	//"github.com/zhufuyi/sponge/pkg/logger"

	sysV1 "github.com/Meikwei/aet_sys/api/sys/v1"
	//"github.com/Meikwei/aet_sys/internal/cache"
	//"github.com/Meikwei/aet_sys/internal/dao"
	//"github.com/Meikwei/aet_sys/internal/ecode"
	//"github.com/Meikwei/aet_sys/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		sysV1.RegisterSysUserServer(server, NewSysUserServer())
	})
}

var _ sysV1.SysUserServer = (*sysUser)(nil)

type sysUser struct {
	sysV1.UnimplementedSysUserServer

	// example:
	//		iDao dao.SysUserDao
}

// NewSysUserServer create a server
func NewSysUserServer() sysV1.SysUserServer {
	return &sysUser{
		// example:
		//		iDao: dao.NewSysUserDao(
		//			model.GetDB(),
		//			cache.NewSysUserCache(model.GetCacheType()),
		//		),
	}
}

// Create sysUser
func (s *sysUser) Create(ctx context.Context, req *sysV1.CreateSysUserRequest) (*sysV1.CreateSysUserReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.Create(ctx, &model.SysUser{
	//     	UserName: req.UserName,
	//     	UserNumber: req.UserNumber,
	//     	UserPassword: req.UserPassword,
	//     	UserPhone: req.UserPhone,
	//     	UserAvatar: req.UserAvatar,
	//     	CreateUser: req.CreateUser,
	//     	UpdateUser: req.UpdateUser,
	//     })
	// 	if err != nil {
	//			logger.Warn("Create error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.CreateSysUserReply{
	//     	Id: reply.Id,
	//     }, nil
}

// DeleteByID delete sysUser by id
func (s *sysUser) DeleteByID(ctx context.Context, req *sysV1.DeleteSysUserByIDRequest) (*sysV1.DeleteSysUserByIDReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.DeleteByID(ctx, &model.SysUser{
	//     	Id: req.Id,
	//     })
	// 	if err != nil {
	//			logger.Warn("DeleteByID error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.DeleteSysUserByIDReply{
	//     }, nil
}

// DeleteByIDs delete sysUser by batch id
func (s *sysUser) DeleteByIDs(ctx context.Context, req *sysV1.DeleteSysUserByIDsRequest) (*sysV1.DeleteSysUserByIDsReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.DeleteByIDs(ctx, &model.SysUser{
	//     	Ids: req.Ids,
	//     })
	// 	if err != nil {
	//			logger.Warn("DeleteByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.DeleteSysUserByIDsReply{
	//     }, nil
}

// UpdateByID update sysUser by id
func (s *sysUser) UpdateByID(ctx context.Context, req *sysV1.UpdateSysUserByIDRequest) (*sysV1.UpdateSysUserByIDReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.UpdateByID(ctx, &model.SysUser{
	//     	Id: req.Id,
	//     	UserName: req.UserName,
	//     	UserNumber: req.UserNumber,
	//     	UserPassword: req.UserPassword,
	//     	UserPhone: req.UserPhone,
	//     	UserAvatar: req.UserAvatar,
	//     	CreateUser: req.CreateUser,
	//     	UpdateUser: req.UpdateUser,
	//     })
	// 	if err != nil {
	//			logger.Warn("UpdateByID error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.UpdateSysUserByIDReply{
	//     }, nil
}

// GetByID get sysUser by id
func (s *sysUser) GetByID(ctx context.Context, req *sysV1.GetSysUserByIDRequest) (*sysV1.GetSysUserByIDReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.GetByID(ctx, &model.SysUser{
	//     	Id: req.Id,
	//     })
	// 	if err != nil {
	//			logger.Warn("GetByID error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.GetSysUserByIDReply{
	//     	SysUser: reply.SysUser,
	//     }, nil
}

// GetByCondition get sysUser by condition
func (s *sysUser) GetByCondition(ctx context.Context, req *sysV1.GetSysUserByConditionRequest) (*sysV1.GetSysUserByConditionReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.GetByCondition(ctx, &model.SysUser{
	//     	Conditions: req.Conditions,
	//     })
	// 	if err != nil {
	//			logger.Warn("GetByCondition error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.GetSysUserByConditionReply{
	//     	SysUser: reply.SysUser,
	//     }, nil
}

// ListByIDs list of sysUser by batch id
func (s *sysUser) ListByIDs(ctx context.Context, req *sysV1.ListSysUserByIDsRequest) (*sysV1.ListSysUserByIDsReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.ListByIDs(ctx, &model.SysUser{
	//     	Ids: req.Ids,
	//     })
	// 	if err != nil {
	//			logger.Warn("ListByIDs error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.ListSysUserByIDsReply{
	//     	SysUsers: reply.SysUsers,
	//     }, nil
}

// ListByLastID list sysUser by last id
func (s *sysUser) ListByLastID(ctx context.Context, req *sysV1.ListSysUserByLastIDRequest) (*sysV1.ListSysUserByLastIDReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.ListByLastID(ctx, &model.SysUser{
	//     	LastID: req.LastID,
	//     	Limit: req.Limit,
	//     	Sort: req.Sort,
	//     })
	// 	if err != nil {
	//			logger.Warn("ListByLastID error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.ListSysUserByLastIDReply{
	//     	SysUsers: reply.SysUsers,
	//     }, nil
}

// List of sysUser by query parameters
func (s *sysUser) List(ctx context.Context, req *sysV1.ListSysUserRequest) (*sysV1.ListSysUserReply, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.List(ctx, &model.SysUser{
	//     	Params: req.Params,
	//     })
	// 	if err != nil {
	//			logger.Warn("List error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &sysV1.ListSysUserReply{
	//     	Total: reply.Total,
	//     	SysUsers: reply.SysUsers,
	//     }, nil
}