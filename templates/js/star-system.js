
class StarSystem {

    constructor() {
        this.world = undefined
        this.scene = undefined
        this.clock = undefined
        this.camera = undefined
        this.renderer = undefined
        this.interaction = undefined
        this.tweenManager = undefined
    }

    init() {
        const canvas = this.canvas()
        const dimension = this.dimension()

        // create the renderer
        this.renderer = new THREE.WebGLRenderer({
            canvas,
            antialias: true
        });
        this.renderer.setPixelRatio(window.devicePixelRatio);
        this.renderer.setSize(dimension.width, dimension.height);

        // create the camera
        this.camera = new THREE.PerspectiveCamera(70, dimension.width / dimension.height, 0.001, worldScale() * 3.0);
        this.camera.position.x = radiusWorld(sunRadius())
        this.camera.position.y = radiusWorld(sunRadius() * 10.0)
        this.camera.position.z = radiusWorld(sunRadius() * 10.0)

        // create the clock for delta ticks
        this.clock = new THREE.Clock()

        // create the scene
        this.scene = new THREE.Scene();

        // create the world
        this.world = new World(this.scene, this.camera)
        this.world.init()

        // create the default control
        this.control = new TrackballControls(this.camera, canvas)

        // create the tweenManager
        this.tweenManager = new TweenManager()

        // create the interaction
        this.interaction = new Interaction(this.world, this.tweenManager, this.control)
        this.interaction.init()

        // resize
        window.addEventListener('resize', this.onWindowResize.bind(this), false);
    }

    dimension() {
        return {
            width: window.innerWidth * 0.95,
            height: window.innerHeight * 0.8
        }
    }

    canvas() {
        const canvas = $('#canvas')
        if (!canvas) {
            throw new Error(`unable to find canvas`)
        }
        return canvas[0]
    }

    onWindowResize() {
        const dimension = this.dimension()
        this.camera.aspect = dimension.width / dimension.height;
        this.camera.updateProjectionMatrix();
        this.renderer.setSize(dimension.width, dimension.height);
    }

    update() {
        // register next update
        requestAnimationFrame(this.update.bind(this));

        // calculate tick delta
        const delta = this.clock.getDelta()

        // check if control enabled
        if (this.control.enabled) {
            // update the control
            this.control.update(delta)
        }

        // check if world updated
        if (this.world.enabled) {
            // update the world
            this.world.update(delta)
        }

        // update the tween manager
        if (this.tweenManager.enabled) {
            this.tweenManager.update(delta)
        }

        // render the scene with camera
        this.renderer.render(this.scene, this.camera);
    }

}