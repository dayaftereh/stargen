
class Interaction {

    constructor(world, tweenManager, control) {
        this.world = world
        this.control = control
        this.tweenManager = tweenManager
    }

    init() {
        // register the sun selection
        $('#3d-dropdown-sun').click(function () {
            this.onSunSelected()
        })

        // register selection for each planet
        this.world.planets.forEach((planet, index) => {
            $(`#3d-dropdown-planet-${index}`).click(function () {
                this.onPlanetSelected(planet)
            })
        })
    }

    onPlanetSelected(planet) {
        const lerpTween = new LerpTween(planet.mesh, this.world.camera, this.control)
        //this.tweenManager.add(lerpTween)
    }

    onSunSelected() {
        const mesh = this.world.sun.mesh
        const lerpTween = new LerpTween(mesh, this.world.camera, this.control)
        //this.tweenManager.add(lerpTween)
    }
}