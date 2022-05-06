export interface IAppOption {
  globalData: {
    // userInfo?: WechatMiniprogram.UserInfo,
    userInfo: Promise<WechatMiniprogram.UserInfo>,
  }
  resolveUserInfo(userInfo: WechatMiniprogram.UserInfo): void,
}