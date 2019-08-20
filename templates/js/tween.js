
class TweenManager {
    constructor() {
        this.queue = []
        this.enabled = true
    }

    update(delta) {
        if (!this.queue || this.queue.length < 1) {
            return
        }

        const head = this.queue[0]

        head.update(delta)

        if (!head.alive()) {
            this.queue.shift()
            head.dispose()
        }
    }

    add(tween) {
        tween.init()
        this.queue.push(tween)
    }
}

class Tween {

    init() {

    }

    update(delta) {

    }

    alive() {
        return true
    }

    dispose() {

    }
}

class LerpTween extends Tween {

    constructor(target, camera, control) {
        this.target = target
        this.camera = camera
        this.control = control
    }

    init() {
        this.control.enabled = false
    }

    update(delta) {
        // lerp camera toward target
        const position = this.camera.position.clone()
        this.camera.position = position.lerp(this.target.position, 0.1*delta)
    }

    alive() {
        // calculate distance between camera and target
        const distance = this.camera.position.distanceTo(this.target.position)
        return distance > 0.1
    }

    dispose() {
        this.control.enabled = true
    }

}