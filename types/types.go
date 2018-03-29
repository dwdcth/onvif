package types

//todo посмотреть все Extensions
//todo что делать ч xs:any = Any

type DeviceEntity struct {
	Token ReferenceToken //required attribute
}

type ReferenceToken struct {
	Token string
}

type Name struct {
	Name string
}

type IntRectangle struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
	Width int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type IntRectangleRange struct {
	XRange IntRange
	YRange IntRange
	WidthRange IntRange
	HeightRange IntRange
}

type IntRange struct {
	Min int
	Max int
}

type FloatRange struct {
	Min float64
	Max float64
}

type OSDConfiguration struct {
	DeviceEntity
	VideoSourceConfigurationToken OSDReference
	Type OSDType
	Position OSDPosConfiguration
	TextString OSDTextConfiguration
	Image OSDImgConfiguration
	Extension OSDConfigurationExtension
}

type OSDType struct {
	Type string
}

type OSDPosConfiguration struct {
	Type string
	Pos Vector
	Extension OSDPosConfigurationExtension
}

type Vector struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type OSDPosConfigurationExtension struct {
	Any string //while
}

type OSDReference struct {
	ReferenceToken
}

type OSDTextConfiguration struct {
	Type string
	DateFormat string
	TimeFormat string
	FontSize int
	FontColor OSDColor
	BackgroundColor OSDColor
	PlainText string
	Extension OSDTextConfigurationExtension
}

type OSDColor struct {
	Color Color
	Transparent int
}

type Color struct {
	X float64 `xml:"X,attr"`
	Y float64 `xml:"Y,attr"`
	Z float64 `xml:"Z,attr"`
	Colorspace anyURI `xml:"Colorspace,attr"`
}


type OSDTextConfigurationExtension struct {
	Any string
}

type OSDImgConfiguration struct {
	ImgPath anyURI
	Extension OSDImgConfigurationExtension
}

type OSDImgConfigurationExtension struct {
	Any string
}

type OSDConfigurationExtension struct {
	Any string //While
}

type Capabilities struct {
	SnapshotUri 		bool `xml:"SnapshotUri,attr"`
	Rotation 			bool `xml:"Rotation,attr"`
	VideoSourceMode 	bool `xml:"VideoSourceMode,attr"`
	OSD 				bool `xml:"OSD,attr"`
	TemporaryOSDText 	bool `xml:"TemporaryOSDText,attr"`
	EXICompression 		bool `xml:"EXICompression,attr"`
	ProfileCapabilities ProfileCapabilities
	StreamingCapabilities StreamingCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int `xml:"MaximumNumberOfProfiles,attr"`
}

type StreamingCapabilities struct {
	RTPMulticast 		bool `xml:"RTPMulticast,attr"`
	RTP_TCP 			bool `xml:"RTP_TCP,attr"`
	RTP_RTSP_TCP 		bool `xml:"RTP_RTSP_TCP,attr"`
	NonAggregateControl bool `xml:"NonAggregateControl,attr"`
	NoRTSPStreaming 	bool `xml:"NoRTSPStreaming,attr"`
}

type VideoSource struct {
	DeviceEntity
	Framerate 	float64
	Resolution 	VideoResolution
	Imaging 	ImagingSettings
	Extension 	VideoSourceExtension
}

type VideoResolution struct {
	Width 	int
	Height 	int
}

type ImagingSettings struct {
	BacklightCompensation 	BacklightCompensation
	Brightness 				float64
	ColorSaturation 		float64
	Contrast 				float64
	Exposure 				Exposure
	Focus 					FocusConfiguration
	IrCutFilter 			IrCutFilterMode
	Sharpness 				float64
	WideDynamicRange 		WideDynamicRange
	WhiteBalance 			WhiteBalance
	Extension 				ImagingSettingsExtension
}

type BacklightCompensation struct {
	Mode 	BacklightCompensationMode
	Level 	float64
}

type BacklightCompensationMode struct {
	BacklightCompensation string
}

type Exposure struct {
	Mode 			ExposureMode
	Priority 		ExposurePriority
	Window 			Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain 		float64
	MaxGain 		float64
	MinIris 		float64
	MaxIris 		float64
	ExposureTime 	float64
	Gain 			float64
	Iris 			float64
}

type ExposureMode struct {
	Mode string
}

type ExposurePriority struct {
	Priority string
}

type Rectangle struct {
	Bottom 	float64 `xml:"bottom,attr"`
	Top 	float64 `xml:"top,attr"`
	Right 	float64 `xml:"right,attr"`
	Left 	float64 `xml:"left,attr"`
}

type FocusConfiguration struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed float64
	NearLimit float64
	FarLimit float64
}

type AutoFocusMode struct {
	FocusMode string
}

type IrCutFilterMode struct {
	FilterMode string
}

type WideDynamicRange struct {
	Mode 	WideDynamicMode
	Level 	float64
}

type WideDynamicMode struct {
	DynamicMode string
}

type WhiteBalance struct {
	Mode WhiteBalanceMode
	CrGain float64
	CbGain float64
}

type WhiteBalanceMode struct {
	BalanceMode string
}

type ImagingSettingsExtension struct {
	Any string
}

type VideoSourceExtension struct {
	Imaging ImagingSettings20
	Extension VideoSourceExtension2
}

type ImagingSettings20 struct {
	BacklightCompensation BacklightCompensation20
	Brightness float64
	ColorSaturation float64
	Contrast float64
	Exposure Exposure20
	Focus FocusConfiguration20
	IrCutFilter IrCutFilterMode
	Sharpness float64
	WideDynamicRange WideDynamicRange20
	WhiteBalance WhiteBalance20
	Extension ImagingSettingsExtension20
}

type BacklightCompensation20 struct {
	Mode BacklightCompensationMode
	Level float64
}

type Exposure20 struct {
	Mode ExposureMode
	Priority ExposurePriority
	Window Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain float64
	MaxGain float64
	MinIris float64
	MaxIris float64
	ExposureTime float64
	Gain float64
	Iris float64
}

type FocusConfiguration20 struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed float64
	NearLimit float64
	FarLimit float64
	Extension FocusConfiguration20Extension
}

type FocusConfiguration20Extension struct {
	Any string
}

type WideDynamicRange20 struct {
	Mode WideDynamicMode
	Level float64
}

type WhiteBalance20 struct {
	Mode WhiteBalanceMode
	CrGain float64
	CbGain float64
	Extension WhiteBalance20Extension
}

type WhiteBalance20Extension struct {
	Any string
}

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization
	Extension ImagingSettingsExtension202
}

type ImageStabilization struct {
	Mode ImageStabilizationMode
	Level float64
	Extension ImageStabilizationExtension
}

type ImageStabilizationMode struct {
	StabilizationMode string
}

type ImageStabilizationExtension struct {
	Any string
}

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment
	Extension ImagingSettingsExtension203
}

type IrCutFilterAutoAdjustment struct {
	BoundaryType string
	BoundaryOffset float64
	ResponseTime duration
	Extension IrCutFilterAutoAdjustmentExtension
}

type IrCutFilterAutoAdjustmentExtension struct {
	Any string
}

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation
	Defogging Defogging
	NoiseReduction NoiseReduction
	Extension ImagingSettingsExtension204
}

type ToneCompensation struct {
	Mode string
	Level float64
	Extension ToneCompensationExtension
}

type ToneCompensationExtension struct {
	Any string
}

type Defogging struct {
	Mode string
	Level float64
	Extension DefoggingExtension
}

type DefoggingExtension struct {
	Any string
}

type NoiseReduction struct {
	Level float64
}

type ImagingSettingsExtension204 struct {
	Any string
}

type VideoSourceExtension2 struct {
	Any string
}

type AudioSource struct {
	DeviceEntity
	Channels int
}

type AudioOutput struct {
	DeviceEntity
}

type Profile struct {
	Token ReferenceToken 	`xml:"token,attr"`
	Fixed bool 				`xml:"fixed,attr"`
	Name Name
	VideoSourceConfiguration VideoSourceConfiguration
	AudioSourceConfiguration AudioSourceConfiguration
	VideoEncoderConfiguration VideoEncoderConfiguration
	AudioEncoderConfiguration AudioEncoderConfiguration
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration
	PTZConfiguration PTZConfiguration
	MetadataConfiguration MetadataConfiguration
	Extension ProfileExtension
}

type VideoSourceConfiguration struct {
	ConfigurationEntity
	ViewMode string `xml:"ViewMode,attr"`
	SourceToken ReferenceToken
	Bounds IntRectangle
	Extension VideoSourceConfigurationExtension
}

type ConfigurationEntity struct {
	Token ReferenceToken `xml:"token,attr"`
	Name Name
	UseCount int
}

type VideoSourceConfigurationExtension struct {
	Rotate Rotate
	Extension VideoSourceConfigurationExtension2
}

type Rotate struct {
	Mode RotateMode
	Degree int
	Extension RotateExtension
}

type RotateMode struct {
	RotateMode string
}

type RotateExtension struct {
	Any string
}

type VideoSourceConfigurationExtension2 struct {
	LensDescription LensDescription
	SceneOrientation SceneOrientation
}

type LensDescription struct {
	FocalLength float64 `xml:"FocalLength,attr"`
	Offset LensOffset
	Projection LensProjection
	XFactor float64
}

type LensOffset struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type LensProjection struct {
	Angle float64
	Radius float64
	Transmittance float64
}

type SceneOrientation struct {
	Mode SceneOrientationMode
	Orientation string
}

type SceneOrientationMode struct {
	OrientationMode string
}

type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken ReferenceToken
}

type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding VideoEncoding
	Resolution VideoResolution
	Quality float64
	RateControl VideoRateControl
	MPEG4 Mpeg4Configuration
	H264 H264Configuration
	Multicast MulticastConfiguration
	SessionTimeout duration
}

type VideoEncoding struct {
	VideoEncoding string
}

type VideoRateControl struct {
	FrameRateLimit int
	EncodingInterval int
	BitrateLimit int
}

type Mpeg4Configuration struct {
	GovLength int
	Mpeg4Profile Mpeg4Profile
}

type Mpeg4Profile struct {
	Profile string
}

type H264Configuration struct {
	GovLength int
	H264Profile H264Profile
}

type H264Profile struct {
	Profile string
}

type MulticastConfiguration struct {
	Address IPAddress
	Port int
	TTL int
	AutoStart bool
}

type IPAddress struct {
	Type IPType
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type IPType struct {
	Type string
}

type IPv4Address struct {
	Address token
}

type IPv6Address struct {
	Address token
}

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding AudioEncoding
	Bitrate int
	SampleRate int
	Multicast MulticastConfiguration
	SessionTimeout duration
}

type AudioEncoding struct {
	Encoding string
}

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration
	RuleEngineConfiguration RuleEngineConfiguration
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule Config
	Extension AnalyticsEngineConfigurationExtension
}

type Config struct {
	Name string `xml:"Name,attr"`
	Type QName 	`xml:"Type,attr"`
	Parameters ItemList
}

type ItemList struct {
	SimpleItem
	ElementItem
	Extension ItemListExtension
}

type SimpleItem struct {
	Name 	string 			`xml:"Name,attr"`
	Value 	anySimpleType 	`xml:"Value,attr"`
}

type ElementItem struct {
	Name string `xml:"Name,attr"`
}

type ItemListExtension struct {
	Any string
}

type AnalyticsEngineConfigurationExtension struct {
	Any string
}

type RuleEngineConfiguration struct {
	Rule Config
	Extension RuleEngineConfigurationExtension
}

type RuleEngineConfigurationExtension struct {
	Any string
}

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp 		int `xml:"MoveRamp,attr"`
	PresetRamp 		int `xml:"PresetRamp,attr"`
	PresetTourRamp 	int `xml:"PresetTourRamp,attr"`
	NodeToken ReferenceToken
	DefaultAbsolutePantTiltPositionSpace anyURI
	DefaultAbsoluteZoomPositionSpace anyURI
	DefaultRelativePanTiltTranslationSpace anyURI
	DefaultRelativeZoomTranslationSpace anyURI
	DefaultContinuousPanTiltVelocitySpace anyURI
	DefaultContinuousZoomVelocitySpace anyURI
	DefaultPTZSpeed PTZSpeed
	DefaultPTZTimeout duration
	PanTiltLimits PanTiltLimits
	ZoomLimits ZoomLimits
	Extension PTZConfigurationExtension
}

type PTZSpeed struct {
	PanTilt Vector2D
	Zoom Vector1D
}

type Vector2D struct {
	X 		float64 `xml:"x,attr"`
	Y 		float64 `xml:"y,attr"`
	Space 	anyURI 	`xml:"space,attr"`
}

type Vector1D struct {
	X 		float64 `xml:"x,attr"`
	Space 	anyURI 	`xml:"space,attr"`
}

type PanTiltLimits struct {
	Range Space2DDescription
}

type Space2DDescription struct {
	URI anyURI
	XRange FloatRange
	YRange FloatRange
}

type ZoomLimits struct {
	Range Space1DDescription
}

type Space1DDescription struct {
	URI anyURI
	XRange FloatRange
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection
	Extension PTZConfigurationExtension2
}

type PTControlDirection struct {
	EFlip EFlip
	Reverse Reverse
	Extension PTControlDirectionExtension
}

type EFlip struct {
	Mode EFlipMode
}

type EFlipMode struct {
	EFlipMode string
}

type Reverse struct {
	Mode ReverseMode
}

type ReverseMode struct {
	ReverseMode string
}

type PTControlDirectionExtension struct {
	Any string
}

type PTZConfigurationExtension2 struct {
	Any string
}

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType 				string `xml:"CompressionType,attr"`
	PTZStatus 						PTZFilter
	Events 							EventSubscription
	Analytics 						bool
	Multicast 						MulticastConfiguration
	SessionTimeout 					duration
	AnalyticsEngineConfiguration 	AnalyticsEngineConfiguration
	Extension 						MetadataConfigurationExtension
}

type PTZFilter struct {
	Status bool
	Position bool
}

type EventSubscription struct {
	Filter FilterType
	SubscriptionPolicy
}

type FilterType struct {
	Any string
}

type SubscriptionPolicy struct {
	Any string
}

type MetadataConfigurationExtension struct {
	Any string
}

type ProfileExtension struct {
	AudioOutputConfiguration AudioOutputConfiguration
	AudioDecoderConfiguration AudioDecoderConfiguration
	Extension ProfileExtension2
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken ReferenceToken
	SendPrimacy anyURI
	OutputLevel int
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type ProfileExtension2 struct {
	Any string
}