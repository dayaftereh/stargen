
class World {
    constructor(scene, camera) {
        this.scene = scene
        this.camera = camera
        this.enabled = true

        this.sun = undefined
        this.planets = undefined
    }

    init() {
        this.initSun()
        this.initLight()
        this.initPlanets()

        this.axesHelper()
    }

    initLight() {
        const directionalLight = new THREE.DirectionalLight(0xffffff);
        directionalLight.position.set(1.0, 0, 1.0);

        this.scene.add(directionalLight);
    }

    initSun() {
        // create the sun
        this.sun = new Sun(sun, this.camera)
        // add the sun to the world
        this.scene.add(this.sun.mesh)
    }

    initPlanets() {
        this.planets = planets.map((data, index) => {
            // create the planet
            const planet = new Planet(index, data, this.camera)
            // add the planet to scene
            this.scene.add(planet.mesh)
            this.scene.add(planet.orbit)

            return planet
        })
    }

    axesHelper() {
        const axesHelper = new THREE.AxesHelper(au2km(10));
        this.scene.add(axesHelper)
    }

    update(delta) {
        // update object
        [this.sun, ...this.planets].filter(object => object.enabled).forEach(object => {
            object.update(delta)
        })
    }

}