// index.js
// 获取应用实例
const app = getApp()
const defaultAvatarUrl = 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0'
Page({
  data: {
    wxOpenId:'',
    wxCode:'',
    motto: '孩童孤独症康复计划小程序',
    avatarUrl: defaultAvatarUrl,
    theme: wx.getSystemInfoSync().theme,
  },
  onLoad() {
    let that = this
    wx.onThemeChange((result) => {
      this.setData({
        theme: result.theme
      })
    })
    //用户登录
    wx.login({
      success: (res) => {
          console.log(res);
          that.setData({
              wxCode: res.code,
          })
          // return
          // ====== 【获取OpenId】
          let m_code = that.data.wxCode; // 获取code
          let m_AppId = "wxace1b2884c347d29"; // appid wxace1b2884c347d29
          wx.request({
              url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/user/getOpenId',//http://43.143.180.52
              // url: 'https://www.lprs.site/api/v1/user/getOpenId',//http://43.143.180.52
              header: {
                "Content-Type": "application/json"
              },
              data:{
                "id":m_AppId,
                "js_code":m_code,
              },
              method: 'get',
              success: (res) => {
                if (res.data.msg=="sucess") {
                  console.log(res);
                  this.setData({
                      wxOpenId: res.data.data.openid
                  })
                  console.log("openid:"+this.data.wxOpenId)
                } else {
                  wx.showModal({
                    title: '提示',
                    content: '获取微信id失败',
                    showCancel:false,
                  })
                }
              },
              fail(){
                wx.showModal({
                  title: '提示',
                  content: '获取微信id失败',
                  showCancel:false,
                })
              },
          })
      }
    })
  },
  onChooseAvatar(e) {
    const { avatarUrl } = e.detail 
    this.setData({
      avatarUrl,
    })
  },
  btn6: function () {
    // wx.navigateTo({
    //   url: '/pages/child/child',
    // })
    // return
    let that = this
    // console.log(that.data.wxOpenId)
    //通过小程序id，code获取当前登陆者的微信openid
   
    wx.request({
      
      url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/user/login',
      // url: 'https://www.lprs.site/api/v1/user/login',
      header: {
        "Content-Type": "application/json"
      },
      method: 'get',
      data: {
        "wechat_id": that.data.wxOpenId,
      },
      success: res => {
        if (res.data.msg=="sucess") {//openid登录成功直接跳转
          wx.navigateTo({
            url: '/pages/child/child',
          })
        } else {
          wx.showModal({
            title: '未注册！',
            content:'请将微信ID：'+that.data.wxOpenId+'\n提供给管理员',
            showCancel:false,
            confirmText: '确定并复制文本',
            success(res) {
              if (res.confirm) {
                wx.setClipboardData({
                  data: that.data.wxOpenId,
                  success(res) {
                    wx.getClipboardData({
                      success(res) {
                        console.log(res.data) // data
                      }
                    })
                  }
                })
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
        })
      },
    })
  },
})
