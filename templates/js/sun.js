
class Sun {

    constructor(data) {
        this.data = data

        const color = new THREE.Color(data.color)

        this.material = new THREE.ShaderMaterial({

            uniforms: {
                time: { value: 1 },
                unRadius: { value: 1 },
                unCenterDir: { value: new THREE.Vector3(1, 1, 1) },
                unColor: { value: new THREE.Vector3(0, 0, 1.0) }
            },
            vertexShader: `{{ template "star.vert" }}`,
            fragmentShader: `{{ template "star.frag" }}`,
        });

        this.geometry = new THREE.SphereBufferGeometry(this.radius(), 32, 32);
        this.mesh = new THREE.Mesh(this.geometry, this.material);
    }

    radius() {
        return sizeScaleRatio(sunRadius());
    }

    update(delta) {
        this.material.uniforms.time.value = Date.now()

    }
}