// import {getSetting, getUserInfo} from "./utils/wxapi"
// import {IAppOption} from "./appoption";
// import {Happycar} from "./service/request";
// // import { happycar } from "./service/proto_gen/trip_pb";
// // import camelcaseKeys from "camelcase-keys";
//
// let resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void
// let rejectUserInfo: (reason?: any) => void
//
// // app.ts
// App<IAppOption>({
//     globalData: {
//         //! 获取用户信息 promise版本
//         // 只用这一个， 就可以应对各种各样的情况了
//         userInfo: new Promise((resolve, reject) => {
//             resolveUserInfo = resolve
//             rejectUserInfo = reject
//         })
//     },
//     async onLaunch() {
//         // wx.request({
//         //   url: 'http://localhost:8080/trip/trip123',
//         //   method: 'GET',
//         //   success: res => {
//         //     const getTripRes = happycar.GetTripResponse.fromObject(camelcaseKeys(res.data as object, {
//         //       deep: true, // 对象的所有层级， 所有的key都改成驼峰
//         //     }))
//         //     // const getTripRes = happycar.GetTripResponse.fromObject(res.data as object)
//         //     console.log(getTripRes)
//         //     console.log('status is', happycar.TripStatus[getTripRes.trip?.status!])
//         //   },
//         //   fail: console.error
//         // })
//
//         // 展示本地存储能力
//         // const logs = wx.getStorageSync('logs') || []
//         // logs.unshift(Date.now())
//         // wx.setStorageSync('logs', logs)
//
//         // 登录 静默登录
//         // wx.login({
//         //     success: res => {
//         //         // 发送 res.code 到后台换取 openId, sessionKey, unionId
//         //         wx.request({
//         //             url: 'http://localhost:8080/v1/auth/login',
//         //             method: 'POST',
//         //             data: {
//         //                 code: res.code
//         //             } as auth.v1.ILoginRequest,
//         //             success: res => {
//         //                 const loginResponse: auth.v1.ILoginResponse =
//         //                     auth.v1.LoginResponse.fromObject(
//         //                         camelcaseKeys(res.data as object, {deep: true})
//         //                     )
//         //                 console.log(loginResponse)
//         //                 wx.request({
//         //                     url: 'http://localhost:8080/v1/trip',
//         //                     method: 'POST',
//         //                     data: {
//         //                         start: "abc",
//         //                     } as rental.v1.ICreateTripRequest,
//         //                     header: {
//         //                         // 业界规范，Bearer加上token
//         //                         authorization: 'Bearer ' + loginResponse.accessToken,
//         //                     },
//         //                     fail: res => {
//         //                         console.log(res)
//         //                     }
//         //                 })
//         //             },
//         //             fail: res => {
//         //                 console.error(res)
//         //             }
//         //         })
//         //     },
//         // })
//         //* 获取用户信息
//
//         // 登录
//         await Happycar.login()
//         try {
//             const setting = await getSetting()
//             if (setting.authSetting['scope.userInfo']) {
//                 const userInfoRes = await getUserInfo()
//                 resolveUserInfo(userInfoRes.userInfo)
//             }
//         } catch (err) {
//             rejectUserInfo(err)
//         }
//         //* 上面的代码是将下面的改写
//         // getSetting().then(res => {
//         //   if (res.authSetting['scope.userInfo']) {
//         //     return getUserInfo()
//         //   }
//         //   return undefined
//         // }).then(res => {
//         //   if (!res) {
//         //     return
//         //   }
//         //   // 通知页面我获得了用户信息
//         //   resolveUserInfo(res.userInfo)
//         // }).catch(err => rejectUserInfo(err))
//         // 获取用户信息
//         // wx.getSetting({
//         //   success: res => {
//         //     if (res.authSetting['scope.userInfo']) {
//         //       // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
//         //       wx.getUserInfo({
//         //         success: res => {
//         //           // 可以将 res 发送给后台解码出 unionId
//         //           this.globalData.userInfo = res.userInfo
//
//         //           // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
//         //           // 所以此处加入 callback 以防止这种情况
//         //           //通知页面我获得了用户信息
//         //           if (this.userInfoReadyCallback) {
//         //             this.userInfoReadyCallback(res)
//         //           }
//         //         },
//         //       })
//         //     }
//         //   }
//         // }),
//     },
//     resolveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
//         // this.resolveUserInfo才是递归调用
//         // 单单一个resolveUserInfo是全局的
//         // 不用担心重名
//         resolveUserInfo(userInfo)
//     }
// })
import { IAppOption } from "./appoption"
import { Happycar } from "./service/request"
import { getSetting, getUserInfo } from "./utils/wxapi"

let resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void
let rejectUserInfo: (reason?: any) => void

// app.ts
App<IAppOption>({
    globalData: {
        userInfo: new Promise((resolve, reject) => {
            resolveUserInfo = resolve
            rejectUserInfo = reject
        })
    },
    async onLaunch() {
        // 登录
        await Happycar.login()

        // 获取用户信息
        try {
            const setting = await getSetting()
            if (setting.authSetting['scope.userInfo']) {
                const userInfoRes = await getUserInfo()
                resolveUserInfo(userInfoRes.userInfo)
            }
        } catch (err) {
            rejectUserInfo(err)
        }
    },
    resolveUserInfo(userInfo: WechatMiniprogram.UserInfo) {
        resolveUserInfo(userInfo)
    }
})