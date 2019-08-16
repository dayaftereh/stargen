precision highp float;
precision highp int;

varying vec3 fPosition;

void main() {   
    fPosition = position;
    gl_Position = projectionMatrix * modelViewMatrix * vec4( position, 1.0 );
    //gl_Position = projectionMatrix * (modelViewMatrix * vec4(0.0, 0.0, 0.0, 1.0) + vec4(position.x, position.y, 0.0, 0.0));
}