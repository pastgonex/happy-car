// pages/register/register.ts
Page({

  /**
   * Page initial data
   */
  data: {
    genders: ['未知', '男', '女', '其他'],
    genderIndex: 0,
    // licImgURL: undefined as string | undefined,
    licImgURL: '',
    birthDate: '1990-01-01',
    licNo: '',
    name: '',
    state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED',
  },

  /**
   * Lifecycle function--Called when page load
   */
  onLoad() {

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
  onUploadLic() {
    // console.log("on clicked")
    wx.chooseImage({
      success: res => {
        // console.log(res)
        if (res.tempFilePaths.length > 0) {
          this.setData({
            licImgURL: res.tempFilePaths[0]
          })
          // Todo upload image  模拟上传照片之后的操作 （自动识别驾照中的信息）
          setTimeout(() => {
            this.setData({
              licNo: '12345678901234567890',
              name: 'Ever Ni',
              genderIndex: 1,
              birthDate: '20000-08-05'
            })
          })
        }
      }
    })
  },
  // 性别选择后的结果
  onGenderChange(e: any) {
    this.setData({
      genderIndex: e.detail.value,

    })
  },
  
  // 日期选择后的结果
  onBirthDateChange(e: any) {
    this.setData({
      birthDate: e.detail.value,
    })
  },

  // 递交审查
  onSubmit() {
    // TODO: submit the form to server
    this.setData({
      state: 'PENDING'
    })
    // TODO: 模拟服务器审查
    setTimeout(() => {
      this.onLicVerified()
    }, 3000)
  },
  onResubmit() {
    this.setData({
      // 先清掉数据
      state: 'UNSUBMITTED',
      licImgURL: '',
    })
  },
  onLicVerified() {
    this.setData({
      state: 'VERIFIED',
    }),
    // 审查结束之后， 跳转到开锁页面
    wx.redirectTo({
      url: '/pages/unlock/unlock',
    })
  }
})