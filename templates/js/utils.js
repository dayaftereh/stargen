
function worldScale() {
    return 1000.0
}

function au2km(au) {
    return au * 149597870.700
}

function km2au(km) {
    return km / 149597870.700
}

function sunRadius() {
    const radiusKM = Math.sqrt((sun.luminosity * 3.828e26) / (4.0 * Math.PI * (5.7e-8) * (Math.pow(sun.effectiveTemperature, 4.0)))) / 1000.0
    return radiusKM
}

function maxSemiMajorAxis(){
    const max = Math.max(...planets.map(planet => {
        return planet.semiMajorAxis
    }))
    return max
}

/* KM -> world */
function radiusWorld(km) {
    const worldKm = km2World(km)
    const sunRatio = km / sunRadius()*2.0
    const fillRatio = (1.0 - sunRatio) * 0.25
    return worldKm + fillRatio + sunRatio
}

/* km -> world */
function km2World(km) {
    const max = maxSemiMajorAxis()
    return (km / au2km(max)) * worldScale()
}

function createOrbitLines(radius, lines) {
    const vertices = []

    for (var i = 0; i <= lines; i++) {
        const a = (2.0 * Math.PI) * (i / lines)
        const x = radius * Math.cos(a)
        const y = radius * Math.sin(a)
        vertices.push(x, 0.0, y)
    }

    const lineGeometry = new THREE.BufferGeometry();
    lineGeometry.addAttribute('position', new THREE.Float32BufferAttribute(vertices, 3));

    var lineMaterial = new THREE.LineDashedMaterial({
        color: 0xdddddd,
    });

    var line = new THREE.Line(lineGeometry, lineMaterial);
    return line
}

console.log(sun.class)
console.log(`Sun Radius: ${sunRadius()} km / ${km2au(sunRadius())} AU`);
console.log(`sunRadiusWorld: ${radiusWorld(sunRadius())}`);

