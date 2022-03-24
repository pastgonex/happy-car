// 固定下面的对象的信息
interface Employee {
    name: string
    salary: number
    bonus?: number
    updateBonus(p: number): void // 如果有这种方法的话，我们很可能选择 类 来定义
}
const emp_11: Employee = {
    name: "Ever",
    salary: 350000,
    updateBonus(p: number) {

    }
}

const emp_11_1: Employee = {
    name: "Ever",
    salary: 350000,
    bonus: 100000,
    // 加了方法之后， 就比较恶心了， 每次都要抄一遍这个方法  接口 => 类
    updateBonus(p: number) {

    }
}

let emp3: Employee // 这样声明， 就是没有定义任何的属性， 只是声明了一个类型


// 使用 "?" 来保护
interface Employee2 extends HasName { //! 将两个接口扩展起来
    name: {
        first: string
        last: string
    }
    salary: number
    bonus?: number
}

interface HasName {
    name?: {
        first?: string
        last: string
    }
}

function hasBadName(e: Employee2) {
    // if (e.name && e.name.first) {
    //     return e.name.first.startsWith('AAA')
    // }
    // return true
    //! 可以使用 "?", 这样就不会报错了
    // return e.name?.first?.startsWith('AAA')

    //* 可以使用 "!", 告诉编译器一定有  非空断言
    return e.name!.first!.startsWith('AAA')
}

console.log(hasBadName({
    name: {
        first: 'john',
        last: 'smith',
    },
    salary: 8000
}))

//! 类型的并， 类型断言， 类型判断
interface WxButton {
    visible: boolean
    enable: boolean
    onClick(): void
}

interface WxImage {
    visible: boolean
    src: string
    width: number
    height: number
}

function processElement(e: WxButton | WxImage) {
    // e. 能点出来的是 这两个interface共有的filed
    //* 低级写法
    // if ((e as any).onClick) {
    //     const btn = e as WxButton  //! e as WxButton 是 类型断言
    //     btn.onClick()
    // } else {
    //     const img = e as WxImage
    //     console.log(img.src)
    // }

    //* 高级写法
    if (isButton(e)) {
        e.onClick()
    } else {
        console.log(e.src)
    }
}

function isButton(e: WxButton | WxImage): e is WxButton {
    return (e as any).onClick !== undefined // 语法再高级，最终还是要通过属性判断
}

processElement({
    visible: true,
    enable: true,
    onClick() {
        console.log('clicked')
    }
})