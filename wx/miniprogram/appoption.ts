export interface IAppOption {
  globalData: {
    // userInfo?: WechatMiniprogram.UserInfo,
    userInfo: Promise<WechatMiniprogram.UserInfo>,
  }
  userInfoReadyCallback?: WechatMiniprogram.GetUserInfoSuccessCallback,
  resolveUserInfo(userInfo: WechatMiniprogram.UserInfo): void,
}