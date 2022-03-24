//! 函数式编程
function cmp1(a: number, b: number) {
    //* 如果 a 小于 b，返回 正数
    //* 如果 a 等于 b，返回 0
    //* 如果 a 大于 b，返回 负数
    return a - b
}

//! 函数作为"一等公民"
/*
函数作为"一等公民"
变量类型可以是函数 
值(literal)可以是函数 
对象的字段可以是函数 
函数的参数可以是函数 
函数的返回值可以是函数
*/

//! 值(literal)可以是函数
// Sorting of Array
let compareNumber = (a: number, b: number) => a - b
let arr1 = [5, 2, 1, 6, 8, 10, 5, 25, 16, 23, 11]
arr1.sort(cmp1) //* 将函数作为参数，这是函数式编程的一种
arr1.sort((a, b) => a - b) //* lambda表达式,箭头函数
arr1.sort(compareNumber)
console.log(arr1)

//! 对象的字段可以是函数
const empp1 = {
    name: 'john',
    salary: 8000,
    // increaseSalary(p: number) {
    //     this.salary *= p
    // }
    increaseSalary: function (p: number) {
        this.salary *= p
    }
}
empp1.increaseSalary(1.1)
console.log(empp1)


//! 函数的返回值可以是函数
function createComparer(p: { smallerFirst: boolean }) {
    if (p.smallerFirst) {
        return (a: number, b: number) => a - b
    }
    return (a: number, b: number) => b - a
}
let a2 = [5, 2, 1, 6, 8, 10, 5, 25, 16, 23, 11]
a2.sort(createComparer({
    smallerFirst: true
}))

function loggingComparer(
    logger: (a: number, b: number) => void,
    comp: (a: number, b: number) => number) {
    return (a: number, b: number) => {
        // 副作用
        // console.log('comparing', a, b)
        // compCount++
        logger(a, b)
        return comp(a, b)
    }
}

//! 闭包
function processArray(a: number[]) {
    let compCount = 0

    // logger: 闭包
    //* 延长局部变量的生命周期
    const logger = (a: number, b: number) => {
        console.log('comparing', a, b)
        compCount++ // 自由变量
    }
    const comp = createComparer({ smallerFirst: true })
    a.sort(loggingComparer(logger, comp))
    return compCount
}

let a3 = [5, 2, 1, 6, 8, 10, 5, 25, 16, 23, 11]
const compCount = processArray(a3)
console.log(a3)
console.log('compare count', compCount)

//! 部分应用函数 
//* 例子1
function isGoodNumber(goodFactor: number, v: number) {
    return v % goodFactor === 0
}

function filterArray(a: number[], f: (v: number) => boolean) {
    return a.filter(f)
}

//config
const GOOD_FACTOR = 2
// end config

const aaa = [1, 2, 3, 4, 5, 6, 7, 8, 9]
console.log(
    //                     虽然参数数量不匹配，不要紧，可以用这种方式运用起来
    filterArray(aaa, (v) => isGoodNumber(GOOD_FACTOR, v))
)

//* 例子2
// a, b => boolean
// 提供了a， 那么 b => f(a, b)
function partiallyApply(f: (a: number, b: number) => boolean, a: number) {
        // a是自由变量，从花括号外面可以访问
        return (b: number) => f(a, b)
}
console.log(
    filterArray(aaa, partiallyApply(isGoodNumber, GOOD_FACTOR))
)