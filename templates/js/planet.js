
class Planet {
    constructor(index, data, camera) {
        this.index = index;
        this.data = data;
        this.camera = camera
        this.epoch = (Math.random() * this.data.orbitPeriod) % this.data.orbitPeriod;

        this.material = new THREE.MeshBasicMaterial();
        this.geometry = new THREE.SphereBufferGeometry(this.radius(), 32, 32);
        this.mesh = new THREE.Mesh(this.geometry, this.material);
        this.orbit = createOrbitLines(this.semiMajorAxis(), 1000);

        const ringGeometry = new THREE.RingBufferGeometry(0.9, 1.0, 100);
        const ringMaterial = new THREE.MeshBasicMaterial();
        const ring = new THREE.Mesh(ringGeometry, ringMaterial);
        ring.rotation.set(-Math.PI / 2.0, 0.0, 0.0)

        this.mesh.add(ring)

        const axesHelper = new THREE.AxesHelper(1);
        ring.add(axesHelper)
        //this.mesh.add(axesHelper)
    }

    radius() {
        return radiusWorld(this.data.radius);
    }

    semiMajorAxis() {
        return km2World(au2km(this.data.semiMajorAxis));
    }

    update(delta) {
        this.epoch = (this.epoch + (10.0 * delta)) % this.data.orbitPeriod;
        const a = (2.0 * Math.PI) * (this.epoch / this.data.orbitPeriod);
        const radius = this.semiMajorAxis();
        this.mesh.position.x = radius * Math.cos(a);
        this.mesh.position.z = radius * Math.sin(a);
    }
}

