import { routing } from "../../utils/routing"

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
  onLoad(opt: Record<'car_id', string>) {
    const o: routing.LockOpts = opt
    console.log('unlocking car', o.car_id)
    if (wx.getUserProfile) {
      this.setData({
        canIUseGetUserProfile: true,
      })
    }
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

  getUserProfile(e) {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认
    // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '用于实时展示用户头像', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
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
        //todo  模拟后端数据
        console.log('starting a trip', {
          location: {
            latitude: loc.latitude,
            longituge: loc.longitude,
          },
          // todo 需要和后端双向绑定（shareLocation）
          avatarUrl: this.data.shareLocation ? this.data.avatarUrl : '',
          carID: 0
        })
        const tripID = 'trip456'
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