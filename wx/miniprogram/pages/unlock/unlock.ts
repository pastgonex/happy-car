import {routing} from "../../utils/routing"
import {IAppOption} from "../../appoption";
import {TripService} from "../../service/trip";

// pages/unlock/unlock.ts
const shareLocationKey = "share_location"
const shareUserInfo = "share_userInfo"

Page({

    /**
     * Page initial data
     */
    data: {
        userInfo: {},
        avatarUrl: '',
        shareLocation: false,
        hasUserInfo: false,
        canIUseGetUserProfile: false,
    },

    /**
     * Lifecycle function--Called when page load
     */
    async onLoad(opt: Record<'car_id', string>) {
        const o: routing.LockOpts = opt
        console.log('unlocking car', o.car_id)
        // if (wx.getUserProfile) {
        //   this.setData({
        //     canIUseGetUserProfile: true,
        //   })
        // }

        // wx.getUserProfile
        // this.setData({
        //   canIUseGetUserProfile: true,
        // })
        wx.getUserProfile({
            lang: 'zh_CN',
            desc: '用户登录',
            success: () => {
            },
            // 失败回调
            fail: (res) => {
                // 弹出错误
                console.log(res)
            }
        })
        this.setData({
            canIUseGetUserProfile: true,
        })
    },

    /**
     * Lifecycle function--Called when page is initially rendered
     */
    onReady() {

    },

    /**
     * Lifecycle function--Called when page show
     */
    onShow() {

    },

    /**
     * Lifecycle function--Called when page hide
     */
    onHide() {

    },

    /**
     * Lifecycle function--Called when page unload
     */
    onUnload() {

    },

    /**
     * Page event handler function--Called when user drop down
     */
    onPullDownRefresh() {

    },

    /**
     * Called when page reach bottom
     */
    onReachBottom() {

    },

    /**
     * Called when user click on the top right corner to share
     */
    onShareAppMessage() {

    },

    getUserProfile() {
        // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认
        // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
        wx.getUserProfile({
            desc: '用于实时展示用户头像', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
            success: (res: any) => {
                getApp<IAppOption>().resolveUserInfo(res.userInfo)
                this.setData({
                    avatarUrl: res.userInfo.avatarUrl,
                    userInfo: res.userInfo,
                    hasUserInfo: true,
                    shareLocation: wx.getStorageSync(shareUserInfo) || false,
                })
            }
        })
    },

    // 2021.4.28日之后不推荐使用此接口
    //   onGetUserInfo(e: any) {
    //     const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
    //     if (userInfo) {
    //         getApp<IAppOption>().resolveUserInfo(userInfo)
    //         this.setData({
    //             shareLocation: true,
    //         })
    //         wx.setStorageSync(shareLocationKey, true)
    //     }
    // },

    onShareLocation(e: any) {
        const shareLocation: boolean = e.detail.value
        // 记住用户的选择
        wx.setStorageSync(shareLocationKey, shareLocation)
    },
    onUnlockTap() {
        wx.getLocation({
            type: 'gcj02',
            success: (loc) => {
                //TODO  send request to server to start a trip
                // 协议 http
                // 服务器地址 api.happycar.cn
                // 路径 /trip
                // 参数
                // 数据类型
                // 数据编码 JSON
                // 安全性 header 带上token
                // 错误处理
                //! starting a trip
                TripService.CreateTrip({
                    start: 'abc',
                })
                return // 先return

                wx.request({
                    url: 'https://api.happycar.cn/trip',
                    dataType: 'json',
                    data: {
                        location: {
                            latitude: loc.latitude,
                            longitude: loc.longitude,
                        },
                        avatarUrl: this.data.shareLocation ? this.data.avatarUrl : '',
                    },
                    header: {
                        authorization: 'jf32i9r89h' // toke
                    },
                    method: 'POST',
                    responseType: 'text', // 和后端两个人共同商讨
                    success: (res) => {
                        if (res.statusCode === 200) {
                            // const tripID = res.data.tripID
                            const tripID = '123'
                            wx.showLoading({
                                title: '开锁中',
                                mask: true, // 让页面不能点击
                            })
                            setTimeout(() => {
                                wx.redirectTo({
                                    //url: `/pages/driving/driving?trip_id=${tripID}`,
                                    url: routing.driving({
                                        trip_id: tripID,
                                    }),
                                    complete: () => {
                                        wx.hideLoading()
                                    }
                                })
                            }, 2000)
                        }
                    }
                })
            },
            fail: () => {
                wx.showToast({
                    icon: "none",
                    title: "请前往设置页授权位置信息"
                })
            }
        })
    }
})