package aravis

// #cgo pkg-config: aravis-0.8
// #cgo pkg-config: glib-2.0
// #include <arv.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Camera struct {
	camera *C.struct__ArvCamera
}

const (
	ACQUISITION_MODE_CONTINUOUS   = C.ARV_ACQUISITION_MODE_CONTINUOUS
	ACQUISITION_MODE_SINGLE_FRAME = C.ARV_ACQUISITION_MODE_SINGLE_FRAME
)

const (
	AUTO_OFF        = C.ARV_AUTO_OFF
	AUTO_ONCE       = C.ARV_AUTO_ONCE
	AUTO_CONTINUOUS = C.ARV_AUTO_CONTINUOUS
)

func NewCamera(name string) (Camera, error) {
	var c Camera
	var err error
	var gerror *C.struct__GError

	cs := C.CString(name)
	c.camera, err = C.arv_camera_new(cs, &gerror)
	C.free(unsafe.Pointer(cs))
	C.free(unsafe.Pointer(gerror))

	return c, err
}

func (c *Camera) CreateStream() (Stream, error) {
	var s Stream
	var err error
	var gerror *C.struct__GError

	s.stream, err = C.arv_camera_create_stream(
		c.camera,
		nil,
		nil,
		&gerror,
	)

	if s.stream == nil {
		return Stream{}, errors.New("Failed to create stream")
	}

	return s, err
}

func (c *Camera) GetDevice() (Device, error) {
	var d Device
	var err error

	d.device, err = C.arv_camera_get_device(c.camera)

	return d, err
}

func (c *Camera) GetVendorName() (string, error) {
	var gerror *C.struct__GError
	name, error := C.arv_camera_get_vendor_name(c.camera, &gerror)
	return C.GoString(name), error
}

func (c *Camera) GetModelName() (string, error) {
	var gerror *C.struct__GError
	name, error := C.arv_camera_get_model_name(c.camera, &gerror)
	return C.GoString(name), error
}

func (c *Camera) GetDeviceId() (string, error) {
	var gerror *C.struct__GError
	id, error := C.arv_camera_get_device_id(c.camera, &gerror)
	return C.GoString(id), error
}

func (c *Camera) GetSensorSize() (int, int, error) {
	var width, height int
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_sensor_size(
		c.camera,
		(*C.gint)(unsafe.Pointer(&width)),
		(*C.gint)(unsafe.Pointer(&height)),
		&gerror,
	)
	return int(width), int(height), err
}

func (c *Camera) SetRegion(x, y, width, height int) {
	var gerror *C.struct__GError
	C.arv_camera_set_region(c.camera,
		C.gint(x),
		C.gint(y),
		C.gint(width),
		C.gint(height),
		&gerror,
	)
}

func (c *Camera) GetRegion() (int, int, int, int, error) {
	var x, y, width, height int
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_region(
		c.camera,
		(*C.gint)(unsafe.Pointer(&x)),
		(*C.gint)(unsafe.Pointer(&y)),
		(*C.gint)(unsafe.Pointer(&width)),
		(*C.gint)(unsafe.Pointer(&height)),
		&gerror,
	)
	return int(x), int(y), int(width), int(height), err
}

func (c *Camera) GetHeightBounds() (int, int, error) {
	var min, max int
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_height_bounds(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gerror,
	)
	return int(min), int(max), err
}

func (c *Camera) GetWidthBounds() (int, int, error) {
	var min, max int
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_width_bounds(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gerror,
	)
	return int(min), int(max), err
}

func (c *Camera) SetBinning() {
	// TODO
}

func (c *Camera) GetBinning() (int, int, error) {
	var min, max int
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_binning(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gerror,
	)
	return int(min), int(max), err
}

func (c *Camera) SetPixelFormat() {
	// TODO
}

func (c *Camera) GetPixelFormat() {
	// TODO
}

func (c *Camera) GetPixelFormatAsString() {
	// TODO
}

func (c *Camera) SetPixelFormatFromString() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormats() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormatsAsDisplayNames() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormatsAsStrings() {
	// TODO
}

func (c *Camera) StartAcquisition() {
	var gerror *C.struct__GError
	C.arv_camera_start_acquisition(c.camera, &gerror)
}

func (c *Camera) StopAcquisition() {
	var gerror *C.struct__GError
	C.arv_camera_stop_acquisition(c.camera, &gerror)
}

func (c *Camera) AbortAcquisition() {
	var gerror *C.struct__GError
	C.arv_camera_abort_acquisition(c.camera, &gerror)
}

func (c *Camera) SetAcquisitionMode(mode int) {
	var gerror *C.struct__GError
	C.arv_camera_set_acquisition_mode(c.camera, C.ArvAcquisitionMode(mode), &gerror)
}

func (c *Camera) SetFrameRate(frameRate float64) {
	var gerror *C.struct__GError
	C.arv_camera_set_frame_rate(c.camera, C.double(frameRate), &gerror)
}

func (c *Camera) GetFrameRate() (float64, error) {
	var gerror *C.struct__GError
	fr, err := C.arv_camera_get_frame_rate(c.camera, &gerror)
	return float64(fr), err
}

func (c *Camera) GetFrameRateBounds() (float64, float64, error) {
	var min, max float64
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_frame_rate_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		&gerror,
	)
	return float64(min), float64(max), err
}

func (c *Camera) SetTrigger(source string) {
	csource := C.CString(source)
	var gerror *C.struct__GError
	C.arv_camera_set_trigger(c.camera, csource, &gerror)
	C.free(unsafe.Pointer(csource))
}

func (c *Camera) SetTriggerSource(source string) {
	csource := C.CString(source)
	var gerror *C.struct__GError
	C.arv_camera_set_trigger_source(c.camera, csource, &gerror)
	C.free(unsafe.Pointer(csource))
}

func (c *Camera) GetTriggerSource() (string, error) {
	var gerror *C.struct__GError
	csource, err := C.arv_camera_get_trigger_source(c.camera, &gerror)
	return C.GoString(csource), err
}

func (c *Camera) SoftwareTrigger() {
	var gerror *C.struct__GError
	C.arv_camera_software_trigger(c.camera, &gerror)
}

func (c *Camera) IsExposureTimeAvailable() (bool, error) {
	var gerror *C.struct__GError
	gboolean, err := C.arv_camera_is_exposure_time_available(c.camera, &gerror)
	return toBool(gboolean), err
}

func (c *Camera) IsExposureAutoAvailable() (bool, error) {

	var gerror *C.struct__GError
	gboolean, err := C.arv_camera_is_exposure_auto_available(c.camera, &gerror)
	return toBool(gboolean), err
}

func (c *Camera) SetExposureTime(time float64) {
	var gerror *C.struct__GError
	C.arv_camera_set_exposure_time(c.camera, C.double(time), &gerror)
}

func (c *Camera) GetExposureTime() (float64, error) {
	var gerror *C.struct__GError
	cdouble, err := C.arv_camera_get_exposure_time(c.camera, &gerror)
	return float64(cdouble), err
}

func (c *Camera) GetExposureTimeBounds() {
	// TODO
}

func (c *Camera) SetExposureTimeAuto(mode int) {
	var gerror *C.struct__GError
	C.arv_camera_set_exposure_time_auto(c.camera, C.ArvAuto(mode), &gerror)
}

func (c *Camera) GetExposureTimeAuto() {
	// TODO
}

func (c *Camera) SetGain(gain float64) {
	var gerror *C.struct__GError
	C.arv_camera_set_gain(c.camera, C.double(gain), &gerror)
}

func (c *Camera) GetGain() (float64, error) {
	var gerror *C.struct__GError
	cgain, err := C.arv_camera_get_gain(c.camera, &gerror)
	return float64(cgain), err
}

func (c *Camera) GetGainBounds() (float64, float64, error) {
	var min, max float64
	var gerror *C.struct__GError
	_, err := C.arv_camera_get_gain_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		&gerror,
	)
	return float64(min), float64(max), err
}

func (c *Camera) SetGainAuto() {
	// TODO
}

func (c *Camera) SetString(key string, value string) error {
	ckey := C.CString(key)
	cvalue := C.CString(value)
	var gerror *C.struct__GError
	C.arv_camera_set_string(
		c.camera,
		ckey,
		cvalue,
		&gerror,
	)
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cvalue))

	return nil
}

func (c *Camera) GetString(key string) (string, error) {
	ckey := C.CString(key)
	var gerror *C.struct__GError
	cvalue, err := C.arv_camera_get_string(c.camera, ckey, &gerror)
	C.free(unsafe.Pointer(ckey))
	return C.GoString(cvalue), err
}

func (c *Camera) GetPayloadSize() (uint, error) {
	var gerror *C.struct__GError
	csize, err := C.arv_camera_get_payload(c.camera, &gerror)
	return uint(csize), err
}

func (c *Camera) IsGVDevice() (bool, error) {
	cbool, err := C.arv_camera_is_gv_device(c.camera)
	return toBool(cbool), err
}

func (c *Camera) GVGetNumStreamChannels() (int, error) {
	var gerror *C.struct__GError
	cint, err := C.arv_camera_gv_get_n_stream_channels(c.camera, &gerror)
	return int(cint), err
}

func (c *Camera) GVSelectStreamChannels(id int) {
	var gerror *C.struct__GError
	C.arv_camera_gv_select_stream_channel(c.camera, C.gint(id), &gerror)
}

func (c *Camera) GVGetCurrentStreamChannel() (int, error) {
	var gerror *C.struct__GError
	cint, err := C.arv_camera_gv_get_current_stream_channel(c.camera, &gerror)
	return int(cint), err
}

func (c *Camera) GVGetPacketDelay() (int64, error) {
	var gerror *C.struct__GError
	cint64, err := C.arv_camera_gv_get_packet_delay(c.camera, &gerror)
	return int64(cint64), err
}

func (c *Camera) GVSetPacketDelay(delay int64) {
	var gerror *C.struct__GError
	C.arv_camera_gv_set_packet_delay(c.camera, C.gint64(delay), &gerror)
}

func (c *Camera) GVGetPacketSize() (int, error) {
	var gerror *C.struct__GError
	csize, err := C.arv_camera_gv_get_packet_size(c.camera, &gerror)
	return int(csize), err
}

func (c *Camera) GVSetPacketSize(size int) {
	var gerror *C.struct__GError
	C.arv_camera_gv_set_packet_size(c.camera, C.gint(size), &gerror)
}

func (c *Camera) GetChunkMode() (bool, error) {
	var gerror *C.struct__GError
	mode, err := C.arv_camera_get_chunk_mode(c.camera, &gerror)
	return toBool(mode), err
}

func (c *Camera) Close() {
	C.g_object_unref(C.gpointer(c.camera))
}
