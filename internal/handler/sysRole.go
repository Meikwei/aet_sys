package handler

import (
	"context"

	sysV1 "github.com/Meikwei/aet_sys/api/sys/v1"
	"github.com/Meikwei/aet_sys/internal/service"
)

var _ sysV1.SysRoleLogicer = (*sysRoleHandler)(nil)

type sysRoleHandler struct {
	server sysV1.SysRoleServer
}

// NewSysRoleHandler create a handler
func NewSysRoleHandler() sysV1.SysRoleLogicer {
	return &sysRoleHandler{
		server: service.NewSysRoleServer(),
	}
}

// Create a record
func (h *sysRoleHandler) Create(ctx context.Context, req *sysV1.CreateSysRoleRequest) (*sysV1.CreateSysRoleReply, error) {
	return h.server.Create(ctx, req)
}

// DeleteByID delete a record by id
func (h *sysRoleHandler) DeleteByID(ctx context.Context, req *sysV1.DeleteSysRoleByIDRequest) (*sysV1.DeleteSysRoleByIDReply, error) {
	return h.server.DeleteByID(ctx, req)
}

// DeleteByIDs delete records by batch id
func (h *sysRoleHandler) DeleteByIDs(ctx context.Context, req *sysV1.DeleteSysRoleByIDsRequest) (*sysV1.DeleteSysRoleByIDsReply, error) {
	return h.server.DeleteByIDs(ctx, req)
}

// UpdateByID update a record by id
func (h *sysRoleHandler) UpdateByID(ctx context.Context, req *sysV1.UpdateSysRoleByIDRequest) (*sysV1.UpdateSysRoleByIDReply, error) {
	return h.server.UpdateByID(ctx, req)
}

// GetByID get a record by id
func (h *sysRoleHandler) GetByID(ctx context.Context, req *sysV1.GetSysRoleByIDRequest) (*sysV1.GetSysRoleByIDReply, error) {
	return h.server.GetByID(ctx, req)
}

// GetByCondition get a record by condition
func (h *sysRoleHandler) GetByCondition(ctx context.Context, req *sysV1.GetSysRoleByConditionRequest) (*sysV1.GetSysRoleByConditionReply, error) {
	return h.server.GetByCondition(ctx, req)
}

// ListByIDs list of records by batch id
func (h *sysRoleHandler) ListByIDs(ctx context.Context, req *sysV1.ListSysRoleByIDsRequest) (*sysV1.ListSysRoleByIDsReply, error) {
	return h.server.ListByIDs(ctx, req)
}

// ListByLastID get records by last id
func (h *sysRoleHandler) ListByLastID(ctx context.Context, req *sysV1.ListSysRoleByLastIDRequest) (*sysV1.ListSysRoleByLastIDReply, error) {
	return h.server.ListByLastID(ctx, req)
}

// List of records by query parameters
func (h *sysRoleHandler) List(ctx context.Context, req *sysV1.ListSysRoleRequest) (*sysV1.ListSysRoleReply, error) {
	return h.server.List(ctx, req)
}
