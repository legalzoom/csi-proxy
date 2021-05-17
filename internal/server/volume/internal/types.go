// Defines all the structs that the server is aware of, because all
// the apis are included as the target for code generation it also
// has definition for older APIs e.g. volume/v1alpha1, volume/v1beta1, etc
// Because of this some structs are needed but are no longer used in newer APIs

package internal

type ListVolumesOnDiskRequest struct {
	// DiskId is deprecated, it's parsed into a DiskNumber in the internal server implementation
	DiskId          string
	DiskNumber      uint32
	PartitionNumber uint32
}

type ListVolumesOnDiskResponse struct {
	VolumeIds []string
}

type MountVolumeRequest struct {
	VolumeId string
	// Path is deprecated, it's copied to TargetPath in the internal server implementation
	Path       string
	TargetPath string
}

type MountVolumeResponse struct {
}

type IsVolumeFormattedRequest struct {
	VolumeId string
}

type IsVolumeFormattedResponse struct {
	Formatted bool
}

type FormatVolumeRequest struct {
	VolumeId string
}

type FormatVolumeResponse struct {
}

type WriteVolumeCacheRequest struct {
	VolumeId string
}

type WriteVolumeCacheResponse struct {
}

type UnmountVolumeRequest struct {
	VolumeId   string
	TargetPath string
}

type UnmountVolumeResponse struct {
}

type ResizeVolumeRequest struct {
	VolumeId string
	// Size is deprecated, it's copied into SizeBytes
	Size      int64
	SizeBytes int64
}

type ResizeVolumeResponse struct {
}

type GetVolumeStatsRequest struct {
	VolumeId string
}

type GetVolumeStatsResponse struct {
	TotalBytes int64
	UsedBytes  int64
}

type GetDiskNumberFromVolumeIDRequest struct {
	VolumeId string
}

type GetDiskNumberFromVolumeIDResponse struct {
	DiskNumber uint32
}

type GetVolumeIDFromTargetPathRequest struct {
	TargetPath string
}

type GetVolumeIDFromTargetPathResponse struct {
	VolumeId string
}

// These structs are deprecated but are needed to provide support for older APIs

type DismountVolumeRequest struct {
	VolumeId string
	Path     string
}
type DismountVolumeResponse struct{}
type VolumeDiskNumberRequest struct {
	VolumeId string
}
type VolumeDiskNumberResponse struct {
	DiskNumber int64
}
type VolumeIDFromMountRequest struct {
	Mount string
}
type VolumeIDFromMountResponse struct {
	VolumeId string
}
type VolumeStatsRequest struct {
	VolumeId string
}
type VolumeStatsResponse struct {
	VolumeSize     int64
	VolumeUsedSize int64
}
