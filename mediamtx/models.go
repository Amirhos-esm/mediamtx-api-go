package mediamtx

import (
	"fmt"

	"gopkg.in/yaml.v2"
)


type GlobalConf struct {
	Api                    *bool                         `json:"api,omitempty"`
	ApiAddress             *string                       `json:"apiAddress,omitempty"`
	ApiAllowOrigin         *string                       `json:"apiAllowOrigin,omitempty"`
	ApiEncryption          *bool                         `json:"apiEncryption,omitempty"`
	ApiServerCert          *string                       `json:"apiServerCert,omitempty"`
	ApiServerKey           *string                       `json:"apiServerKey,omitempty"`
	ApiTrustedProxies      *[]string                     `json:"apiTrustedProxies,omitempty"`
	AuthHTTPAddress        *string                       `json:"authHTTPAddress,omitempty"`
	AuthHTTPExclude        *[]AuthInternalUserPermission `json:"authHTTPExclude,omitempty"`
	AuthInternalUsers      *[]AuthInternalUser           `json:"authInternalUsers,omitempty"`
	AuthJWTClaimKey        *string                       `json:"authJWTClaimKey,omitempty"`
	AuthJWTJWKS            *string                       `json:"authJWTJWKS,omitempty"`
	AuthMethod             *string                       `json:"authMethod,omitempty"`
	Hls                    *bool                         `json:"hls,omitempty"`
	HlsAddress             *string                       `json:"hlsAddress,omitempty"`
	HlsAllowOrigin         *string                       `json:"hlsAllowOrigin,omitempty"`
	HlsAlwaysRemux         *bool                         `json:"hlsAlwaysRemux,omitempty"`
	HlsDirectory           *string                       `json:"hlsDirectory,omitempty"`
	HlsEncryption          *bool                         `json:"hlsEncryption,omitempty"`
	HlsMuxerCloseAfter     *string                       `json:"hlsMuxerCloseAfter,omitempty"`
	HlsPartDuration        *string                       `json:"hlsPartDuration,omitempty"`
	HlsSegmentCount        *int                          `json:"hlsSegmentCount,omitempty"`
	HlsSegmentDuration     *string                       `json:"hlsSegmentDuration,omitempty"`
	HlsSegmentMaxSize      *string                       `json:"hlsSegmentMaxSize,omitempty"`
	HlsServerCert          *string                       `json:"hlsServerCert,omitempty"`
	HlsServerKey           *string                       `json:"hlsServerKey,omitempty"`
	HlsTrustedProxies      *[]string                     `json:"hlsTrustedProxies,omitempty"`
	HlsVariant             *string                       `json:"hlsVariant,omitempty"`
	LogDestinations        *[]string                     `json:"logDestinations,omitempty"`
	LogFile                *string                       `json:"logFile,omitempty"`
	LogLevel               *string                       `json:"logLevel,omitempty"`
	Metrics                *bool                         `json:"metrics,omitempty"`
	MetricsAddress         *string                       `json:"metricsAddress,omitempty"`
	MetricsAllowOrigin     *string                       `json:"metricsAllowOrigin,omitempty"`
	MetricsEncryption      *bool                         `json:"metricsEncryption,omitempty"`
	MetricsServerCert      *string                       `json:"metricsServerCert,omitempty"`
	MetricsServerKey       *string                       `json:"metricsServerKey,omitempty"`
	MetricsTrustedProxies  *[]string                     `json:"metricsTrustedProxies,omitempty"`
	MulticastIPRange       *string                       `json:"multicastIPRange,omitempty"`
	MulticastRTCPPort      *int                          `json:"multicastRTCPPort,omitempty"`
	MulticastRTPPort       *int                          `json:"multicastRTPPort,omitempty"`
	Playback               *bool                         `json:"playback,omitempty"`
	PlaybackAddress        *string                       `json:"playbackAddress,omitempty"`
	PlaybackAllowOrigin    *string                       `json:"playbackAllowOrigin,omitempty"`
	PlaybackEncryption     *bool                         `json:"playbackEncryption,omitempty"`
	PlaybackServerCert     *string                       `json:"playbackServerCert,omitempty"`
	PlaybackServerKey      *string                       `json:"playbackServerKey,omitempty"`
	PlaybackTrustedProxies *[]string                     `json:"playbackTrustedProxies,omitempty"`
	Pprof                  *bool                         `json:"pprof,omitempty"`
	PprofAddress           *string                       `json:"pprofAddress,omitempty"`
	PprofAllowOrigin       *string                       `json:"pprofAllowOrigin,omitempty"`
	PprofEncryption        *bool                         `json:"pprofEncryption,omitempty"`
	PprofServerCert        *string                       `json:"pprofServerCert,omitempty"`
	PprofServerKey         *string                       `json:"pprofServerKey,omitempty"`
	PprofTrustedProxies    *[]string                     `json:"pprofTrustedProxies,omitempty"`
	ReadTimeout            *string                       `json:"readTimeout,omitempty"`
	RtcpAddress            *string                       `json:"rtcpAddress,omitempty"`
	Rtmp                   *bool                         `json:"rtmp,omitempty"`
	RtmpAddress            *string                       `json:"rtmpAddress,omitempty"`
	RtmpEncryption         *string                       `json:"rtmpEncryption,omitempty"`
	RtmpServerCert         *string                       `json:"rtmpServerCert,omitempty"`
	RtmpServerKey          *string                       `json:"rtmpServerKey,omitempty"`
	RtmpsAddress           *string                       `json:"rtmpsAddress,omitempty"`
	RtpAddress             *string                       `json:"rtpAddress,omitempty"`
	Rtsp                   *bool                         `json:"rtsp,omitempty"`
	RtspAddress            *string                       `json:"rtspAddress,omitempty"`
	RtspAuthMethods        *[]string                     `json:"rtspAuthMethods,omitempty"`
	RtspEncryption         *string                       `json:"rtspEncryption,omitempty"`
	RtspServerCert         *string                       `json:"rtspServerCert,omitempty"`
	RtspServerKey          *string                       `json:"rtspServerKey,omitempty"`
	RtspTransports         *[]string                     `json:"rtspTransports,omitempty"`
	RtspsAddress           *string                       `json:"rtspsAddress,omitempty"`
	RunOnConnect           *string                       `json:"runOnConnect,omitempty"`
	RunOnConnectRestart    *bool                         `json:"runOnConnectRestart,omitempty"`
	RunOnDisconnect        *string                       `json:"runOnDisconnect,omitempty"`
	Srt                    *bool                         `json:"srt,omitempty"`
	SrtAddress             *string                       `json:"srtAddress,omitempty"`
	UdpMaxPayloadSize      *int                          `json:"udpMaxPayloadSize,omitempty"`
	Webrtc                 *bool                         `json:"webrtc,omitempty"`
	WebrtcAdditionalHosts  *[]string                     `json:"webrtcAdditionalHosts,omitempty"`
	WebrtcAddress          *string                       `json:"webrtcAddress,omitempty"`
	WebrtcAllowOrigin      *string                       `json:"webrtcAllowOrigin,omitempty"`
	WebrtcEncryption       *bool                         `json:"webrtcEncryption,omitempty"`
	WebrtcHandshakeTimeout *string                       `json:"webrtcHandshakeTimeout,omitempty"`
	WebrtcICEServers2      *[]struct {
		ClientOnly *bool   `json:"clientOnly,omitempty"`
		Password   *string `json:"password,omitempty"`
		Url        *string `json:"url,omitempty"`
		Username   *string `json:"username,omitempty"`
	} `json:"webrtcICEServers2,omitempty"`
	WebrtcIPsFromInterfaces     *bool     `json:"webrtcIPsFromInterfaces,omitempty"`
	WebrtcIPsFromInterfacesList *[]string `json:"webrtcIPsFromInterfacesList,omitempty"`
	WebrtcLocalTCPAddress       *string   `json:"webrtcLocalTCPAddress,omitempty"`
	WebrtcLocalUDPAddress       *string   `json:"webrtcLocalUDPAddress,omitempty"`
	WebrtcServerCert            *string   `json:"webrtcServerCert,omitempty"`
	WebrtcServerKey             *string   `json:"webrtcServerKey,omitempty"`
	WebrtcTrackGatherTimeout    *string   `json:"webrtcTrackGatherTimeout,omitempty"`
	WebrtcTrustedProxies        *[]string `json:"webrtcTrustedProxies,omitempty"`
	WriteQueueSize              *int      `json:"writeQueueSize,omitempty"`
	WriteTimeout                *string   `json:"writeTimeout,omitempty"`
}

// AuthInternalUser defines model for AuthInternalUser.
type AuthInternalUser struct {
	Ips         *[]string                     `json:"ips,omitempty"`
	Pass        *string                       `json:"pass,omitempty"`
	Permissions *[]AuthInternalUserPermission `json:"permissions,omitempty"`
	User        *string                       `json:"user,omitempty"`
}

// AuthInternalUserPermission defines model for AuthInternalUserPermission.
type AuthInternalUserPermission struct {
	Action *string `json:"action,omitempty"`
	Path   *string `json:"path,omitempty"`
}

// PathConf defines model for PathConf.
type PathConf struct {
	Fallback                   *string    `json:"fallback,omitempty"`
	MaxReaders                 *int       `json:"maxReaders,omitempty"`
	Name                       *string    `json:"name,omitempty"`
	OverridePublisher          *bool      `json:"overridePublisher,omitempty"`
	Record                     *bool      `json:"record,omitempty"`
	RecordDeleteAfter          *string    `json:"recordDeleteAfter,omitempty"`
	RecordFormat               *string    `json:"recordFormat,omitempty"`
	RecordPartDuration         *string    `json:"recordPartDuration,omitempty"`
	RecordPath                 *string    `json:"recordPath,omitempty"`
	RecordSegmentDuration      *string    `json:"recordSegmentDuration,omitempty"`
	RpiCameraAWB               *string    `json:"rpiCameraAWB,omitempty"`
	RpiCameraAWBGains          *[]float32 `json:"rpiCameraAWBGains,omitempty"`
	RpiCameraAfMode            *string    `json:"rpiCameraAfMode,omitempty"`
	RpiCameraAfRange           *string    `json:"rpiCameraAfRange,omitempty"`
	RpiCameraAfSpeed           *string    `json:"rpiCameraAfSpeed,omitempty"`
	RpiCameraAfWindow          *string    `json:"rpiCameraAfWindow,omitempty"`
	RpiCameraBitrate           *int       `json:"rpiCameraBitrate,omitempty"`
	RpiCameraBrightness        *float32   `json:"rpiCameraBrightness,omitempty"`
	RpiCameraCamID             *int       `json:"rpiCameraCamID,omitempty"`
	RpiCameraCodec             *string    `json:"rpiCameraCodec,omitempty"`
	RpiCameraContrast          *float32   `json:"rpiCameraContrast,omitempty"`
	RpiCameraDenoise           *string    `json:"rpiCameraDenoise,omitempty"`
	RpiCameraEV                *float32   `json:"rpiCameraEV,omitempty"`
	RpiCameraExposure          *string    `json:"rpiCameraExposure,omitempty"`
	RpiCameraFPS               *float32   `json:"rpiCameraFPS,omitempty"`
	RpiCameraFlickerPeriod     *int       `json:"rpiCameraFlickerPeriod,omitempty"`
	RpiCameraGain              *float32   `json:"rpiCameraGain,omitempty"`
	RpiCameraHDR               *bool      `json:"rpiCameraHDR,omitempty"`
	RpiCameraHFlip             *bool      `json:"rpiCameraHFlip,omitempty"`
	RpiCameraHeight            *int       `json:"rpiCameraHeight,omitempty"`
	RpiCameraIDRPeriod         *int       `json:"rpiCameraIDRPeriod,omitempty"`
	RpiCameraLensPosition      *float32   `json:"rpiCameraLensPosition,omitempty"`
	RpiCameraLevel             *string    `json:"rpiCameraLevel,omitempty"`
	RpiCameraMetering          *string    `json:"rpiCameraMetering,omitempty"`
	RpiCameraMode              *string    `json:"rpiCameraMode,omitempty"`
	RpiCameraProfile           *string    `json:"rpiCameraProfile,omitempty"`
	RpiCameraROI               *string    `json:"rpiCameraROI,omitempty"`
	RpiCameraSaturation        *float32   `json:"rpiCameraSaturation,omitempty"`
	RpiCameraSharpness         *float32   `json:"rpiCameraSharpness,omitempty"`
	RpiCameraShutter           *int       `json:"rpiCameraShutter,omitempty"`
	RpiCameraTextOverlay       *string    `json:"rpiCameraTextOverlay,omitempty"`
	RpiCameraTextOverlayEnable *bool      `json:"rpiCameraTextOverlayEnable,omitempty"`
	RpiCameraTuningFile        *string    `json:"rpiCameraTuningFile,omitempty"`
	RpiCameraVFlip             *bool      `json:"rpiCameraVFlip,omitempty"`
	RpiCameraWidth             *int       `json:"rpiCameraWidth,omitempty"`
	RtspAnyPort                *bool      `json:"rtspAnyPort,omitempty"`
	RtspRangeStart             *string    `json:"rtspRangeStart,omitempty"`
	RtspRangeType              *string    `json:"rtspRangeType,omitempty"`
	RtspTransport              *string    `json:"rtspTransport,omitempty"`
	RunOnDemand                *string    `json:"runOnDemand,omitempty"`
	RunOnDemandCloseAfter      *string    `json:"runOnDemandCloseAfter,omitempty"`
	RunOnDemandRestart         *bool      `json:"runOnDemandRestart,omitempty"`
	RunOnDemandStartTimeout    *string    `json:"runOnDemandStartTimeout,omitempty"`
	RunOnInit                  *string    `json:"runOnInit,omitempty"`
	RunOnInitRestart           *bool      `json:"runOnInitRestart,omitempty"`
	RunOnNotReady              *string    `json:"runOnNotReady,omitempty"`
	RunOnRead                  *string    `json:"runOnRead,omitempty"`
	RunOnReadRestart           *bool      `json:"runOnReadRestart,omitempty"`
	RunOnReady                 *string    `json:"runOnReady,omitempty"`
	RunOnReadyRestart          *bool      `json:"runOnReadyRestart,omitempty"`
	RunOnRecordSegmentComplete *string    `json:"runOnRecordSegmentComplete,omitempty"`
	RunOnRecordSegmentCreate   *string    `json:"runOnRecordSegmentCreate,omitempty"`
	RunOnUnDemand              *string    `json:"runOnUnDemand,omitempty"`
	RunOnUnread                *string    `json:"runOnUnread,omitempty"`
	Source                     *string    `json:"source,omitempty"`
	SourceFingerprint          *string    `json:"sourceFingerprint,omitempty"`
	SourceOnDemand             *bool      `json:"sourceOnDemand,omitempty"`
	SourceOnDemandCloseAfter   *string    `json:"sourceOnDemandCloseAfter,omitempty"`
	SourceOnDemandStartTimeout *string    `json:"sourceOnDemandStartTimeout,omitempty"`
	SourceRedirect             *string    `json:"sourceRedirect,omitempty"`
	SrtPublishPassphrase       *string    `json:"srtPublishPassphrase,omitempty"`
	SrtReadPassphrase          *string    `json:"srtReadPassphrase,omitempty"`
}

// PathReader defines model for PathReader.
type PathReader struct {
	Id   *string         `json:"id,omitempty"`
	Type *PathReaderType `json:"type,omitempty"`
}

// PathReaderType defines model for PathReader.Type.
type PathReaderType string

// PathSource defines model for PathSource.
type PathSource struct {
	Id   *string         `json:"id,omitempty"`
	Type *PathSourceType `json:"type,omitempty"`
}
type PathSourceType string


type AllPathConfiguration struct {
	PageCount int        `json:"pageCount"`
	ItemCount int        `json:"itemCount"`
	Items     []PathConf `json:"items"`
}



func (g GlobalConf) String() string {
	out, err := yaml.Marshal(&g)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(out)
}

func (a AuthInternalUser) String() string {
	out, err := yaml.Marshal(&a)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(out)
}

func (a AuthInternalUserPermission) String() string {
	out, err := yaml.Marshal(&a)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(out)
}

func (p PathConf) String() string {
	out, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(out)
}

func (p PathReader) String() string {
	out, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(out)
}

func (p PathSource) String() string {
	out, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return string(out)
}

