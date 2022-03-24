function add_10(a: number, b: number): Promise<number> {
    return new Promise((resolve, reject) => {
        if (b % 17 === 0) {
            reject(`bad number: ${b}`)
        }
        setTimeout(() => {
            resolve(a + b)
        }, 1000)
    })
}

function mul_10(a: number, b: number): Promise<number> {
    return new Promise((resolve, reject) => {
        if (b % 17 === 0) {
            reject(`bad number: ${b}`)
        }
        setTimeout(() => {
            resolve(a * b)
        }, 2000)
    })
}

//! async await 语法糖

// (2+3) * (4+5)
// promise可以通过async来使用
// await一定要在async函数中使用
async function calc() { // 返回的是一个 Promise
    try {
        // const a = await add_10(2, 3) // await只是一个语法糖，就相当于add(2,3).then(...)
        // console.log('2+3', a)
        // const b = await add_10(4, 5)
        // console.log('4+5', b)
        // const c = await mul_10(a, b)
        // console.log('(2+3)*(4+5)', c)

        //* 改写同时算 2+3 和 4+5
        const [a, b] = await Promise.all([add_10(2, 3), add_10(4, 5)])
        console.log('2+3', a)
        console.log('4+5', b)
        return await (mul_10(a, b))
    } catch {
        console.log('error')
        return undefined
    }
}

calc()