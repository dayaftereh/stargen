
// {{ template "sun.js" . }}
// {{ template "world.js" . }}
// {{ template "tween.js" . }}
// {{ template "planet.js" . }}
// {{ template "star-system.js" . }}
// {{ template "interaction.js" . }}
// {{ template "fly-controls.js" . }}
// {{ template "trackball-controls.js" . }}

$(function () {
    const starSystem = new StarSystem()
    starSystem.init()
    // start the first update
    starSystem.update()
});

