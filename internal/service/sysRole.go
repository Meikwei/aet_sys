package service

import (
	"context"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	sysV1 "github.com/Meikwei/aet_sys/api/sys/v1"
	"github.com/Meikwei/aet_sys/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		sysV1.RegisterSysRoleServer(server, NewSysRoleServer()) // register service to the rpc service
	})
}

var _ sysV1.SysRoleServer = (*sysRole)(nil)
var _ time.Time

type sysRole struct {
	sysV1.UnimplementedSysRoleServer

	// iDao dao.SysRoleDao
}

// NewSysRoleServer create a new service
func NewSysRoleServer() sysV1.SysRoleServer {
	return &sysRole{
		// iDao: dao.NewSysRoleDao(
		// 	model.GetDB(),
		// 	cache.NewSysRoleCache(model.GetCacheType()),
		// ),
	}
}

// Create a record
func (s *sysRole) Create(ctx context.Context, req *sysV1.CreateSysRoleRequest) (*sysV1.CreateSysRoleReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// record := &model.SysRole{}
	// err = copier.Copy(record, req)
	// if err != nil {
	// 	return nil, ecode.StatusCreateSysRole.Err()
	// }

	// err = s.iDao.Create(ctx, record)
	// if err != nil {
	// 	logger.Error("Create error", logger.Err(err), logger.Any("sysRole", record), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// return &sysV1.CreateSysRoleReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *sysRole) DeleteByID(ctx context.Context, req *sysV1.DeleteSysRoleByIDRequest) (*sysV1.DeleteSysRoleByIDReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// err = s.iDao.DeleteByID(ctx, req.Id)
	// if err != nil {
	// 	logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// return &sysV1.DeleteSysRoleByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *sysRole) DeleteByIDs(ctx context.Context, req *sysV1.DeleteSysRoleByIDsRequest) (*sysV1.DeleteSysRoleByIDsReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// err = s.iDao.DeleteByIDs(ctx, req.Ids)
	// if err != nil {
	// 	logger.Error("DeleteByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// return &sysV1.DeleteSysRoleByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *sysRole) UpdateByID(ctx context.Context, req *sysV1.UpdateSysRoleByIDRequest) (*sysV1.UpdateSysRoleByIDReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// record := &model.SysRole{}
	// err = copier.Copy(record, req)
	// if err != nil {
	// 	return nil, ecode.StatusUpdateByIDSysRole.Err()
	// }
	// record.ID = req.Id

	// err = s.iDao.UpdateByID(ctx, record)
	// if err != nil {
	// 	logger.Error("UpdateByID error", logger.Err(err), logger.Any("sysRole", record), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// return &sysV1.UpdateSysRoleByIDReply{}, nil
}

// GetByID get a record by id
func (s *sysRole) GetByID(ctx context.Context, req *sysV1.GetSysRoleByIDRequest) (*sysV1.GetSysRoleByIDReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// record, err := s.iDao.GetByID(ctx, req.Id)
	// if err != nil {
	// 	if errors.Is(err, model.ErrRecordNotFound) {
	// 		logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
	// 		return nil, ecode.StatusNotFound.Err()
	// 	}
	// 	logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// data, err := convertSysRole(record)
	// if err != nil {
	// 	logger.Warn("convertSysRole error", logger.Err(err), logger.Any("sysRole", record), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusGetByIDSysRole.Err()
	// }

	// return &sysV1.GetSysRoleByIDReply{SysRole: data}, nil
}

// GetByCondition get a record by id
func (s *sysRole) GetByCondition(ctx context.Context, req *sysV1.GetSysRoleByConditionRequest) (*sysV1.GetSysRoleByConditionReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// conditions := &query.Conditions{}
	// for _, v := range req.Conditions.GetColumns() {
	// 	column := query.Column{}
	// 	_ = copier.Copy(&column, v)
	// 	conditions.Columns = append(conditions.Columns, column)
	// }
	// err = conditions.CheckValid()
	// if err != nil {
	// 	logger.Warn("Parameters error", logger.Err(err), logger.Any("conditions", conditions), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }

	// record, err := s.iDao.GetByCondition(ctx, conditions)
	// if err != nil {
	// 	if errors.Is(err, model.ErrRecordNotFound) {
	// 		logger.Warn("GetByCondition error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 		return nil, ecode.StatusNotFound.Err()
	// 	}
	// 	logger.Error("GetByCondition error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// data, err := convertSysRole(record)
	// if err != nil {
	// 	logger.Warn("convertSysRole error", logger.Err(err), logger.Any("sysRole", record), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusGetByConditionSysRole.Err()
	// }

	// return &sysV1.GetSysRoleByConditionReply{
	// 	SysRole: data,
	// }, nil
}

// ListByIDs list of records by batch id
func (s *sysRole) ListByIDs(ctx context.Context, req *sysV1.ListSysRoleByIDsRequest) (*sysV1.ListSysRoleByIDsReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// sysRoleMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	// if err != nil {
	// 	logger.Error("GetByIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// sysRoles := []*sysV1.SysRole{}
	// for _, id := range req.Ids {
	// 	if v, ok := sysRoleMap[id]; ok {
	// 		record, err := convertSysRole(v)
	// 		if err != nil {
	// 			logger.Warn("convertSysRole error", logger.Err(err), logger.Any("sysRole", v), interceptor.ServerCtxRequestIDField(ctx))
	// 			return nil, ecode.StatusInternalServerError.ToRPCErr()
	// 		}
	// 		sysRoles = append(sysRoles, record)
	// 	}
	// }

	// return &sysV1.ListSysRoleByIDsReply{SysRoles: sysRoles}, nil
}

// ListByLastID list sysRole by last id
func (s *sysRole) ListByLastID(ctx context.Context, req *sysV1.ListSysRoleByLastIDRequest) (*sysV1.ListSysRoleByLastIDReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.CtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// if req.LastID == 0 {
	// 	req.LastID = math.MaxInt32
	// }
	// if req.Limit == 0 {
	// 	req.Limit = 10
	// }

	// records, err := s.iDao.GetByLastID(ctx, req.LastID, int(req.Limit), req.Sort)
	// if err != nil {
	// 	logger.Error("ListByLastID error", logger.Err(err), interceptor.CtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// sysRoles := []*sysV1.SysRole{}
	// for _, record := range records {
	// 	data, err := convertSysRole(record)
	// 	if err != nil {
	// 		logger.Warn("convertSysRole error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
	// 		continue
	// 	}
	// 	sysRoles = append(sysRoles, data)
	// }

	// return &sysV1.ListSysRoleByLastIDReply{
	// 	SysRoles: sysRoles,
	// }, nil
}

// List of records by query parameters
func (s *sysRole) List(ctx context.Context, req *sysV1.ListSysRoleRequest) (*sysV1.ListSysRoleReply, error) {
	panic("implement me")
	// err := req.Validate()
	// if err != nil {
	// 	logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInvalidParams.Err()
	// }
	// ctx = interceptor.WrapServerCtx(ctx)

	// params := &query.Params{}
	// err = copier.Copy(params, req.Params)
	// if err != nil {
	// 	return nil, ecode.StatusListSysRole.Err()
	// }
	// params.Size = int(req.Params.Limit)

	// records, total, err := s.iDao.GetByColumns(ctx, params)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "query params error:") {
	// 		logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
	// 		return nil, ecode.StatusInvalidParams.Err()
	// 	}
	// 	logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
	// 	return nil, ecode.StatusInternalServerError.ToRPCErr()
	// }

	// sysRoles := []*sysV1.SysRole{}
	// for _, record := range records {
	// 	data, err := convertSysRole(record)
	// 	if err != nil {
	// 		logger.Warn("convertSysRole error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
	// 		continue
	// 	}
	// 	sysRoles = append(sysRoles, data)
	// }

	// return &sysV1.ListSysRoleReply{
	// 	Total:    total,
	// 	SysRoles: sysRoles,
	// }, nil
}

func convertSysRole(record *model.SysRole) (*sysV1.SysRole, error) {
	value := &sysV1.SysRole{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	value.Id = record.ID
	// todo if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.CreatedAt = record.CreatedAt.Format(time.RFC3339)
	value.UpdatedAt = record.UpdatedAt.Format(time.RFC3339)

	return value, nil
}
