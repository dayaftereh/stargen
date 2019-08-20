
class Sun {

    constructor(data, camera) {
        this.data = data
        this.camera = camera
        this.enabled = true

        this.dt = 0;

        const color = new THREE.Color(this.data.color)
        this.material = new THREE.ShaderMaterial({
            uniforms: {
                time: { value: 1 },
                radius: { value: this.radius() },
                color: { value: new THREE.Vector3(color.r, color.g, color.b) }
            },
            vertexShader: `{{ template "star.vert" }}`,
            fragmentShader: `{{ template "star.frag" }}`,
        });

        this.geometry = new THREE.SphereBufferGeometry(this.radius(), 100, 100);
        this.mesh = new THREE.Mesh(this.geometry, this.material);
        this.mesh.name = "sun"

        const coronaSize = (this.radius() * 2.0) * 3.0
        const coronaGeometry = new THREE.PlaneBufferGeometry(coronaSize, coronaSize, 32);
        this.coronaMaterial = new THREE.ShaderMaterial({
            transparent: true,
            uniforms: {
                time: { value: 1 },
                coronaSize: { value: coronaSize },
                radius: { value: this.radius() }
            },
            vertexShader: `{{ template "corona.vert" }}`,
            fragmentShader: `{{ template "corona.frag" }}`,
        })
        this.corona = new THREE.Mesh(coronaGeometry, this.coronaMaterial);
        this.mesh.add(this.corona)
    }

    radius() {
        return radiusWorld(sunRadius());
    }

    update(delta) {
        this.corona.lookAt(this.camera.position)

        this.dt += delta * 0.05
        this.material.uniforms.time.value = this.dt
        this.coronaMaterial.uniforms.time.value = this.dt
    }

    color() {
        const r = this.data.effectiveTemperature * (0.0534 / 255.0) - (43.0 / 255.0)
        const g = this.data.effectiveTemperature * (0.0628 / 255.0) - (77.0 / 255.0)
        const b = this.data.effectiveTemperature * (0.0735 / 255.0) - (115.0 / 255.0)
        //effectiveTemperature
        const color = new THREE.Color(r, g, b)
        return color
    }
}