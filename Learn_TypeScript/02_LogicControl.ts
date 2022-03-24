//! == 与 === 的区别
// ==： 如果两个值类型不同， 会自动转化为相同的类型进行比较
// ===： 就是其他语言中的 ==
function processHttpStatus(s:
    200 | 404 | 500 | '200' | '404' | '500') {

    //! 把 number 转化成 string
    let statusStr = ''
    if (typeof s === 'number') {
        statusStr = s.toString()
    } else {
        statusStr = s
    }

    //! 把 string 转化成 number
    let statusNum = 0
    if (typeof s === 'string') {
        statusNum = parseInt(s)
    } else {
        statusNum = s
    }
    //! 问号表达式
    //* 在 TypeScript 中，要经常使用 const，除非这个值是会变的
    //* 希望尽可能多地使用 const
    const statusNum2 = typeof s === 'string' ? parseInt(s) : s

    //* 在TypeScript中，一律使用 三个'='
    //* 不需要处理unkown的情况，因为传参的时候根本传不过来， 非常放心
    if (s === 200 || s === '200') {
        console.log('ok')
    } else if (s === 404 || s === '404') {
        console.log('not found')
    } else if (s === 500 || s === '500') {
        console.log('internal server error')
    }

    //! switch
    switch (statusNum) {
        case 200:
            console.log('ok')
            break
        case 404:
            console.log('not found')
            break
        case 500:
            console.log('internal server error')
            break
        // default: console.log('unkown error')
        //     break

    }
}
processHttpStatus(200)


//! 循环 和 错误处理
//* for
let sum = 0
for (let i = 0; i < 100; i++) {
    try {
        sum += i
        if (i % 17 === 0) {
            throw `bad number ${i}` //* 主动抛出一个错误
        }
    } catch (err) {
        console.error(err)
    }
}
console.log(sum)

//* while
let i = 0
while (i < 100) {
    sum += i
    i++
}
console.log(sum)
