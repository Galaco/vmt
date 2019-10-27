package vmt

// Material represents expected properties in a particular material
// It is important to define the following property in any implementation:
// 	ShaderName string `vmt:"__SHADER_NAME__"`
// Properties should be defined in the following format:
// 	BaseTexture string `vmt:"$basetexture"`
type Material interface{}
