//! 函数
//* 函数定义
// 建议： 如果是自己用，就不写返回值类型
//                                 小技巧            默认值       可变参数列表
function add(a: number, b: number, c?: number, d: number = 999, ...e: number[]): number {
    let sum = a + b + (c || 0) + d
    for (let i = 0; i < e.length; i++) {
        sum += e[i]
    }
    return sum
}
console.log(add(1, 2))
console.log(add(9, 1, 10))

const numbers = [5, 6, 7, 8, 9, 10]
console.log(add(1, 2, 3, ...numbers))

//! 参数过长， 就使用一个对象参数
function sendRequest(parames: {
    url: string,
    method: 'GET' | 'POST' | 'PUT' | 'DELETE',
    header: object
    data?: string,
    requireAuth: boolean,
    retry: boolean,
    retryTimeout?: number,
}) {
    console.log(parames)
}  

// 调用
sendRequest({
    url: 'http://www.baidu.com',
    method: 'GET',
    header: {
        'Content-Type': 'application/json'
    },
    requireAuth: true,
    retry: true,
})

//! 方法
const emp_06 = {
    name: 'john',
    salary: 8000,
    bonus: undefined as (number | undefined),
    performance: 3.5,
    updateBonus() {
        if (!this.bonus) {
            this.bonus = this.salary * this.performance
        }
    }
}
emp_06.updateBonus()

//* this关键字
function registerEmployee(p: any) {

}
registerEmployee({
    name: 'john',
    salary: 8000,
    bonus: undefined as (number | undefined),
    performance: 3.5,
    updateBonus() {
        if (!this.bonus) { // this.bonus
            this.bonus = this.salary * this.performance
        }
    }
})