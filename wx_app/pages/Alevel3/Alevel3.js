var app=getApp()
Page({
  data: {
    maintenanceList: [],
    memlist2: [], //列表数据
    searchValue: '', //input框输入内容
  },
  //页面加载就会触发
  onLoad: function (options) {
    let that = this
    wx.request({
      url: 'https://jimucloud.cn/basedata/part/select',
      data: {
        pageNum: 0,
        pageSize: 0, 
        maintenanceOrder:app.globalData.workNumber, 
        openId:app.globalData.openid,
      },
      header: {
        'content-type': 'application/json' // 默认值
      },
      method: 'POST',
      success: (res) => {
        console.log(res);
        that.setData({
          maintenanceList: res.data.data.list
        })
        for (var i = 0; i < that.data.maintenanceList.length; i++) {
          that.data.maintenanceList[i].index = i
        }
        that.setData({
          maintenanceList: that.data.maintenanceList
        })
        console.log(that.data.maintenanceList.length);
        if (that.data.maintenanceList.length==0) {
          wx.showToast({
            title: '暂未填写内容',
            icon: 'error',
            duration: 2000, 
          })
        }
      }
    })
  },

  bind1: function (e) {
    wx.navigateTo({
      url: '/pages/Alevel4/Alevel4',
    })
    app.globalData.partdata = this.data.maintenanceList[e.currentTarget.id].partData
    app.globalData.isNormal = this.data.maintenanceList[e.currentTarget.id].isNormal
    app.globalData.partModel = this.data.maintenanceList[e.currentTarget.id].partModel
    app.globalData.partremark = this.data.maintenanceList[e.currentTarget.id].remark
    app.globalData.standardImgUrl = this.data.maintenanceList[e.currentTarget.id].standardImgUrl
    app.globalData.troubleImgUrl = this.data.maintenanceList[e.currentTarget.id].troubleImgUrl
    app.globalData.updateBy = this.data.maintenanceList[e.currentTarget.id].updateBy
    app.globalData.nextInspectionTime = this.data.maintenanceList[e.currentTarget.id].nextInspectionTime
    app.globalData.function = this.data.maintenanceList[e.currentTarget.id].partFunction
    app.globalData.maintenancePart = this.data.maintenanceList[e.currentTarget.id].maintenancePart
    app.globalData.updateTime = this.data.maintenanceList[e.currentTarget.id].updateTime
    app.globalData.resultImgUrl = this.data.maintenanceList[e.currentTarget.id].resultImgUrl

    

  },




})