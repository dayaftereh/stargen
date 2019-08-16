precision highp float;
precision highp int;

uniform float time;
uniform float radius;
uniform float coronaSize;

varying vec3 fPosition;

void main() {   

    vec3 nDistVec = normalize(fPosition);
    float noise2 = abs(.01/(abs(nDistVec.y)+.05)) * 2.3;

    float dist = length(fPosition);
    float scaleDist = (dist-radius)/coronaSize;
    float fade = 1.0 - 2.0*scaleDist;
    float totalStrength  = fade*(1.0 + noise2 );
    float intense = 0.8*totalStrength;
    float alpha = totalStrength - 2.0*scaleDist;
    
    gl_FragColor = vec4(intense,intense,intense,alpha);
}