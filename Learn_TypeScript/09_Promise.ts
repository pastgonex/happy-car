//! 回调函数的缺点
function add1(a: number, b: number) {
    return a + b
}

//* 必须要等add函数结束才能执行下面的代码， 假设add函数很花时间， 那么程序就一直卡在这里了
//* 前端单线程， 这个等待是不能等的， 否则就会卡住
console.log(add1(2, 3))

function callbackAdd(a: number, b: number, callback: (res: number) => void): void {
    setTimeout(() => {
        callback(a + b)
    }, 2000)
}
callbackAdd(2, 3, res => {
    console.log('2+3', res)
    //* 如果要再加4， 这里就还得再套一层
    callbackAdd(res, 4, res2 => {
        console.log('2 + 3 +4', res2)
    })
}) 

//! Promise
function promiseAdd(a: number, b: number): Promise<number> {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve(a + b)
        }, 2000)
    })
}

//* Promise的好处是可以把多个异步操作串起来
promiseAdd(2, 3).then(res => {
    console.log('2+3', res)
    return promiseAdd(res, 4)
}).then(res => {
    console.log('2+3+4', res)
})

// 不打log版本
promiseAdd(2, 3).then(res => promiseAdd(res, 4)).then(res => promiseAdd(res, 6)).then(res => {
    console.log('final result', res)
}).catch(err => { // 错误处理
    console.log('caugh error', err)
})

function mul(a: number, b: number): Promise<number> {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve(a * b)
        }, 3000)
    })
}
// (2+3)*4+5
promiseAdd(2, 3).then(res => {
    console.log('2+3', res)
    return mul(res, 4)
}).then(res => {
    console.log('(2+3)*4', res)
    return promiseAdd(res, 5)
}).then(res => {
    console.log('(2+3)*4+5', res)
}).catch(err => {
    console.log('caugh error', err)
})

// (2+3)*(4+5)
//* 我们希望等待2+3 和 4+5同时执行后， 做一个乘法
//* 同时
Promise.all([promiseAdd(2, 3), promiseAdd(4, 5)]).then(([a, b]) => {
    console.log('result', a, b)
    return mul(a, b)
}).then(res => {
    console.log('fianl result', res)
})

//* Promise.race 等待一个返回即可
Promise.race([promiseAdd(2, 3), promiseAdd(4,5)]).then(res => {
    console.log(res)
})

