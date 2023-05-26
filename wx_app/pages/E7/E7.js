// pages/homes/homes.js
var app = getApp(); //取得全局App({..})实例
var util = require('../../utils/util.js') //引用外部文件，为了获取当前时间
Page({

  /**
   * 页面的初始数据
   */
  data: {
    workNum: '',//工单号
    maintenanceStartTime: '', //维保开始时间
    liftNum: '', //电梯编号
    liftBrand: '', //电梯品牌
    manager: '', //负责人
    time: '', //用时
    maintenanceStartTime:'',
    readonly:'',
    typeName:'',
    flag1:"已提交",
    flag2:"未提交"
  },

  //跳转到二级页面,接口上传一级表
  nextLevelpage() {
    let that = this
    // wx.navigateTo({
    //   url: '/pages/facerecognition/index',
    // })
    if (app.globalData.issubmit=='已提交') {
      wx.navigateTo({
        url: '/pages/Alevel3/Alevel3',
      })
    }else{
      wx.navigateTo({
      url: '/pages/Alevel2/Alevel2',
    })
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    this.setData({
      readonly:app.globalData.issubmit
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {
    this.setData({
      elevatorname: app.globalData.entryname,
      liftNum:app.globalData.elevatorNumber,
      liftBrand:app.globalData.brand,
      maintenanceStartTime:app.globalData.maintenanceStartTime,
      time:app.globalData.useTime,
      manager:app.globalData.responsiblePerson,
      typeName:app.globalData.typeName,
    })
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  },

})