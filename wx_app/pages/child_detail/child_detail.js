// pages/child_detail/child_detail.js
var app = getApp(); //取得全局App({..})实例
const num2options = new Map([
  [1, '儿童感知觉发展能力'],
  [2, '儿童粗大动作能力'],
  [3, '儿童精细动作能力'],
  [4, '儿童语言与沟通发展能力'],
  [5, '儿童认知发展能力'],
  [6, '儿童社会交往发展能力'],
  [7, '儿童生活自理发展能力'],
]);
Page({

  /**
   * 页面的初始数据
   */
  data: {
    options: [
      {
        text: '1.儿童感知觉发展能力',
        value: 1
      },
      {
        text: '2.儿童粗大动作能力',
        value: 2
      },
      {
        text: '3.儿童精细动作能力',
        value: 3
      },
      {
        text: '4.儿童语言与沟通发展能力',
        value: 4
      },
      {
        text: '5.儿童认知发展能力',
        value: 5
      },
      {
        text: '6.儿童社会交往发展能力',
        value: 6
      },
      {
        text: '7.儿童生活自理发展能力',
        value: 7
      },
      {
        text: '8.训练建议',
        value: 8
      },
      {
        text: '9.孤独症儿童发展情况剖面图',
        value: 9
      }
    ], //内容下拉菜单
    optionvalue: 1, //初始选择
    times: [
      {
        text: '1',
        value: 1
      },
      {
        text: '2',
        value: 2
      },
      {
        text: '3',
        value: 3
      },
      {
        text: '4',
        value: 4
      }
    ], //次数下拉菜单
    timesvalue: 1, //初始选择
    flag: false, //选择
    child_id:'',
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
   let that = this
   that.setData({
     child_id: app.globalData.detail_child_id
   })
  },
  
  //查看内容改变
  menuChange(e) {
    let that = this
    var temp = !that.data.flag
    that.setData({
      optionvalue:e.detail,
      flag: temp,
    })
    console.log(that.data.optionvalue)
  },
  //次数选择改变
  timeChange(e) {
   
    let that = this
    var temp = !that.data.flag
    that.setData({
      timesvalue:e.detail,
      flag: temp,
    })
    console.log(that.data.timesvalue)
  },
  //返回刷新
  changeData:function(){
    this.onLoad();
  },

  btn1: function (e) {
    let that = this
    app.globalData.detail_optionvalue = that.data.optionvalue
    app.globalData.detail_timesvalue = that.data.timesvalue
    if (that.data.optionvalue == 9) {//跳转到剖面图
      wx.navigateTo({
        url: '/pages/E7/E7',
      })
    } else if (that.data.optionvalue == 8) {//跳转到训练指标
      wx.navigateTo({
        url: '/pages/suggestion/suggestion',
      })
    } else {//跳转到测试表结果
      wx.navigateTo({//把当前表的名字传给下一页作为题目
        url: '/pages/result/result?title='+num2options.get(that.data.optionvalue),
      })
    }
    
  },


//刷新
onRefresh(){
 
},
//网络请求，获取数据
getData(){
  
},
/**
* 页面相关事件处理函数--监听用户下拉动作
*/
onPullDownRefresh: function () {
  //调用刷新时将执行的方法
  this.onRefresh();
}
})