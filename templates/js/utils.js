
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

/* KM -> world */
function sizeScaleRatio(size) {
    const x = size / sunRadius()
    return x * 10.0
}

function semiMajorAxis(km) {
    const kmOffset = km - sunRadius()
    return sizeScaleRatio(sunRadius()) + distanceRatio(kmOffset)
}

/* KM -> world */
function distanceRatio(km) {
    const max = Math.max(...planets.map(planet => {
        return planet.semiMajorAxis
    }))
    return (km / au2km(max)) * 100.0
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