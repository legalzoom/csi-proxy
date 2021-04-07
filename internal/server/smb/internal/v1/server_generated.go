// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/api/smb/v1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server/smb/internal"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v1")

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1.RegisterSmbServer(grpcServer, s)
}

func (s *versionedAPI) NewSmbGlobalMapping(context context.Context, versionedRequest *v1.NewSmbGlobalMappingRequest) (*v1.NewSmbGlobalMappingResponse, error) {
	request := &internal.NewSmbGlobalMappingRequest{}
	if err := Convert_v1_NewSmbGlobalMappingRequest_To_internal_NewSmbGlobalMappingRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.NewSmbGlobalMapping(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.NewSmbGlobalMappingResponse{}
	if err := Convert_internal_NewSmbGlobalMappingResponse_To_v1_NewSmbGlobalMappingResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) RemoveSmbGlobalMapping(context context.Context, versionedRequest *v1.RemoveSmbGlobalMappingRequest) (*v1.RemoveSmbGlobalMappingResponse, error) {
	request := &internal.RemoveSmbGlobalMappingRequest{}
	if err := Convert_v1_RemoveSmbGlobalMappingRequest_To_internal_RemoveSmbGlobalMappingRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.RemoveSmbGlobalMapping(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.RemoveSmbGlobalMappingResponse{}
	if err := Convert_internal_RemoveSmbGlobalMappingResponse_To_v1_RemoveSmbGlobalMappingResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
