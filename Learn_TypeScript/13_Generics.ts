//! 泛型

const a: Array<string> = []

const p = new Promise<string>((resolve, reject) => {
    resolve('122')
})

// 定义泛型参数
class MyArray<T> {
    data: T[] = []
    add(t: T) {
        this.data.push(t)
    }
    print() {
        console.log(this.data)
    }

    map<U>(f: (v: T) => U) {
        return this.data.map(f)
    }
}

const a_13_1 = new MyArray<number>()
a_13_1.add(1)
a_13_1.add(2)
console.log(a_13_1.map(v => v.toExponential()))
a_13_1.print()


//! 对传入的T做一个限制, 用接口限制
interface hasWeight {
    weight: number
}

class MyArray2<T extends hasWeight> {
    data: T[] = []
    add(t: T) {
        this.data.push(t)
    }
    print() {
        console.log(this.data)
    }

    map<U>(f: (v: T) => U) {
        return this.data.map(f)
    }

    sortByWeight() {
        this.data.sort((a, b) => a.weight - b.weight)
    }
}


class T1 {
    weight: 15
}
const a_13_2 = new MyArray2<T1>()
a_13_2