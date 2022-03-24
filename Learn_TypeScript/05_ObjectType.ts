//! 对象类型
//* 直接定义一个对象
const emp1 = {
    // object 套 object
    name: {
        first: '三',
        last: '张'
    },

    // literal type
    gender: 'male' as 'male' | 'female' | 'other' | 'unknown',
    salary: 8000,
    // bonus: undefined,
    //* 也可以直接告诉编译器类型
    bonus: undefined as (number | undefined),
    performance: 3.5,
    bedges: ['优秀员工', '迟到王'], // 习惯是加行末逗号
}
emp1.gender = 'other'
console.log(`${emp1.name} has a salary of ${emp1.salary}`)

//! JSON   JavaScript Object Notation
//* console.log 可以将对象转换成json string 或者 JSON.stringify(item)
console.log(emp1)
const s: string = JSON.stringify(emp1)
console.log(s)

//* 通过 JSON.parse(jsonString) 可以将json string 转换成对象
const emp2 = JSON.parse(s)
console.log(emp2.name.first) // any类型用起来比较吃力，只能希望自己打字打正确了_(:з」∠)_

//! JSON object没有比较的能力
//* 不能 emp1 === emp2
