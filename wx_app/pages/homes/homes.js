// pages/homes/homes.js
var app = getApp(); //取得全局App({..})实例
var util = require('../../utils/util.js') //引用外部文件，为了获取当前时间
const columns = [];

Page({

  /**
   * 页面的初始数据
   */
  data: {
    code: '',//孩子编码
    name: '',//孩子名字
    age: '',//孩子月龄
    manager_name: '', //负责人
    test_n: '',//第几次测试
    alltimes:'',//总次数
    // hidden: false,//是否隐藏
    addmodel: true,//是否添加模式
    label_testn:'',
    buttonvalue: '',
    title:'',
    allList:[],
  },


  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {

    let that = this
    console.log(options)
    that.setData({
      // hidden:(options.addmodel==="true"),
      addmodel:(options.addmodel==="true"),
    })
    if (options.addmodel ==="false") {//查看模式
      that.setData({
        label_testn:"总次数",
        title:"查看记录设置",
        buttonvalue:"查看评估表",
        code: app.globalData.detail_child_id,
        name: app.globalData.detail_name,
        manager_name: app.globalData.detail_manager_name,
        test_n: app.globalData.detail_times,
        alltimes: app.globalData.detail_times,
      })
    } else {
      that.setData({
        label_testn:"第几次评估",
        title:"新增记录设置",
        buttonvalue:"新增评估表",
      })

      console.log("getData")
      //再次获取记录表，方便后面查找编号对应的最大次数
      wx.request({
        url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/children/showAll',
        // url: 'https:///www.lprs.site/api/v1/children/showAll',
        data: {
        },
        header: {
          'content-type': 'application/json' // 默认值
        },
        method: 'get',
        success: (res) => {
          // console.log(res)
          that.setData({
            allList:res.data.data
          })
          // console.log(that.data.allList)
        },
        fail(){
          console.log("接口showall失败")
        },
      })
    }
    wx.setNavigationBarTitle({
      title: that.data.title,
  })
  },

  //新增/查看评估表
  nextLevelpage() {
    let that = this
    if (that.data.addmodel) {//添加新记录
      //通过首页查询想要新增的孩子的最大次数
      let maxtimes = 0
      for (var i = 0; i < that.data.allList.length; i++) {
        if (that.data.allList[i].code == that.data.code) {
          console.log("getData")
          maxtimes = that.data.allList[i].count_test
          break
        }
      }
      console.log("得到编号为"+that.data.code+"的最大次数为"+maxtimes)
      if (maxtimes+1 < that.data.test_n || that.data.code == '' || that.data.name == '' || that.data.age == '' || that.data.manager_name == '' || that.data.test_n == '') {
        let content = '缺少信息'
        if (maxtimes != that.data.test_n-1) {
          content = '之前一共做过'+maxtimes+'次评估，请输入正确数字'
          that.setData({
            test_n:maxtimes+1,
          })
        }
        wx.showModal({   //添加确认提示
          title: '提示',
          content: content,
          showCancel:false,
        })
        return
      } else {
        wx.showModal({   //添加确认提示
          title: '提示',
          content: '确认添加新记录吗',
          success (res) {
            if (res.confirm) {   //确认则跳转
              console.log('用户点击确定')
              app.globalData.code = that.data.code
              app.globalData.name = that.data.name
              app.globalData.age = that.data.age
              app.globalData.manager_name = that.data.manager_name
              app.globalData.test_n = that.data.test_n
              wx.navigateTo({
                url: '/pages/Alevel2/Alevel2?addmodel='+that.data.addmodel,
                // success: function () {
                //   let pages = getCurrentPages();
                //   let beforePage = pages[pages.length - 2];
                //   beforePage.changeData();   
                //  }
              })
            } else if (res.cancel) {  //取消
              console.log('用户点击取消')
            }
          }
        })
      }
    } else {//查看记录
      app.globalData.detail_child_id = that.data.code
      wx.navigateTo({
        url: '/pages/Alevel2/Alevel2?addmodel='+that.data.addmodel+'&alltimes='+that.data.alltimes,
        // success: function () {
        //   let pages = getCurrentPages();
        //   let beforePage = pages[pages.length - 2];
        //   beforePage.changeData();   
        //   }
      })
    }
  },

  codeChange() {
    let that = this
    let maxtimes = 0
    for (var i = 0; i < that.data.allList.length; i++) {
      if (that.data.allList[i].code == that.data.code) {
        console.log("getData")
        maxtimes = that.data.allList[i].count_test
        break
      }
    }
    that.setData({
      test_n : maxtimes+1,
    })
    console.log("得到编号为"+that.data.code+"的最大次数为"+maxtimes)
  },

  suggestion() {
    wx.navigateTo({
      url: '/pages/suggestion/suggestion?alltimes='+this.data.alltimes,
    })
  },
  graph() {
    wx.navigateTo({
      url: '/pages/graph/graph?child_id='+this.data.code,
    })
  },
  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

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