function f(goodFactor: number, v: number) {
    return v % goodFactor === 0
}

function filterA(a: number[], f: (v: number) => boolean) {
    return a.filter(f)
}

const gf = 2
const a4 = [1, 2, 3, 4 , 5, 6, 7, 8, 9]
console.log(filterA(a4, (v) => f(gf, v)))
