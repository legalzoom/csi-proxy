// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	v1beta1 "github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta1"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/volume/internal"
)

func autoConvert_v1beta1_DismountVolumeRequest_To_internal_DismountVolumeRequest(in *v1beta1.DismountVolumeRequest, out *internal.DismountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Path = in.Path
	return nil
}

// Convert_v1beta1_DismountVolumeRequest_To_internal_DismountVolumeRequest is an autogenerated conversion function.
func Convert_v1beta1_DismountVolumeRequest_To_internal_DismountVolumeRequest(in *v1beta1.DismountVolumeRequest, out *internal.DismountVolumeRequest) error {
	return autoConvert_v1beta1_DismountVolumeRequest_To_internal_DismountVolumeRequest(in, out)
}

func autoConvert_internal_DismountVolumeRequest_To_v1beta1_DismountVolumeRequest(in *internal.DismountVolumeRequest, out *v1beta1.DismountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Path = in.Path
	return nil
}

// Convert_internal_DismountVolumeRequest_To_v1beta1_DismountVolumeRequest is an autogenerated conversion function.
func Convert_internal_DismountVolumeRequest_To_v1beta1_DismountVolumeRequest(in *internal.DismountVolumeRequest, out *v1beta1.DismountVolumeRequest) error {
	return autoConvert_internal_DismountVolumeRequest_To_v1beta1_DismountVolumeRequest(in, out)
}

func autoConvert_v1beta1_DismountVolumeResponse_To_internal_DismountVolumeResponse(in *v1beta1.DismountVolumeResponse, out *internal.DismountVolumeResponse) error {
	return nil
}

// Convert_v1beta1_DismountVolumeResponse_To_internal_DismountVolumeResponse is an autogenerated conversion function.
func Convert_v1beta1_DismountVolumeResponse_To_internal_DismountVolumeResponse(in *v1beta1.DismountVolumeResponse, out *internal.DismountVolumeResponse) error {
	return autoConvert_v1beta1_DismountVolumeResponse_To_internal_DismountVolumeResponse(in, out)
}

func autoConvert_internal_DismountVolumeResponse_To_v1beta1_DismountVolumeResponse(in *internal.DismountVolumeResponse, out *v1beta1.DismountVolumeResponse) error {
	return nil
}

// Convert_internal_DismountVolumeResponse_To_v1beta1_DismountVolumeResponse is an autogenerated conversion function.
func Convert_internal_DismountVolumeResponse_To_v1beta1_DismountVolumeResponse(in *internal.DismountVolumeResponse, out *v1beta1.DismountVolumeResponse) error {
	return autoConvert_internal_DismountVolumeResponse_To_v1beta1_DismountVolumeResponse(in, out)
}

func autoConvert_v1beta1_FormatVolumeRequest_To_internal_FormatVolumeRequest(in *v1beta1.FormatVolumeRequest, out *internal.FormatVolumeRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta1_FormatVolumeRequest_To_internal_FormatVolumeRequest is an autogenerated conversion function.
func Convert_v1beta1_FormatVolumeRequest_To_internal_FormatVolumeRequest(in *v1beta1.FormatVolumeRequest, out *internal.FormatVolumeRequest) error {
	return autoConvert_v1beta1_FormatVolumeRequest_To_internal_FormatVolumeRequest(in, out)
}

func autoConvert_internal_FormatVolumeRequest_To_v1beta1_FormatVolumeRequest(in *internal.FormatVolumeRequest, out *v1beta1.FormatVolumeRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_FormatVolumeRequest_To_v1beta1_FormatVolumeRequest is an autogenerated conversion function.
func Convert_internal_FormatVolumeRequest_To_v1beta1_FormatVolumeRequest(in *internal.FormatVolumeRequest, out *v1beta1.FormatVolumeRequest) error {
	return autoConvert_internal_FormatVolumeRequest_To_v1beta1_FormatVolumeRequest(in, out)
}

func autoConvert_v1beta1_FormatVolumeResponse_To_internal_FormatVolumeResponse(in *v1beta1.FormatVolumeResponse, out *internal.FormatVolumeResponse) error {
	return nil
}

// Convert_v1beta1_FormatVolumeResponse_To_internal_FormatVolumeResponse is an autogenerated conversion function.
func Convert_v1beta1_FormatVolumeResponse_To_internal_FormatVolumeResponse(in *v1beta1.FormatVolumeResponse, out *internal.FormatVolumeResponse) error {
	return autoConvert_v1beta1_FormatVolumeResponse_To_internal_FormatVolumeResponse(in, out)
}

func autoConvert_internal_FormatVolumeResponse_To_v1beta1_FormatVolumeResponse(in *internal.FormatVolumeResponse, out *v1beta1.FormatVolumeResponse) error {
	return nil
}

// Convert_internal_FormatVolumeResponse_To_v1beta1_FormatVolumeResponse is an autogenerated conversion function.
func Convert_internal_FormatVolumeResponse_To_v1beta1_FormatVolumeResponse(in *internal.FormatVolumeResponse, out *v1beta1.FormatVolumeResponse) error {
	return autoConvert_internal_FormatVolumeResponse_To_v1beta1_FormatVolumeResponse(in, out)
}

func autoConvert_v1beta1_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest(in *v1beta1.IsVolumeFormattedRequest, out *internal.IsVolumeFormattedRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta1_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest is an autogenerated conversion function.
func Convert_v1beta1_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest(in *v1beta1.IsVolumeFormattedRequest, out *internal.IsVolumeFormattedRequest) error {
	return autoConvert_v1beta1_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest(in, out)
}

func autoConvert_internal_IsVolumeFormattedRequest_To_v1beta1_IsVolumeFormattedRequest(in *internal.IsVolumeFormattedRequest, out *v1beta1.IsVolumeFormattedRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_IsVolumeFormattedRequest_To_v1beta1_IsVolumeFormattedRequest is an autogenerated conversion function.
func Convert_internal_IsVolumeFormattedRequest_To_v1beta1_IsVolumeFormattedRequest(in *internal.IsVolumeFormattedRequest, out *v1beta1.IsVolumeFormattedRequest) error {
	return autoConvert_internal_IsVolumeFormattedRequest_To_v1beta1_IsVolumeFormattedRequest(in, out)
}

func autoConvert_v1beta1_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse(in *v1beta1.IsVolumeFormattedResponse, out *internal.IsVolumeFormattedResponse) error {
	out.Formatted = in.Formatted
	return nil
}

// Convert_v1beta1_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse is an autogenerated conversion function.
func Convert_v1beta1_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse(in *v1beta1.IsVolumeFormattedResponse, out *internal.IsVolumeFormattedResponse) error {
	return autoConvert_v1beta1_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse(in, out)
}

func autoConvert_internal_IsVolumeFormattedResponse_To_v1beta1_IsVolumeFormattedResponse(in *internal.IsVolumeFormattedResponse, out *v1beta1.IsVolumeFormattedResponse) error {
	out.Formatted = in.Formatted
	return nil
}

// Convert_internal_IsVolumeFormattedResponse_To_v1beta1_IsVolumeFormattedResponse is an autogenerated conversion function.
func Convert_internal_IsVolumeFormattedResponse_To_v1beta1_IsVolumeFormattedResponse(in *internal.IsVolumeFormattedResponse, out *v1beta1.IsVolumeFormattedResponse) error {
	return autoConvert_internal_IsVolumeFormattedResponse_To_v1beta1_IsVolumeFormattedResponse(in, out)
}

// detected external conversion function
// Convert_v1beta1_ListVolumesOnDiskRequest_To_internal_ListVolumesOnDiskRequest(in *v1beta1.ListVolumesOnDiskRequest, out *internal.ListVolumesOnDiskRequest) error
// skipping generation of the auto function

// detected external conversion function
// Convert_internal_ListVolumesOnDiskRequest_To_v1beta1_ListVolumesOnDiskRequest(in *internal.ListVolumesOnDiskRequest, out *v1beta1.ListVolumesOnDiskRequest) error
// skipping generation of the auto function

func autoConvert_v1beta1_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse(in *v1beta1.ListVolumesOnDiskResponse, out *internal.ListVolumesOnDiskResponse) error {
	out.VolumeIds = *(*[]string)(unsafe.Pointer(&in.VolumeIds))
	return nil
}

// Convert_v1beta1_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse is an autogenerated conversion function.
func Convert_v1beta1_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse(in *v1beta1.ListVolumesOnDiskResponse, out *internal.ListVolumesOnDiskResponse) error {
	return autoConvert_v1beta1_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse(in, out)
}

func autoConvert_internal_ListVolumesOnDiskResponse_To_v1beta1_ListVolumesOnDiskResponse(in *internal.ListVolumesOnDiskResponse, out *v1beta1.ListVolumesOnDiskResponse) error {
	out.VolumeIds = *(*[]string)(unsafe.Pointer(&in.VolumeIds))
	return nil
}

// Convert_internal_ListVolumesOnDiskResponse_To_v1beta1_ListVolumesOnDiskResponse is an autogenerated conversion function.
func Convert_internal_ListVolumesOnDiskResponse_To_v1beta1_ListVolumesOnDiskResponse(in *internal.ListVolumesOnDiskResponse, out *v1beta1.ListVolumesOnDiskResponse) error {
	return autoConvert_internal_ListVolumesOnDiskResponse_To_v1beta1_ListVolumesOnDiskResponse(in, out)
}

// detected external conversion function
// Convert_v1beta1_MountVolumeRequest_To_internal_MountVolumeRequest(in *v1beta1.MountVolumeRequest, out *internal.MountVolumeRequest) error
// skipping generation of the auto function

// detected external conversion function
// Convert_internal_MountVolumeRequest_To_v1beta1_MountVolumeRequest(in *internal.MountVolumeRequest, out *v1beta1.MountVolumeRequest) error
// skipping generation of the auto function

func autoConvert_v1beta1_MountVolumeResponse_To_internal_MountVolumeResponse(in *v1beta1.MountVolumeResponse, out *internal.MountVolumeResponse) error {
	return nil
}

// Convert_v1beta1_MountVolumeResponse_To_internal_MountVolumeResponse is an autogenerated conversion function.
func Convert_v1beta1_MountVolumeResponse_To_internal_MountVolumeResponse(in *v1beta1.MountVolumeResponse, out *internal.MountVolumeResponse) error {
	return autoConvert_v1beta1_MountVolumeResponse_To_internal_MountVolumeResponse(in, out)
}

func autoConvert_internal_MountVolumeResponse_To_v1beta1_MountVolumeResponse(in *internal.MountVolumeResponse, out *v1beta1.MountVolumeResponse) error {
	return nil
}

// Convert_internal_MountVolumeResponse_To_v1beta1_MountVolumeResponse is an autogenerated conversion function.
func Convert_internal_MountVolumeResponse_To_v1beta1_MountVolumeResponse(in *internal.MountVolumeResponse, out *v1beta1.MountVolumeResponse) error {
	return autoConvert_internal_MountVolumeResponse_To_v1beta1_MountVolumeResponse(in, out)
}

// detected external conversion function
// Convert_v1beta1_ResizeVolumeRequest_To_internal_ResizeVolumeRequest(in *v1beta1.ResizeVolumeRequest, out *internal.ResizeVolumeRequest) error
// skipping generation of the auto function

// detected external conversion function
// Convert_internal_ResizeVolumeRequest_To_v1beta1_ResizeVolumeRequest(in *internal.ResizeVolumeRequest, out *v1beta1.ResizeVolumeRequest) error
// skipping generation of the auto function

func autoConvert_v1beta1_ResizeVolumeResponse_To_internal_ResizeVolumeResponse(in *v1beta1.ResizeVolumeResponse, out *internal.ResizeVolumeResponse) error {
	return nil
}

// Convert_v1beta1_ResizeVolumeResponse_To_internal_ResizeVolumeResponse is an autogenerated conversion function.
func Convert_v1beta1_ResizeVolumeResponse_To_internal_ResizeVolumeResponse(in *v1beta1.ResizeVolumeResponse, out *internal.ResizeVolumeResponse) error {
	return autoConvert_v1beta1_ResizeVolumeResponse_To_internal_ResizeVolumeResponse(in, out)
}

func autoConvert_internal_ResizeVolumeResponse_To_v1beta1_ResizeVolumeResponse(in *internal.ResizeVolumeResponse, out *v1beta1.ResizeVolumeResponse) error {
	return nil
}

// Convert_internal_ResizeVolumeResponse_To_v1beta1_ResizeVolumeResponse is an autogenerated conversion function.
func Convert_internal_ResizeVolumeResponse_To_v1beta1_ResizeVolumeResponse(in *internal.ResizeVolumeResponse, out *v1beta1.ResizeVolumeResponse) error {
	return autoConvert_internal_ResizeVolumeResponse_To_v1beta1_ResizeVolumeResponse(in, out)
}

func autoConvert_v1beta1_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest(in *v1beta1.VolumeDiskNumberRequest, out *internal.VolumeDiskNumberRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta1_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest is an autogenerated conversion function.
func Convert_v1beta1_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest(in *v1beta1.VolumeDiskNumberRequest, out *internal.VolumeDiskNumberRequest) error {
	return autoConvert_v1beta1_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest(in, out)
}

func autoConvert_internal_VolumeDiskNumberRequest_To_v1beta1_VolumeDiskNumberRequest(in *internal.VolumeDiskNumberRequest, out *v1beta1.VolumeDiskNumberRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_VolumeDiskNumberRequest_To_v1beta1_VolumeDiskNumberRequest is an autogenerated conversion function.
func Convert_internal_VolumeDiskNumberRequest_To_v1beta1_VolumeDiskNumberRequest(in *internal.VolumeDiskNumberRequest, out *v1beta1.VolumeDiskNumberRequest) error {
	return autoConvert_internal_VolumeDiskNumberRequest_To_v1beta1_VolumeDiskNumberRequest(in, out)
}

func autoConvert_v1beta1_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse(in *v1beta1.VolumeDiskNumberResponse, out *internal.VolumeDiskNumberResponse) error {
	out.DiskNumber = in.DiskNumber
	return nil
}

// Convert_v1beta1_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse is an autogenerated conversion function.
func Convert_v1beta1_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse(in *v1beta1.VolumeDiskNumberResponse, out *internal.VolumeDiskNumberResponse) error {
	return autoConvert_v1beta1_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse(in, out)
}

func autoConvert_internal_VolumeDiskNumberResponse_To_v1beta1_VolumeDiskNumberResponse(in *internal.VolumeDiskNumberResponse, out *v1beta1.VolumeDiskNumberResponse) error {
	out.DiskNumber = in.DiskNumber
	return nil
}

// Convert_internal_VolumeDiskNumberResponse_To_v1beta1_VolumeDiskNumberResponse is an autogenerated conversion function.
func Convert_internal_VolumeDiskNumberResponse_To_v1beta1_VolumeDiskNumberResponse(in *internal.VolumeDiskNumberResponse, out *v1beta1.VolumeDiskNumberResponse) error {
	return autoConvert_internal_VolumeDiskNumberResponse_To_v1beta1_VolumeDiskNumberResponse(in, out)
}

func autoConvert_v1beta1_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest(in *v1beta1.VolumeIDFromMountRequest, out *internal.VolumeIDFromMountRequest) error {
	out.Mount = in.Mount
	return nil
}

// Convert_v1beta1_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest is an autogenerated conversion function.
func Convert_v1beta1_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest(in *v1beta1.VolumeIDFromMountRequest, out *internal.VolumeIDFromMountRequest) error {
	return autoConvert_v1beta1_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest(in, out)
}

func autoConvert_internal_VolumeIDFromMountRequest_To_v1beta1_VolumeIDFromMountRequest(in *internal.VolumeIDFromMountRequest, out *v1beta1.VolumeIDFromMountRequest) error {
	out.Mount = in.Mount
	return nil
}

// Convert_internal_VolumeIDFromMountRequest_To_v1beta1_VolumeIDFromMountRequest is an autogenerated conversion function.
func Convert_internal_VolumeIDFromMountRequest_To_v1beta1_VolumeIDFromMountRequest(in *internal.VolumeIDFromMountRequest, out *v1beta1.VolumeIDFromMountRequest) error {
	return autoConvert_internal_VolumeIDFromMountRequest_To_v1beta1_VolumeIDFromMountRequest(in, out)
}

func autoConvert_v1beta1_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse(in *v1beta1.VolumeIDFromMountResponse, out *internal.VolumeIDFromMountResponse) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta1_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse is an autogenerated conversion function.
func Convert_v1beta1_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse(in *v1beta1.VolumeIDFromMountResponse, out *internal.VolumeIDFromMountResponse) error {
	return autoConvert_v1beta1_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse(in, out)
}

func autoConvert_internal_VolumeIDFromMountResponse_To_v1beta1_VolumeIDFromMountResponse(in *internal.VolumeIDFromMountResponse, out *v1beta1.VolumeIDFromMountResponse) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_VolumeIDFromMountResponse_To_v1beta1_VolumeIDFromMountResponse is an autogenerated conversion function.
func Convert_internal_VolumeIDFromMountResponse_To_v1beta1_VolumeIDFromMountResponse(in *internal.VolumeIDFromMountResponse, out *v1beta1.VolumeIDFromMountResponse) error {
	return autoConvert_internal_VolumeIDFromMountResponse_To_v1beta1_VolumeIDFromMountResponse(in, out)
}

func autoConvert_v1beta1_VolumeStatsRequest_To_internal_VolumeStatsRequest(in *v1beta1.VolumeStatsRequest, out *internal.VolumeStatsRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta1_VolumeStatsRequest_To_internal_VolumeStatsRequest is an autogenerated conversion function.
func Convert_v1beta1_VolumeStatsRequest_To_internal_VolumeStatsRequest(in *v1beta1.VolumeStatsRequest, out *internal.VolumeStatsRequest) error {
	return autoConvert_v1beta1_VolumeStatsRequest_To_internal_VolumeStatsRequest(in, out)
}

func autoConvert_internal_VolumeStatsRequest_To_v1beta1_VolumeStatsRequest(in *internal.VolumeStatsRequest, out *v1beta1.VolumeStatsRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_VolumeStatsRequest_To_v1beta1_VolumeStatsRequest is an autogenerated conversion function.
func Convert_internal_VolumeStatsRequest_To_v1beta1_VolumeStatsRequest(in *internal.VolumeStatsRequest, out *v1beta1.VolumeStatsRequest) error {
	return autoConvert_internal_VolumeStatsRequest_To_v1beta1_VolumeStatsRequest(in, out)
}

func autoConvert_v1beta1_VolumeStatsResponse_To_internal_VolumeStatsResponse(in *v1beta1.VolumeStatsResponse, out *internal.VolumeStatsResponse) error {
	out.VolumeSize = in.VolumeSize
	out.VolumeUsedSize = in.VolumeUsedSize
	return nil
}

// Convert_v1beta1_VolumeStatsResponse_To_internal_VolumeStatsResponse is an autogenerated conversion function.
func Convert_v1beta1_VolumeStatsResponse_To_internal_VolumeStatsResponse(in *v1beta1.VolumeStatsResponse, out *internal.VolumeStatsResponse) error {
	return autoConvert_v1beta1_VolumeStatsResponse_To_internal_VolumeStatsResponse(in, out)
}

func autoConvert_internal_VolumeStatsResponse_To_v1beta1_VolumeStatsResponse(in *internal.VolumeStatsResponse, out *v1beta1.VolumeStatsResponse) error {
	out.VolumeSize = in.VolumeSize
	out.VolumeUsedSize = in.VolumeUsedSize
	return nil
}

// Convert_internal_VolumeStatsResponse_To_v1beta1_VolumeStatsResponse is an autogenerated conversion function.
func Convert_internal_VolumeStatsResponse_To_v1beta1_VolumeStatsResponse(in *internal.VolumeStatsResponse, out *v1beta1.VolumeStatsResponse) error {
	return autoConvert_internal_VolumeStatsResponse_To_v1beta1_VolumeStatsResponse(in, out)
}
