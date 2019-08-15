precision highp float;
precision highp int;

varying vec3 fPosition;

void main() {
    float dist = length(fPosition) * 4.0
    float brightness = (1.0/(dist*dist)-0.1)*0.7
    gl_FragColor = vec4(brightness, brightness, brightness, 1.0);
}