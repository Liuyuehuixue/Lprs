// 1. 获取数据库引用
var app = getApp(); //取得全局App({..})实例
var output = [];
const list = {
  "Surveyname":"感知觉",
  "AllTimesItems":[
    {
      "p_n":2,
      "e_n":2,
      "f_n":2,
      "test_n":1,
      "ItemsWithTest":[
        {
          "survey_project_cnname":"追视飘动物体",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        },
        {
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "score":"P"
        }
      ],
    }
  ]
};

Page({
  data: {
    isScroll: true, //主页面是否滚动
    showModal: false, //弹窗
    showList: [],
    flag: false, //选择
  },
  onLoad: function (options) {
    let that = this
    that.getbaseData() 
    wx.setNavigationBarTitle({
      title: options.title,
    })
  },
  getbaseData() {
    let that = this
    that.setData({
      showList:list.AllTimesItems[0].ItemsWithTest,
    })
    return
    wx.request({//获取孩子某张表的单次评估结果
      url: '/api/v1/survey/showSurveyById',
      header: {
        "Content-Type": "application/json"
      },
      method: 'get',
      data: {
        child_id: app.globalData.detail_child_id,//孩子id
        survey_id: app.globalData.detail_optionvalue,//表id
        test_n: app.globalData.detail_timesvalue,//次数
      },
      success: res => {
        if (res.msg == "success") {
          that.setData({
            showList: res.data.AllTimesItems[0].ItemsWithTest,
          })
        } else {
          wx.showModal({
            title: '提示',
            content: '请求失败或没有此记录',
            showCancel: false,
            complete () {
              wx.navigateTo({
                url: '/pages/child_detail/child_detail',
              })
            }
          })
        }
      },
      fail(){
        wx.showModal({
          title: '提示',
          content: '请求失败或没有此记录',
          showCancel: false,
          complete () {
            wx.navigateTo({
              url: '/pages/child_detail/child_detail',
            })
          },
        })
      }
    })
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