precision highp float;
precision highp int;

// uniform 
uniform float time;
uniform float radius;
uniform vec3 color;

// input
varying vec3 fPosition;

{{ template "snoise3.glsl" }}

void main() {
    vec3 position = fPosition + time;
    float n = (noise(position, 6, 7.0, 0.15) + 1.0) * 0.5;

    // Get worldspace position
    vec3 sPosition = position * radius;

    // sunspots
    float s = 0.10;
    float frequency = 0.05;
    float t1 = snoise(sPosition*frequency) - s;
    float t2 = snoise((sPosition+radius)*frequency) - s;
    float ss = (max(t1, 0.0) * max(t2, 0.0))*2.0;

    // Accumulate total noise
    float total = n - ss;
    gl_FragColor = vec4(color + (total-0.5), 1.0);
}