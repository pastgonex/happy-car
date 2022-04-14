// pages/driving/driving.ts

import { routing } from "../../utils/routing"

const centPerSec = 0.7

// const updateIntervalSec = 5
const initialLat = 30
const initialLng = 120

function formatDuration(sec: number) {
  const padString = (n: number) => {
    return n < 10 ? '0' + n.toFixed(0) : n.toFixed(0)
  }
  const h = Math.floor(sec / 3600)
  sec -= 3600 * h
  const m = Math.floor(sec / 60)
  sec -= 60 * m
  const s = Math.floor(sec)
  return `${padString(h)}:${padString(m)}:${padString(s)}`
}

function formatFee(cents: number) {
  // 元
  return (cents / 100).toFixed(2)
}

Page({

  /**
   * Page initial data
   */
  timer: undefined as number | undefined,
  tripID: '',
  data: {
    location: {
      latitude: initialLat,
      longitude: initialLng,
    },
    scale: 12,
    elapsed: '00:00:00',
    fee: '0.00',
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: initialLat,
        longitude: initialLng,
        width: 20,
        height: 20,
      },
    ],
  },

  /**
   * Lifecycle function--Called when page load
   */
  //               这种写法， 定死了只能写 trip_id, 符合了接口的规定
  onLoad(opt: Record<'trip_id', string>) {
    const o: routing.DrivingOpts = opt
    console.log('current trip', o.trip_id)
    this.setupLocationUpdator()
    this.setupTimer()
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
    // 停止位置更新
    wx.stopLocationUpdate()
    if (this.timer) {
      clearInterval(this.timer)
    }
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

  setupLocationUpdator() {
    wx.startLocationUpdate({
      fail: console.error,

    })
    wx.onLocationChange(loc => {
      console.log('location', loc)
      this.setData({
        location: {
          latitude: loc.latitude,
          longitude: loc.longitude,
        }
      })
    })
  },

  setupTimer() {
    let elapsedSec = 0
    let cents = 0
    this.timer = setInterval(() => {
      elapsedSec++
      cents += centPerSec
      this.setData({
        elapsed: formatDuration(elapsedSec),
        fee: formatFee(cents),
      })
    }, 1000)
  },

  onEndTripTap() {
    wx.redirectTo({
      url: routing.mytrips(),
    })
  }
})