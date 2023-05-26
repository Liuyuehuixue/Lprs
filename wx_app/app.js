App({
  /**
   * 当小程序初始化完成时，会触发 onLaunch（全局只触发一次）
   */
  onLaunch: function () {
//     wx.cloud.init({
//       env: 'file-download-5gbptut32c562d45',
//       traceUser: true,
//     })
  },

  /**
   * 当小程序启动，或从后台进入前台显示，会触发 onShow
   */
  onShow: function (options) {
  },

  /**
   * 当小程序从前台进入后台，会触发 onHide
   */
  onHide: function () {

  },

  /**
   * 当小程序发生脚本错误，或者 api 调用失败时，会触发 onError 并带上错误信息
   */
  onError: function (msg) {

  },

  //全局变量
  globalData: {
    code: '',//孩子编码
    name: '',//孩子名字
    age: '',//孩子月龄
    manager_name: '', //负责人
    test_n: '',//第几次测试
    commited:false,

    detail_child_id: '',//孩子编码
    detail_name: '',//孩子名字
    detail_times: '',//孩子总次数
    detail_manager_name: '', //负责人
    detail_table_id: '',
    detail_time_id: '',
  }

})