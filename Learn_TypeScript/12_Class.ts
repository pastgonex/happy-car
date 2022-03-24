//! 在TypeScript中使用类的场景不多， 因为有object
class Employee_12 {
    name: string
    private bonus: number = 0
    constructor(name: string) {
        this.name = name
    }
}

// 可以使用传统方式 实现 接口， 这样报错就更加清晰了
class Employee_12_1 implements interfaceEmploy {
    // 在这里加上pulic 和 private， 就相当于上面定义了field
    private allocatedBonus?: number
    constructor(public name: string, public salary: number) {
        this.name = name, this.salary = salary
    }
    updateBonus() {
        if (!this.bonus) {
            this.bonus = 20000;
        }
    }

    //! getter/setter
    set bonus(v: number) {
        this.allocatedBonus = v
    }

    get bonus() {
        return this.allocatedBonus || 0
    }
}



const emp_12 = new Employee_12('john')
const emp_12_1 = new Employee_12_1('Ever', 350000)
emp_12_1.bonus = 80000 // setter就像普通的 field一样使用


//! 继承
class Manager extends Employee_12_1 {
    private reporters: Employee_12_1[]
    // 这里就不用加public， 不然就重复了， 这里的name 和  salary就是传给super的参数
    // 加上public就相当于在这里 定义了一个 name 和 salary
    constructor(name: string, salary: number) {
        super(name, salary)
        this.reporters = []
    }
    addReporter(e: Employee_12_1) {
        this.reporters.push(e)
    }
}

const manager1 = new Manager('mary', 18000)
manager1.addReporter(emp_12_1)
console.log(manager1)


interface interfaceEmploy {
    name: string
    salary: number
    bonus?: number
}

//! 使用类来是实现接口(自动实现)
//* TypeScript中的接口不需要使用implement关键字， 因为它们是自动实现的
const empImpl = new Employee_12_1('mary', 18000)
const emp1_12: interfaceEmploy = empImpl // 自动实现这个接口 


//! 如何选择隐式实现还是显示实现
// service.ts
// 定义接口
interface Service {
    login(): void
    getTrips(): string
    getLic(): string
    startTrip(): void
    updateLic(lic: string): void
}

// 实现接口
class RPCService implements Service {
    login() {
        console.log('login')
    }
    getTrips() {
        return 'getTrips'
    }
    getLic() {
        return 'getLic'
    }
    startTrip() {
        console.log('startTrip')
    }
    updateLic(lic: string) {
        console.log(lic)
    }
}

// Login page
// file: login.ts

interface LoginService {
    login(): void
}

const page = {
    // 我作为一个使用者，我才知道我要什么，我只要Login，不需要其他那么多的方法
    //! 所以可以自定义一个 接口， 使用 as 的方式
    service: new RPCService() as LoginService,
    onLoginButtonClicked() {
        this.service.login()
    },
}
