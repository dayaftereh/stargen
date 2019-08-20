
class Planet {
    constructor(index, data, camera) {
        this.index = index;
        this.data = data;
        this.camera = camera
        this.enabled = true
        this.epoch = (Math.random() * this.data.orbitPeriod) % this.data.orbitPeriod;

        // create the planet mesh
        this.material = new THREE.MeshBasicMaterial();
        this.geometry = new THREE.SphereBufferGeometry(this.radius(), 32, 32);
        this.mesh = new THREE.Mesh(this.geometry, this.material);
        this.mesh.name = `planet-${index}`

        // create the orbit
        this.orbit = createOrbitLines(this.semiMajorAxis(), 1000);

        // create the ring
        const ringGeometry = new THREE.RingBufferGeometry(0.9, 1.0, 100);
        const ringMaterial = new THREE.MeshBasicMaterial();
        this.ring = new THREE.Mesh(ringGeometry, ringMaterial);
        this.ring.rotation.set(-Math.PI / 2.0, 0.0, 0.0)
        this.mesh.add(this.ring)

        // add axes helper
        const axesHelper = new THREE.AxesHelper(this.radius() * 4.0);
        this.mesh.add(axesHelper)
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

        const position = new THREE.Vector3()
        position.x = radius * Math.cos(a);
        position.z = radius * Math.sin(a);

        this.mesh.position.copy(position)

        const normal = position.clone().multiplyScalar(-1.0)

        const tangent = new THREE.Vector3(-normal.x, normal.y, normal.z)

        const quaternion = new THREE.Quaternion()
        quaternion.setFromUnitVectors(tangent.clone().normalize(), normal.clone().normalize())

        this.mesh.rotation.setFromQuaternion(quaternion)

    }

}

