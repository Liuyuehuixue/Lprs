// pages/child/child.js
var app = getApp(); //取得全局App({..})实例
const data = [
  {"name":"张一","child_id":1,"times":1,"manager_name":"张六"},
  {"name":"张二","child_id":2,"times":1,"manager_name":"张六"},
  {"name":"张三","child_id":3,"times":1,"manager_name":"张六"},
  {"name":"张四","child_id":4,"times":1,"manager_name":"张六"},
  {"name":"张五","child_id":5,"times":1,"manager_name":"张六"},
]
Page({

  /**
   * 页面的初始数据
   */
  data: {
    focus:false,
    searchValue: '', //input框输入内容
    NumberList: [],
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
    
    let that = this
    that.getData();
  },
  //返回刷新
  changeData:function(){
    this.onLoad();
  },
  focusHandler(e){
    this.setData({focus:true});
  },
  cancelHandler(e)
  {
    this.setData({
      focus:false,
      searchValue:""
    });
    this.onLoad();
  },
  // 搜索框
  query(e){
    // console.log("query in");
    this.setData({
      searchValue: e.detail.value
    })  
    //首先回显输入的字符串
    //实现搜索的功能
    var list = this.data.NumberList;		//先把第二条json存起来
    // console.log(list)
    var newlist = [];		//定义一个数组
    //循环去取数据
    for (var i = 0; i < list.length; i++) {
      var childname = list[i].name;
      var managername = list[i].manager_name;
      if (childname.indexOf(e.detail.value) >= 0 || managername.indexOf(e.detail.value) >= 0 ) {
        newlist.push(list[i])
      }
    }
    // console.log(newlist)
    //到这里list2就已经是你查出的数据
    //如果输入的关键词为空就加载原来的全部数据，不是空就加载搜索到的数据
    if(e.detail.value == ""){
      //加载全部
      this.onLoad()
    } else {
      this.setData({
        NumberList: newlist
      })
    }
  },

  detail: function (e) {
    let that = this
    // console.log(that.data.NumberList[e.currentTarget.dataset.index])
    // console.log(e.currentTarget.dataset.index)
    app.globalData.detail_child_id = that.data.NumberList[e.currentTarget.dataset.index].code
    app.globalData.detail_name = that.data.NumberList[e.currentTarget.dataset.index].name
    app.globalData.detail_times = that.data.NumberList[e.currentTarget.dataset.index].count_test
    app.globalData.detail_manager_name = that.data.NumberList[e.currentTarget.dataset.index].manager_name
    wx.navigateTo({
      url: '/pages/homes/homes?addmodel=false',
    })
  },

  delete: function (e) {
  },

  btn1: function () {
    app.globalData.issubmit=0
    wx.navigateTo({
      url: '/pages/homes/homes?addmodel=true',
    }) 
  },

//刷新
onRefresh(){
  //在当前页面显示导航条加载动画
  wx.showNavigationBarLoading(); 
  //显示 loading 提示框。需主动调用 wx.hideLoading 才能关闭提示框
  wx.showLoading({
    title: '刷新中...',
  })
  this.getData();
},
//网络请求，获取数据
getData(){
  let that =this
  
  
  // that.setData({
  //   NumberList: data
  // })
  // return
  console.log("getData")
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
          console.log(res)
          setTimeout(function () {
            //隐藏loading 提示框
            wx.hideLoading();
            //隐藏导航条加载动画
            wx.hideNavigationBarLoading();
          }, 1000)

          //停止下拉刷新
          wx.stopPullDownRefresh();
          that.setData({
            NumberList: res.data.data
          })
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
/**
* 页面相关事件处理函数--监听用户下拉动作
*/
onPullDownRefresh: function () {
  //调用刷新时将执行的方法
  this.onRefresh();
}
})