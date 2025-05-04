package shaders 

import (
	"log"
	"github.com/go-gl/gl/v4.1-core/gl"
)

const vertexShaderSource = `
#version 410 core
layout (location = 0) in vec3 aPos;

void main(){
	gl_Position = vec4(aPos, 1.0);
}`

const fragmentShaderSource = `
#version 410 core
out vec4 FragColor;

void main(){
	FragColor = vec4(1.0, 1.0, 1.0, 0.7);
}`

func CompileAndAttachShaders(program uint32) {
	vertexShader := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
}

func compileShader(source string, shaderType uint32) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource, freeMemory := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSource, nil)
	freeMemory()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &logMsg[0])

		log.Fatalf("\nShader compile error. \nType: %d\n%s\n", logLength, logMsg)
	}
	return shader
}
