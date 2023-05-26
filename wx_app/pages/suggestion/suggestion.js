// 1. 获取数据库引用
var app = getApp(); //取得全局App({..})实例
var output = [];
const data = "仰卧时，头转向一侧，同侧上下肢伸白，对侧上下肤屈曲;俯卧时，头转向一侧，脸侧贴在床上，手臂两膝关节均巨由;托起要儿背颈部，双腿有陵腿现象;仰卧时，头转向一侧，同侧上下肢伸白，对侧上下肤屈曲;俯卧时，头转向一侧，脸侧贴在床上，手臂两膝关节均巨由;托起要儿背颈部，双腿有陵腿现象;仰卧时，头转向一侧，同侧上下肢伸白，对侧上下肤屈曲;俯卧时，头转向一侧，脸侧贴在床上，手臂两膝关节均巨由;托起要儿背颈部，双腿有陵腿现象";

Page({
  data: {
    times: [], //次数下拉菜单
    timesvalue: 1, //初始选择
    alltimes:'',
    isScroll: true, //主页面是否滚动
    showModal: false, //弹窗
    showStr:"",
    showList: [],
    flag: false, //选择
    hidechoose:true,
  },
  onLoad: function (options) {
    let that = this
    
    var newtimes = []
    for (var i = 1; i <= options.alltimes; i++) {
      newtimes.push(
        {
          text: String(i),
          value: i,
        }
      )
    }
    that.setData({
      times:newtimes,
      alltimes:options.alltimes,
      timesvalue:options.alltimes,
    })
    // that.getbaseData() 
  },
  getbaseData(e) {
    let that = this
    wx.request({
      url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/children/suggestion',
      // url: 'https://www.lprs.site/api/v1/children/suggestion',
      header: {
        "Content-Type": "application/json"
      },
      method: 'get',
      data: {
        id: Number(app.globalData.detail_child_id),//孩子id
        test_n: that.data.timesvalue,//次数
      },
      success: res => {
        if (res.data.msg != "sucess"||res.data.data == "") {
          wx.showModal({ 
            title: '提示',
            content: '当前选择暂无建议，可能是无建议或未保存评估表。',
            showCancel:false,
          })
          
          that.setData({
            showList: [],
          })
        } else {
          var temp = res.data.data.split("。")
          temp.pop()
          that.setData({
            showList: temp,
          })
        }
      },
      fail(){
        wx.showModal({
          title: '提示',
          content: '数据请求失败',
          showCancel:false,
        })
      },
    })
  },
  menuChange(){
    let that = this
    that.setData({
      hidechoose:false,
    })
    that.getbaseData() 
  },
  //返回刷新
  changeData: function () {
    // this.onLoad();
  },

  onShow() {
    // this.onLoad();
  },
  /**
   * 弹出框蒙层截断touchmove事件
   */
  preventTouchMove: function () {},
  /**
   * 隐藏模态对话框
   */
  hideModal: function () {
    this.setData({
      showModal: false
    });
  },
  /**
   * 对话框取消按钮点击事件
   */
  onCancel: function () {
    this.setData({

      isScroll: true
    });
    this.hideModal();
  },
})