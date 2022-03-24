const a = [1, 2, 3, 4]

// 遍历
for (let i = 0; i < a.length; i++) {

}

// 有副作用的
const b: number[] = []
a.forEach(v => {
    b.push(v * v)
})
console.log(b)

// 无副作用的
const b2 = a.map(v => v * v)
console.log(b2)
const c2 = b.reduce((s, v) => s + v)

// 手写数组的 reduce
function reduce(b: number[], r: (s: number, v: number) => number) {
    let previousVale = 0
    b.forEach(currentValue => {
        previousVale = r(previousVale, currentValue)
    })
    return previousVale
}

// 整合一下 
console.log(
    [1, 2, 3, 4]
    .map(v=>v*v)
    .reduce((s, v) => s + v)
)