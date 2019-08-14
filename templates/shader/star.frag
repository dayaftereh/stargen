precision highp float;
precision highp int;

uniform float time;

uniform vec3 unColor;
uniform vec3 unCenterDir;
uniform float unRadius;

// input
varying vec3 vPosition;

{{ template "snoise3.glsl" }}

void main() {
    vec3 position = vPosition + time;
    float n = (noise(position, 4, 40.0, 0.7) + 1.0) * 0.5;

    // Get worldspace position
    vec3 sPosition = position * unRadius;
    
    // Sunspots
    float s = 0.36;
    float frequency = 0.00001;
    float t1 = snoise(sPosition * frequency) - s;
    float t2 = snoise((sPosition + unRadius) * frequency) - s;
	float ss = (max(t1, 0.0) * max(t2, 0.0)) * 2.0;
    // Accumulate total noise
    float total = n - ss;
	
	float theta = 1.0 - dot(unCenterDir, vPosition);
	
	gl_FragColor = vec4(unColor + (total - 0.5) - theta, 1.0);
}