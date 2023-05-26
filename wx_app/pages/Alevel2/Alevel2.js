
var app = getApp(); //取得全局App({..})实例
var output = [];
const num2options = new Map([
  [1, '儿童感知觉发展能力'],
  [2, '儿童粗大动作能力'],
  [3, '儿童精细动作能力'],
  [4, '儿童语言与沟通发展能力'],
  [5, '儿童认知发展能力'],
  [6, '儿童社会交往发展能力'],
  [7, '儿童生活自理发展能力'],
]);
const addlist = {//测试用数组，可删除
  "code":"111",
  "AllItems":[
    {
      "survey_name":"感知觉",
      "Items":[
        {
          "survey_sub_range":"视觉追视",
          "survey_project_cnname":"追视飘动物体",
          "survey_project_enname": "4_pdwt",
          "score":"P"
        },
        {
          "survey_sub_range":"听觉辨别",
          "survey_project_cnname":"别人叫自己名字时会作出反应，而对其它名字则没有反应",
          "survey_project_enname": "26_imfy",
          "score":"P"
        }
      ],
  },
  {
    "survey_name":"粗大动作1",
    "Items":[
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"坐姿，双手离地，转动躯干",
        "survey_project_enname": "4_pdwt",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      }
    ]
  },
  {
    "survey_name":"粗大动作2",
    "Items":[
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"坐姿，双手离地，转动躯干",
        "survey_project_enname": "4_pdwt",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      }
    ]
  },
  {
    "survey_name":"粗大动作3",
    "Items":[
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"坐姿，双手离地，转动躯干",
        "survey_project_enname": "4_pdwt",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      }
    ]
  },
  {
    "survey_name":"粗大动作4",
    "Items":[
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"坐姿，双手离地，转动躯干",
        "survey_project_enname": "4_pdwt",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      }
    ]
  },
  {
    "survey_name":"粗大动作5",
    "Items":[
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"坐姿，双手离地，转动躯干",
        "survey_project_enname": "4_pdwt",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      },
    ]
  },
  {
    "survey_name":"粗大动作6",
    "Items":[
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"坐姿，双手离地，转动躯干",
        "survey_project_enname": "4_pdwt",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      },
      {
        "survey_sub_range":"坐姿",
        "survey_project_cnname":"扶桌子由站转至坐地",
        "survey_project_enname": "26_imfy",
        "score":"P"
      },
    ]
  },
]
};

const checklist = {//测试用数组，可删除
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
    allList: [],
    showList: [],
    firstList: [],
    secondList: [],
    thirdList: [],
    fourthList: [],
    fivthList: [],
    sixthList: [],
    seventhList: [],
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
      }
    ], //下拉菜单
    optionvalue: 1, //初始选择
    alltimes:1,
    times: [], //次数下拉菜单
    timesvalue: 1, //初始选择
    flag: false, //选择
    addmodel: true,
    submited:true,
    disabled:true,
  },
  onLoad: function (options) {
    let that = this
    console.log(options)
    that.setData({
      addmodel: (options.addmodel === "true"),
      disabled: (options.addmodel === "true"),
      alltimes:options.alltimes,
      timesvalue:options.alltimes,
    })
    that.data.times.pop()
    console.log(that.data.times)
    var newtimes = []
    for (var i = 1; i <= that.data.alltimes; i++) {
      newtimes.push(
        {
          text: String(i),
          value: i,
        }
      )
    }
    that.setData({
      times:newtimes,
    })
    // console.log(that.data.alltimes)
    // console.log(that.data.timesvalue)
    console.log(that.data.times)
    that.getbaseData()
    // that.onReady()
  },

  getbaseData() {
    let that = this
    if (that.data.addmodel) {//添加数据请求
      wx.request({
        url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/children/selectItemsByAge',
        // url: 'https://www.lprs.site/api/v1/children/selectItemsByAge',
        header: {
          "Content-Type": "application/json"
        },
        method: 'post',
        data: {
          "code": Number(app.globalData.code),
          "name": app.globalData.name,
          "age": Number(app.globalData.age),
          "manager_name": app.globalData.manager_name, 
          "test_n": app.globalData.test_n,
        },
        success: res => {
          if (res.data.msg=="sucess") {
            that.data.allList = res.data.data
            // console.log(that.data.allList)
            that.setData({
              firstList: that.data.allList.AllItems[0].Items,
              secondList: that.data.allList.AllItems[1].Items,
              thirdList: that.data.allList.AllItems[2].Items,
              fourthList: that.data.allList.AllItems[3].Items,
              fivthList: that.data.allList.AllItems[4].Items,
              sixthList: that.data.allList.AllItems[5].Items,
              seventhList: that.data.allList.AllItems[6].Items,
            })
            that.setData({
              showList:that.data.firstList,
            })
          } else {
            wx.showModal({
              title: '输入设置有误或已提交',
              content:'请输入正确设置或者查看本次记录！',
              showCancel:false,
              success (res) {
                if (res.confirm) {   //确认则跳转
                  wx.navigateBack({})
                }
              }
            })
          }
        },
        fail(){
          wx.showModal({
            title: '提示',
            content: '数据请求失败',
            showCancel:false,
            success (res) {
              if (res.confirm) {   //确认则跳转
                wx.navigateBack({})
              }
            }
          })
        },
      })
      if (that.data.allList == []) {
        wx.showModal({   //提示没有表项
          title: '提示',
          content: '当前次数无评估表，可能是当前年龄范围内无可用评估项',
          showCancel:false,
        })
      }
    }
    // console.log(this.data.timesvalue)
    this.onReady()
  },
  getCheckData(){//查看数据请求
    let that = this
    wx.request({
      url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/survey/showSurveyById',
      // url: 'https://www.lprs.site/api/v1/survey/showSurveyById',
      header: {
        "Content-Type": "application/json"
      },
      method: 'get',
      data: {
        "child_id": app.globalData.detail_child_id,//孩子id
        "test_n": that.data.timesvalue,//次数
      },
      success: res => {
        if (res.data.msg=="sucess") {
          that.data.allList = res.data.data
          // console.log(that.data.allList)
          that.setData({
            submited:(res.data.data.status === 1),
            firstList: that.data.allList.AllItems[0].Items,
            secondList: that.data.allList.AllItems[1].Items,
            thirdList: that.data.allList.AllItems[2].Items,
            fourthList: that.data.allList.AllItems[3].Items,
            fivthList: that.data.allList.AllItems[4].Items,
            sixthList: that.data.allList.AllItems[5].Items,
            seventhList: that.data.allList.AllItems[6].Items,
          })
          // console.log(that.data.submited)
          that.setData({
            disabled:(!that.data.submited) | (that.data.addmodel),
            showList:that.data.firstList,
          })
          // console.log(that.data.submited)
          // console.log(that.data.addmodel)
          // console.log(that.data.disabled)
        } else {
          wx.showModal({
            title: '提示',
            content: '当前选择暂无数据，请先保存或提交评估表',
            showCancel:false,
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
  //表格选择
  tableChange(e) {
    let that = this
    var temp = !that.data.flag
    that.setData({
      flag: temp,
      showList:that.data.allList.AllItems[e.detail-1].Items,
    })
    // console.log(that.data.showList)
  },
  
  //次数选择
  timeChange(e) {
    let that = this
    // that.setData({
    //   timesvalue
    // })
    if (!that.data.addmodel) {
      that.getCheckData()
    }
  },

  //返回刷新
  changeData: function (options) {
    // this.onLoad(options);
  },
  onReady(){
    // this.setData({
    //   timesvalue:this.data.alltimes,
    // })
  },
  onShow: function (options) {
    // this.onLoad(options);
  },
  radioChange: function (e) {
    let that = this
    var index = e.target.dataset.index//当前showList的数组索引
    var value = e.detail//新选择的选项
    //将选项对应的数组内score修改为选项的值
    that.setData({
      [`showList[${index}].score`]: value,
    })
    console.log(that.data.showList[index].score)
  },
  //保存调用接口
  save: function (e) {
    let that = this
    console.log("保存调用接口")
    
    output = that.data.allList
    console.log(output)
    wx.request({
      url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/survey/insertItems/0',
      // url: 'https://www.lprs.site/api/v1/survey/insertItems/0',
      data: {
        "code":output.code,
        "age":output.age,
        "test_n":output.test_n,
        "AllItems":output.AllItems,
      },
      header: {
        'content-type': "application/json" // 默认值
      },
      method: 'post',
      success: (res) => {
        if (res.data.msg == "sucess") {
          wx.showModal({
            title: '数据已保存',
            showCancel:false,
          })
        } else {
          wx.showModal({
            title: '保存失败',
            showCancel:false,
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

  //提交调用接口
  submit: function (e) {
    let that = this
    
    output = that.data.allList
    console.log("提交调用接口")
    for (var i = 0; i < output.AllItems.length; i++) {
      for (var j = 0; j < output.AllItems[i].Items.length; j++) {
        if(output.AllItems[i].Items[j].score == '') {
          console.log("有选项未填无法提交")
          wx.showModal({   //添加确认提示
            title: '提示',
            content: '有选项未填无法提交',
            showCancel:false,
          })
          that.save()
          return
        }
      }
    }
    
    console.log(output)
    wx.showModal({
      title: '请确认是否提交',
      success: function (sm) {
        if (sm.confirm) {
          let that = this
          wx.request({
            url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/survey/insertItems/1',
            // url: 'https://www.lprs.site/api/v1/survey/insertItems/1',
            data: {
              "code":output.code,
              "age":output.age,
              "test_n":output.test_n,
              "AllItems":output.AllItems,
            },
            dataType:"json",
            header: {
              'content-type': "application/json" // 默认值
            },
            method: 'post',
            success: (res) => {
              console.log(res)
              if (res.data.msg == "sucess") {
                wx.showModal({
                  title: '数据已录入',
                  showCancel:false,
                })
                wx.redirectTo({
                  url: '/pages/child/child',
                })
              } else {
                wx.showModal({
                  title: '录入失败',
                  showCancel:false,
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
        } else if (sm.cancel) {
          console.log('用户点击取消')
        }
      }
    })
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