var app = getApp(); //取得全局App({..})实例
// 1. 获取数据库引用

Page({

  /**
   * 页面的初始数据
   */
  data: {
    faultList: [],
    datakey: "",
    partModel: "",
    partremark: "",
    updateTime: "",
  },
  /**
   * 生命周期函数--监听页面加载
   */

  onLoad: function (options) {
    console.log(app.globalData);
    this.setData({
      datakey: app.globalData.partdata,
      maintenancePart: app.globalData.maintenancePart,
      partdata: app.globalData.partdata,
      partModel: app.globalData.partModel,
      partremark: app.globalData.partremark,
      standardImgUrl: app.globalData.standardImgUrl,
      troubleImgUrl: app.globalData.troubleImgUrl,
      nextInspectionTime: app.globalData.nextInspectionTime,
      radio1: app.globalData.isNormal,
      radio2: app.globalData.isNormal,
      activeNames: app.globalData.nextInspectionTime,
      standardFlag: false,
      troubleFlag: false,
      resultImgUrl: app.globalData.resultImgUrl
    })
    if (app.globalData.standardImgUrl != null && app.globalData.standardImgUrl != '') {
      this.setData({
        standardList: app.globalData.standardImgUrl.split(','),
        standardFlag: true
      })
    }
    if (app.globalData.troubleImgUrl != null && app.globalData.troubleImgUrl != '') {
      this.setData({
        troubleList: app.globalData.troubleImgUrl.split(','),
        troubleFlag: true
      })
    }
    if (app.globalData.resultImgUrl != null && app.globalData.resultImgUrl != '') {
      this.setData({
        resultFlag: true
      })
    }
  },
  btn5: function () {
    wx.switchTab({

      url: '/pages/B1/B1',

    })
  },
  previewImg(e) {
    var url = e.target.dataset.url
    wx.previewImage({
      urls: [url],
    })
  }

})