
// {{ template "sun.js" . }}
// {{ template "planet.js" . }}
// {{ template "fly-controls.js" . }}
// {{ template "trackball-controls.js" . }}

const width = window.innerWidth
const height = window.innerHeight;

const canvas = document.getElementById('canvas')

var renderer, scene, camera, control, clock, world
var planetObjects = []
var sunObject

function init() {
    // create the renderer
    renderer = new THREE.WebGLRenderer({
        canvas,
        antialias: true
    });
    renderer.setPixelRatio(window.devicePixelRatio);
    renderer.setSize(width, height);

    // create the camera
    camera = new THREE.PerspectiveCamera(70, window.innerWidth / window.innerHeight, 0.001, km2World(au2km(400)));
    camera.position.z = radiusWorld(sunRadius()*2.0)
    camera.position.y = km2World(au2km(1))
    camera.position.x = km2World(au2km(1))
    camera.lookAt(0, 0, 0)

    // create the scene
    scene = new THREE.Scene();

    const directionalLight = new THREE.DirectionalLight(0xffffff);
    directionalLight.position.set(au2km(4), 0, au2km(4));
    scene.add(directionalLight);

    const axesHelper = new THREE.AxesHelper(au2km(10));
    scene.add(axesHelper)

    world = new THREE.Object3D()
    scene.add(world)

    clock = new THREE.Clock()

    control = new TrackballControls(camera, canvas)
    //control.movementSpeed = km2World(au2km(0.1))

    // resize
    window.addEventListener('resize', onWindowResize, false);
}

function loadObjects() {
    // load the sun
    sunObject = new Sun(sun, camera)
    world.add(sunObject.mesh)

    // load all planets
    planets.forEach((planet, index) => {
        const planetObject = new Planet(index, planet, camera)
        planetObjects.push(planetObject)
        world.add(planetObject.mesh)
        world.add(planetObject.orbit)
    })
}

function onWindowResize() {
    camera.aspect = window.innerWidth / window.innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(window.innerWidth, window.innerHeight);
}

function animate() {
    requestAnimationFrame(animate);

    const delta = clock.getDelta()

    control.update(delta)

    sunObject.update(delta)
    planetObjects.forEach(planetObject => {
        planetObject.update(delta)
    })

    renderer.render(scene, camera);
}

init()
loadObjects()
animate()