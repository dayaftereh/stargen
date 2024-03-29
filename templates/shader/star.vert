precision highp float;
precision highp int;

varying vec3 fPosition;

void main() {
    fPosition = position.xyz; 
    gl_Position = projectionMatrix * modelViewMatrix * vec4( position, 1.0 );
}