import { routing } from "../../utils/routing"

interface Trip {
  id: string,
  shortId: string,
  start: string,
  end: string,
  duration: string,
  fee: string,
  distance: string,
  status: string,
  inProgress: boolean
}

interface MainItem {
  id: string,
  navId: string,
  navScrollId: string
  data: Trip,
}

interface NavItem {
  id: string
  mainId: string,
  label: string,
}

interface MainItemQueryResult {
  id: string
  top: number
  dataset: {
    navId: string
    navScrollId: string
  }
}

// pages/mytrips/mytrips.ts
Page({
  scrollStatus: {
    mainItems: [] as MainItemQueryResult[]
  },

  /**
   * Page initial data
   */
  data: {
    canIUseGetUserProfile: false,
    hasUserInfo: false,
    indicatorDots: true, //是否显示面板指示点
    autoPlay: true, //是否自动切换
    interval: 3000,//自动切换时间间隔
    duration: 500, //滑动动画时长
    circular: true, //是否采用衔接滑动
    multiItemCount: 1, //每个页面的数据
    prevMargin: '',//上一个页面的margin-left
    nextMargin: '', //下一个页面的margin-right
    vertical: false, //是否竖向滑动
    current: 0, //当前页面
    avatarUrl: '',
    navSel: '',
    promotionItems: [
      {
        img: 'https://images.unsplash.com/photo-1551334787-21e6bd3ab135?w=640',
        promotionID: 1,
      },
      {
        img: 'https://images.unsplash.com/photo-1551214012-84f95e060dee?w=640',
        promotionID: 2,
      },
      {
        img: 'https://images.unsplash.com/photo-1551446591-142875a901a1?w=640',
        promotionID: 3,
      },
    ],
    trips: [] as Trip[],
    tripsHeight: 0,
    mainScroll: '',
    mainItems: [] as MainItem[],
    navItem: [] as NavItem[],
    navCount: 0,
  },

  /**
   * Lifecycle function--Called when page load
   */
  onLoad() {

    this.populateTrips()
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
  populateTrips() {
    // const trips: Trip[] = [] // 限定类型， 做保护
    const mainItems: MainItem[] = []
    const navItems: NavItem[] = []
    let navSel = ''
    let prevNav = ''
    for (let i = 0; i < 100; i++) {
      const mainId = 'main-' + i
      const navId = 'nav-' + i
      const tripId = (10001 + i).toString()
      const shortId = tripId
      if (!prevNav) {
        prevNav = navId
      }
      mainItems.push({
        id: mainId,
        navId: navId,
        navScrollId: prevNav,
        data: {
          id: tripId,
          shortId: '****' + shortId,
          start: '东方明珠',
          end: '迪士尼',
          distance: '27.0公里',
          duration: '0时44分',
          fee: '128.00元',
          status: "已完成",
          inProgress: false,

        }
      })
      navItems.push({
        id: navId,
        mainId: mainId,
        label: tripId,
      })
      if (i === 0) {
        navSel = navId
      }
      prevNav = navId
    }
    this.setData({
      // trips: trips, // trips, 这样写， 就不用写 trips: trips 
      mainItems,
      navItems,
      navSel,
    }, () => {
      this.prepareScrollStates()
    })
  },

  prepareScrollStates() {
    wx.createSelectorQuery().selectAll('.main-item')
      .fields({
        id: true,
        dataset: true,
        rect: true,
      }).exec(res => {
        this.scrollStatus.mainItems = res[0]
      })
  },

  /**
   * Lifecycle function--Called when page is initially rendered
   */
  // 渲染完成
  onReady() {
    // 获取 id为heading的元素
    wx.createSelectorQuery().select('#heading')
      .boundingClientRect(rect => {
        const height = wx.getSystemInfoSync().windowHeight - rect.height
        this.setData({
          tripsHeight: height,
          navCount: Math.round(height / 50)
        })
      }).exec() // 执行
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
  onSwiperChange(e: any) {
    // console.log(e)
    // 自动变的
    if (!e.detail.source) {
      // caused by our program
      return
    }
    // todo process
  },
  onPromotionItemTap(e: any) {
    // console.log(e)
    const promotionID = e.currentTarget.dataset.promotionid
    if (promotionID) {
      // claim this promotion 
    }
  },
  onRegisterTap() {
    wx.navigateTo({
      // url: '/pages/register/register',
      url: routing.register(),
    })
  },

  // TODO  
  getUserProfile() {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认
    // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '用于实时展示用户头像', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res: any) => {
        this.data.avatarUrl = res.avatarUrl
        this.data.hasUserInfo = true
      }
    })
  },

  onNavItemTap(e: any) {
    const mainId: string = e.currentTarget?.dataset?.mainId
    const navId: string = e.currentTarget?.id
    if (mainId) {
      this.setData({
        mainScroll: mainId,
        navSel: navId,
      })
    }
  },

  onMainScroll(e: any) {
    // console.log(e)
    const top: number = e.currentTarget?.offsetTop + e.detail.scrollTop
    if (top === undefined) {
      return
    }
    const selItem = this.scrollStatus.mainItems.find(v => {
      return v.top >= top
    })
    if (!selItem) return
    this.setData({
      navSel: selItem.dataset.navId,
      navScroll: selItem.dataset.navScrollId
    })
  },
})