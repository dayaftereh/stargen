
// {{ template "fly-controls.js" . }}

const width = window.innerWidth
const height = window.innerHeight;

const canvas = document.getElementById('canvas')

var renderer, scene, camera, control, clock, world

function init() {
    // create the renderer
    renderer = new THREE.WebGLRenderer({
        canvas,
        antialias: true
    });
    renderer.setPixelRatio(window.devicePixelRatio);
    renderer.setSize(width, height);

    // create the camera
    camera = new THREE.PerspectiveCamera(70, window.innerWidth / window.innerHeight, 0.00001, 1e9);
    camera.position.z = 10000
    camera.position.y = 10
    camera.lookAt(0, 0, 0)

    // create the scene
    scene = new THREE.Scene();

    world = new THREE.Object3D()
    scene.add(world)

    clock = new THREE.Clock()
    control = new FlyControls(camera, canvas)

    // resize
    window.addEventListener('resize', onWindowResize, false);
}

function loadObjects() {
    // load the sun
    const geometry = new THREE.SphereGeometry(1, 32, 32);
    const material = new THREE.MeshBasicMaterial({
        color: sun.color,
    });

    const sunMesh = new THREE.Mesh(geometry, material);
    world.add(sunMesh)

    // load all planets
    planets.forEach(planet => {
        console.log(km2au(planet.radius)*1000.0, planet.semiMajorAxis * 1000.0)
        const geometry = new THREE.SphereGeometry(km2au(planet.radius) * 1000.0, 32, 32);
        const material = new THREE.MeshBasicMaterial({
            //color: 'rgb(123,45,212)'
        });

        const planetMesh = new THREE.Mesh(geometry, material);
        planetMesh.position.x = planet.semiMajorAxis * 1000.0
        world.add(planetMesh)
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

    renderer.render(scene, camera);
}

init()
loadObjects()
animate()