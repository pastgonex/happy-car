//! 数组
//* number[] 等价于 Array<number> 泛型
//* array.length 取数组的长度
let a = [1, 2, 3, 4, 5, 6, 7, 8]
// let a: number[] = [1, 2, 3]
// let a: Array<number> = [1, 2, 3]

//* 这两种写法等价
let bb = [1, 2, 3, "Hello world"]
let c: (string | number)[] = [1, 2, 3, "Hello world"]

//* 如果是空的话，类型是一定要加上去的
let d: Array<number> = []

//! 遍历
//* 不推荐使用 for 循环，因为 for 循环的语法比较简单，但是它的性能比 forEach 慢很多

//! 在 Typescript 中判断数组是否为空，可以使用 Array.isEmpty 或者 Array.length方法
if (a.length) {
    console.log('a is not empty')
} else {
    console.log('a is empty')
}

//! 数组操作
//* 插    删
//* push pop 从尾部操作
//* unshift shift 从头部操作

//* slice(start, end) 从指定位置开始截取  [start, end)
//* slice(start) 从指定位置到最后
//* splice(startDist, deleteCount, ...) 从指定位置开始删除 deleteCount 个元素, 后面可以添加元素（用逗号隔开）

//* 查找 indexOf(item) 从头部开始查找，返回第一个找到的元素的索引，如果没有找到，返回-1
//*     indexOf(item, start) 从start位置开始找，返回第一个找到的元素的索引，如果没有找到，返回-1
//*     lastIndexOf(item) 从后面开始查找，返回第一个找到的元素的索引，如果没有找到，返回-1

//? 排序 sort() 默认是字典序

a.push(999)
a.pop()
a.unshift(111)
a.shift()
console.log(a.slice(2, 5))
a.splice(2, 2) // 从数组下标为2的元素开始， 删除两个元素
console.log(a)
a.splice(2, 2, 999, 999, 999)
console.log(a)

//! const Array
//* 表示只能是当前这个数组， 不能重新赋值一个新数组
const constArray: number[] = []
// constArray = [1, 2, 3] // error
constArray.push(1) //* ok



//! 元组 tuple
const baseArray = [1, 2, 3]
const [tuple1, tuple2] = baseArray // 会把 baseArray展开
console.log(tuple1, tuple2) // 1， 2

//! split/join
console.log('1, 2, 3, a, b, c'.split(',')) // 将string通过指定的分隔符分割->数组
console.log([1, 2, 3, 'a', 'b', 'c'].join()) // 将Array转换为string，并用指定的分隔符连接(默认是逗号)


function cmp(a:number, b:number) {
    //* 如果 a 小于 b，返回 正数
    //* 如果 a 等于 b，返回 0
    //* 如果 a 大于 b，返回 负数
    return a - b
}

// Sorting of Array
let array1 = [5, 2, 1, 6, 8, 10, 5, 25, 16, 23, 11]
array1.sort(cmp)
console.log(array1)
