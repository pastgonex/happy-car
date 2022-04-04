// 这里一定要保证正确
export namespace routing {
    export interface DrivingOpts {
        trip_id: string
    }

    // 使用方式: routing.driving()
    export function driving(o: DrivingOpts) {
        return `/pages/driving/driving?trip_id=${o.trip_id}`
    }

    export interface LockOpts {
        car_id: string
    }

    export function unlock(o: LockOpts) {
        return `/pages/unlock/unlock?car_id=${o.car_id}`
    }


    // 使用Typescript作为限制
    export interface RegisterOpts {
        redirect?: string
    }

    export interface  RegisterParams {
        redirectURL: string
    }

    export function register(p?: RegisterParams) {
        const page = '/pages/register/register'
        if (!p) {
            return page
        }
        return `${page}?redirect=${encodeURIComponent(p.redirectURL)}`
    }

    export function mytrips() {
        return '/pages/mytrips/mytrips'
    }

}