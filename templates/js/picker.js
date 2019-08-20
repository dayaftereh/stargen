
function onMouseUp(event) {
    const location = getScreenLocation(event)
    const hits = raycast(location)

    console.log("isSun:", isSun(hits))

    const planet = findPlanet(hits)

    console.log("Planet:", planet)
}

function getScreenLocation(event) {
    const rect = canvas.getBoundingClientRect();
    const x = (event.clientX - rect.left) / (rect.right - rect.left) * 2.0 - 1.0
    const y = -(event.clientY - rect.top) / (rect.bottom - rect.top) * 2.0 + 1.0
    const location = new THREE.Vector2(x, y);
    return location
}

function raycast(location) {
    const raycaster = new THREE.Raycaster();
    //raycaster.linePrecision = 0.1
    raycaster.setFromCamera(location, camera)
    const hits = raycaster.intersectObjects(world.children)

    if (!hits || hits.length < 1) {
        return []
    }

    return hits
}

function isSun(hits) {
    const found = hits.some(hit => {
        return hit.object === sunObject.mesh
    })
    return found
}

function findPlanet(hits) {
    const planet = planetObjects.find(planetObject => {
        const isHit = hits.some(hit => {
            return isPlanetHit(hit, planetObject)
        })
        return isHit
    })

    return planet
}

function isPlanetHit(hit, planet) {
    return hit.object === planet.mesh || hit.object === planet.ring // || hit.object === planet.orbit
}

function initPicker() {
    
    planets.forEach((planet, index) => {
        $(`#3d-dropdown-planet-${index}`).click(function () {
            console.log("planet", planet)
        })
    })
    // mouse up
    canvas.addEventListener("mouseup", onMouseUp, false)
}