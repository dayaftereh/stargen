precision highp float;
precision highp int;

varying vec3 fPosition;

void main() {
    gl_Position = projectionMatrix * modelViewMatrix * vec4( position, 1.0 );
}