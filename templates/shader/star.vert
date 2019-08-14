precision highp float;
precision highp int;

// Uniforms
uniform mat4 unWVP;

// Input
//attribute vec3 position;

// Output
varying vec3 vPosition;

void main() {
    vPosition = position.xyz; 
    gl_Position = projectionMatrix * modelViewMatrix * vec4( position, 1.0 );
}