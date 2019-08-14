
class Planet {
    constructor(index, data) {
        this.index = index;
        this.data = data;
        this.epoch = (Math.random() * this.data.orbitPeriod) % this.data.orbitPeriod;

        this.material = new THREE.MeshBasicMaterial();
        this.geometry = new THREE.SphereBufferGeometry(this.radius(), 32, 32);
        this.mesh = new THREE.Mesh(this.geometry, this.material);
        this.orbit = createOrbitLines(this.semiMajorAxis(), 1000);
    }

    radius() {
        return sizeScaleRatio(this.data.radius);
    }

    semiMajorAxis() {
        return semiMajorAxis(au2km(this.data.semiMajorAxis));
    }

    update(delta) {
        this.epoch = (this.epoch + (10.0 * delta)) % this.data.orbitPeriod;
        const a = (2.0 * Math.PI) * (this.epoch / this.data.orbitPeriod);
        const radius = this.semiMajorAxis();
        this.mesh.position.x = radius * Math.cos(a);
        this.mesh.position.z = radius * Math.sin(a);
    }
}

