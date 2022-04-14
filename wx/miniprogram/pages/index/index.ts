import { routing } from "../../utils/routing"
import {IAppOption} from "../../appoption";

Page({
  isPageShowing: false,
  data: {
    userInfo: {},
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: "",
      layerStyle: -1,
      enableZoom: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      latitude: 31,
      longitude: 120,
    },
    scale: 10,
    markers: [{
      iconPath: "/resources/car.png",
      id: 0,
      latitude: 23.099994,
      longitude: 113.324520,
      width: 25,
      height: 25,
    },
    {
      iconPath: "/resources/car.png",
      id: 1,
      latitude: 23.099994,
      longitude: 114.324520,
      width: 25,
      height: 25,
    }],
  },
  async onLoad() {
    const UserInfo = await getApp<IAppOption>().globalData.userInfo
    this.setData({
      userInfo: UserInfo
    })
  },
  onShow() {

  },
  onHide() {
    this.isPageShowing = false
  },
  onMyLocationTap() {
    wx.getLocation({
      type: "gcj02",
      success: res => {
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          }
        })
      },
      fail: () => {
        wx.showToast({
          icon: "none",
          title: "请前往设置页授权"
        })
      }
    })
  },
  onScanTap() {
    wx.scanCode({
      // onlyFromCamera: false,
      success: () => {
        wx.showModal({
          title: "身份认证",
          content: "需要身份认证才能租车",
          success: (res) => {
            if (res.confirm) {
              // console.log(res)
              // todo get car id from scan result 
              const carID = 'car123'
              // const redirectURL = `/pages/unlock/unlock?car_id=${carID}`
              const redirectURL = routing.unlock({
                car_id: carID
              })
              wx.navigateTo({
                // url: `/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`,
                url: routing.register({
                  redirectURL: redirectURL
                })
              })
              // todo
            } else if (res.cancel) {
              console.log('用户点击取消')
            }

          },
          fail: () => {

          }
        })
      },
      fail: res => {
        console.log(res)
      }
    })
  },
  moveCars() {
    const map = wx.createMapContext("map") // map的id
    const dest = {
      // latitude: 23.099994,
      // longitude: 113.324520,
      latitude: this.data.markers[0].latitude,
      longitude: this.data.markers[0].longitude,
    }
    // 做成一个回调函数，每隔一段时间调用一次
    const moveCar = () => {
      dest.latitude += 0.1
      dest.longitude += 0.1
      map.translateMarker({
        destination: {
          latitude: dest.latitude,
          longitude: dest.longitude,
        },
        markerId: 0,
        autoRotate: false,
        rotate: 0,
        duration: 5000,
        animationEnd: () => {
          if (this.isPageShowing) {
            this.data.markers[0].latitude = dest.latitude
            this.data.markers[0].longitude = dest.longitude
            moveCar()
          } else {
            this.data.markers[0].latitude = dest.latitude
            this.data.markers[0].longitude = dest.longitude
          }
        },
      })
    }
    moveCar()
  },
  onMyTripsTap() {
    wx.navigateTo({
      // url: '/pages/mytrips/mytrips',
      url: routing.mytrips()
    })
  }
})   
