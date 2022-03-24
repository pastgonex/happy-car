//! 基本数据类型 number string boolean 
let anExampleNum = 123
const anExampleVariable = 'Hello World'
let anExampleBool: boolean = true
console.log(anExampleVariable, anExampleNum, anExampleBool)

//! literal types
let httpStatus: 200 | 404 | 500 = 200
let answer: 'yes' | 'no' | 'maybe' = 'maybe'
// 这里有两种类型，在给anExampleUnionType复制的时候，主要看'='后面的类型
let anExampleUnionType: 200 | 404 | 500 | '200' | '404' | '500' = '200'

//! 函数的定义
function f(s: 200 | 404 | 500 | '200' | '404' | '500') {
    let status: string | number = s // 这就是union of types
}
//! 函数的调用
f(httpStatus)

//! any类型
//* 使用场景：在操作一些第三方库的时候（缺失一些类型），我们只能定义成any，根据文档操作（代码比较难写）
// 加了any之后， 就放弃了 typescript的所有检查
// javascript怎么样， typescript就怎么样
// 不报错了，但是不代表代码是错的
/*
    let a: any = 'abc'
    a = 123
    a.name = 'John'
    console.log(a.salary)
*/

//! undefined 类型
let b: undefined
console.log(b)
// 有时候会有不回答的场景，那么这个时候就使用undefined
// let answer: 'yes' | 'no' | 'maybe' | undefined= undefined