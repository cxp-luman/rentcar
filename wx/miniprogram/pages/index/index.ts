Page({
  data: {
    avatarURL: '',
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      latitude: 41.826363,
      longitude: 123.571858,
    },
    scale: 16,
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 23.099994,
        longitude:114.32452,
        width: 50,
        height: 50
      },
      {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 23.099994,
        longitude:113.32452,
        width: 50,
        height: 50
      }
    ]
  },
  onMyLocationTap() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        })
      }, 
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '请前往设置页授权',
        })
      }
    })
  },
  onScanClicked() {
    wx.scanCode({
      success: () => {
        wx.navigateTo({
          url: '/pages/register/register',
        })
      },
      fail: console.error,
    })
  }
 
  
})