package vmt

// Properties is an example material that contains all known
// material properties. This should be updated as more properties are added
// or discovered in games.
type Properties struct {
	// ShaderName is special, as it is the only property where
	// the value we want is the key itself. __SHADER_NAME__ is a custom
	// key that the parser can manually map to.
	ShaderName string `vmt:"__SHADER_NAME__"`

	// Basics
	BaseTexture string `vmt:"$basetexture"`
	SurfaceProp string `vmt:"$surfaceprop"`
	Detail      string `vmt:"$detail"`
	Model       int    `vmt:"$model"`

	// Adjustment
	BaseTextureTransform string `vmt:"$basetexturetransform"`
	Color                string `vmt:"$color"`
	SeamlessScale        string `vmt:"$seamless_scale"`

	// Transparency
	Alpha                float32 `vmt:"$alpha"`
	AlphaTest            string  `vmt:"$alphatest"`
	BlendModulateTexture string  `vmt:"$blendmodulatetexture"`
	DistanceAlpha        string  `vmt:"$distancealpha"`
	NoCull               string  `vmt:"$nocull"`
	Translucent          int     `vmt:"$translucent"`

	// Lighting
	Bumpmap          string `vmt:"$bumpmap"`
	SSBump           string `vmt:"$ssbump"`
	ParallaxMap      string `vmt:"$parallaxmap"`
	Selfillum        int    `vmt:"$selfillum"`
	LightwarpTexture string `vmt:"$lightwarptexture"`
	HalfLambert      string `vmt:"$halflambert"`
	AmbientOcclusion string `vmt:"$ambientocclusion"`
	Rimlight         string `vmt:"$rimlight"`

	// Reflection
	Reflectivity string `vmt:"$reflectivity"`
	Phong        string `vmt:"$phong"`
	Envmap       string `vmt:"$envmap"`

	// Optimization
	IgnoreZ int `vmt:"$ignorez"`

	// Texture Organization
	CompileKeywords      string `vmt:"%keywords"`
	CompileNoToolTexture int    `vmt:"%notooltexture"`
	CompileToolTexture   string `vmt:"%tooltexture"`

	// Effect
	TreeSway string `vmt:"$treeSway"`
	NoFog    int    `vmt:"$nofog"`

	// Misc
	CompileClip        int `vmt:"%compileclip"`
	CompileDetail      int `vmt:"%compiledetail"`
	CompileLadder      int `vmt:"%compileladder"`
	CompileNoDraw      int `vmt:"%compilenodraw"`
	CompileNoLight     int `vmt:"%compilenolight"`
	CompileNonSolid    int `vmt:"%compilenonsolid"`
	CompilePassBullets int `vmt:"%compilepassbullets"`
	CompileSkip        int `vmt:"%compileskip"`
	CompileSky         int `vmt:"%compilesky"`

	AllowDiffuseModulation string `vmt:"$allowdiffusemodulation"`

	BasetextureNoEnvmap            string `vmt:"$basetexturenoenvmap"`
	BlendTintByBaseAlpha           string `vmt:"$blendtintbybasealpha"`
	BumpmapBaseTexture2WithBumpmap string `vmt:"$bumpbasetexture2withbumpmap"`
	BumpFrame                      string `vmt:"$bumpframe"`
	Bumpoffset                     string `vmt:"$bumpoffset"`
	BumpScale                      string `vmt:"$bumpscale"`
	BumpTransform                  string `vmt:"$bumptransform"`

	DisplacementMap string `vmt:"$displacementmap"`

	EmissiveBlend       string `vmt:"$emissiveblend"`
	EnvmapContrast      string `vmt:"$envmapcontrast"`
	EnvmapFrame         string `vmt:"$envmapframe"`
	EnvmapMask          string `vmt:"$envmapmask"`
	EnvmapMaskFrame     string `vmt:"$envmapmaskframe"`
	EnvmapMaskScale     string `vmt:"$envmapmaskscale"`
	EnvmpaMaskTransform string `vmt:"$envmapmasktransform"`
	EnvmapSaturation    string `vmt:"$envmapsaturation"`
	EnvmpaTint          string `vmt:"$envmaptint"`

	Flesh string `vmt:"$flesh"`

	MaxFogDensityScalar string `vmt:"$maxfogdensityscalar"`

	NoDraw    string `vmt:"$no_draw"`
	NoDecal   string `vmt:"$nodecal"`
	NormalMap string `vmt:"$normalmap"`
	NoTint    string `vmt:"$notint"`

	PointSampleMagFilter string `vmt:"$PointSampleMagFilter"`

	VertexAlpha int `vmt:"$vertexalpha"`
	VertexColor int `vmt:"$vertexcolor"`

	WriteZ int `vmt:"$writeZ"`
}

// NewProperties returns *Properties
func NewProperties() *Properties {
	return &Properties{}
}
