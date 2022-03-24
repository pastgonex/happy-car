// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  data: {
    motto: 'Hello World from TypeScript',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    canIUseGetUserProfile: false,
    canIUseOpenData: wx.canIUse('open-data.type.userAvatarUrl') && wx.canIUse('open-data.type.userNickName') // 如需尝试获取用户信息可改为false
  },
  // 事件处理函数
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs',
    })
  },
  async onLoad() {
    // @ts-ignore
    // if (wx.getUserProfile) {
    //   this.setData({
    //     canIUseGetUserProfile: true
    //   })
    // }
    const  userInfo = await app.globalData.userInfo
    this.setData({
      // userInfo: app.globalData.userInfo,
      userInfo,
      hasUserInfo: true,
    })
    //* onLoad的时候等userInfo，等到了就（then）把界面更新
    //* 上面的改写下面的代码
    // app.globalData.userInfo.then(userInfo => {
    //   this.setData({
    //     // userInfo: app.globalData.userInfo,
    //     userInfo,
    //     hasUserInfo: true,
    //   })
    // })
  },
  getUserProfile() {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }
    })
  },
  getUserInfo(e: any) {
    // 不推荐使用getUserInfo获取用户信息，预计自2021年4月13日起，getUserInfo将不再弹出弹窗，并直接返回匿名的用户个人信息
    console.log(e)
    // app.globalData.userInfo = e.detail.userInfo 这个是错误的
    const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo // 强行受到编译器保护
    app.globalData.userInfo = Promise.resolve(userInfo)
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  }
})
