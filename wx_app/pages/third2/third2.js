// pages/third1/third1.js
var app = getApp(); //取得全局App({..})实例
// 1. 获取数据库引用

var util = require('../../utils/date.js');
var myexif = require('../../utils/myexif.js');
/////////////////////////
const date = new Date();
const years = [];
const months = [];
const days = [];
const hours = [];
const minutes = [];
const seconds = [];
const M = date.getMonth();
const Date1 = date.getDate() - 1;
const H = date.getHours();
const minute = date.getMinutes();
const second = date.getSeconds();

//获取年
for (let i = 2023; i <= date.getFullYear() + 5; i++) {
  years.push("" + i);
}
//获取月份
for (let i = 1; i <= 12; i++) {
  if (i < 10) {
    i = "0" + i;
  }
  months.push("" + i);
}
//获取日期
for (let i = 1; i <= 31; i++) {
  if (i < 10) {
    i = "0" + i;
  }
  days.push("" + i);
}
//获取小时
for (let i = 0; i < 24; i++) {
  if (i < 10) {
    i = "0" + i;
  }
  hours.push("" + i);
}
//获取分钟
for (let i = 0; i < 60; i++) {
  if (i < 10) {
    i = "0" + i;
  }
  minutes.push("" + i);
}
//获取秒
for (let i = 0; i < 60; i++) {
  if (i < 10) {
    i = "0" + i;
  }
  seconds.push("" + i);
}
Page({
  /**
   * 页面的初始数据
   */
  data: {
    ///////////////////////////////deng
    previewImagePrPath: './images/preview.jpg',
    previewImageStPath: './images/preview.jpg',
    cameraImagePath: './images/camera.jpg',
    options: 0,
    eleId: app.globalData.elevatorNumber,
    partindex: app.globalData.partindex,
    x: '',
    y: '',
    z: '',
    // ipPort: 'http://127.0.0.1:5000',
    // ipPort: 'http://172.20.10.5:5000',
    // ipPortimg:'http://172.20.10.5:5000/images',
    ipPort: 'https://jimucloud.cn/pushimg',
    ipPortimg: 'https://jimucloud.cn/getimg',
    ///////////////////////////////////////////////////
    startTime: '',
    endTime: '',
    multiArray: [years, months, days, hours, minutes, seconds],
    multiArray1: [years, months, days, hours, minutes, seconds],
    multiIndex: [0, M, Date1, H, minute, second],
    multiIndex1: [0, M, Date1, H, minute, second],
    choose_year: '',
    fileList1: [{
      url: app.globalData.bigImg1,
      name: '图片1',
    }, ], //图片组件相关数据

    fileList2: [{
      url: app.globalData.bigImg2,
      name: '图片2',
    }, ], //图片组件相关数据

    fileList3: [{
      url: app.globalData.bigImg3,
      name: '图片3',
    }, ], //图片组件相关数据

    fileList4: [{
      url: app.globalData.bigImg4,
      name: '图片4',
    }, ], //图片组件相关数据

    show: false, //弹出框
    date: '', //显示时间
    minHour: 10,
    maxHour: 20,
    minDate: new Date().getTime(),
    maxDate: new Date(2030, 12, 31).getTime(),
    currentDate: new Date().getTime(),
    datakey: '', //数据值
    partmodel: '', //部件型号
    memo: '', //备注
    Pphoto: 0, //问题图片是否拍照
    Sphoto: 0, //标准图片是否拍照
    radio1: '1', //单选框 有无该功能
    radio2: '1', //单选框 工作是否正常
    activeNames: ['1', '2'],
    isTrouble: '',
  },

  // 上传图片

  //0代表问题图片
  afterRead(event) {
    const {
      file
    } = event.detail;
    // //0代表问题图片，1代表标准图片，2代表水印拍照，3代表拍照结果
    app.globalData.partinfor = 0
    // wx.navigateTo({

    wx.getFileSystemManager().readFile({
      filePath: file.url, //选择图片返回的相对路径
      encoding: 'base64', //编码格式
      success: res => { //成功的回调
        let userImageBase64 = 'data:image/jpg;base64,' + res.data;
        var aa = {
          'pic': userImageBase64,
          ele: app.globalData.elevatorNumber,
          'part': app.globalData.partindex,
          'button': app.globalData.partinfor
        }

        wx.request({
          url: that.data.ipPort + '/get_image',
          data: {
            aa
          },
          header: {
            'content-type': 'application/json'
          },
          method: 'POST',
          dataType: 'json',
          responseType: 'text',
          success: (result) => {
            app.globalData.bigImg1 = that.data.ipPortimg + app.globalData.elevatorNumber + '_' + app.globalData.partindex + '_' + app.globalData.partinfor + '.png'
          },
          fail: (res) => {},
          complete: () => {}
        });
      }
    })
    this.data.fileList1.url = file.url
    app.globalData.bigImg1 = file.url
  },

  ////////////////////////////////////////Deng

  //预览问题图片
  previewImage4: function (e) {
    wx.showLoading({
      title: '图片加载中...',
    })
    var that = this
    console.log(app.globalData.Problemimgurls);
    wx.downloadFile({
      url: app.globalData.Problemimgurls[0],
      success(res1) {
        wx.downloadFile({
          url: app.globalData.Problemimgurls[1],
          success(res2) {
            wx.downloadFile({
              url: app.globalData.Problemimgurls[2],
              success(res3) {
                wx.previewImage({
                  current: '',
                  urls: [res1.tempFilePath, res2.tempFilePath, res3.tempFilePath],
                  success: function (res) {
                    wx.hideLoading()
                  }
                })
              },
              fail: (res3) => {
                wx.previewImage({
                  current: '',
                  urls: [res1.tempFilePath, res2.tempFilePath],
                  success: function (res) {
                    wx.hideLoading()
                  }
                })
              }

            })
          },
          fail: (res2) => {
            wx.previewImage({
              current: '',
              urls: [res1.tempFilePath],
              success: function (res) {
                wx.hideLoading()
              }
            })
          }
        })
      },

      fail: (res1) => {
        wx.previewImage({
          current: '',
          urls: [],
          success: function (res) {
            wx.hideLoading()
          },
          fail: function (res) {
            wx.hideLoading()
          }
        })
      }


    })

  },
  //预览标准图片
  previewImage5: function (e) {
    wx.showLoading({
      title: '图片加载中...',
    })
    var that = this
    wx.downloadFile({
      url: app.globalData.Standardimgurls[0],
      success(res1) {
        wx.downloadFile({
          url: app.globalData.Standardimgurls[1],
          success(res2) {
            wx.downloadFile({
              url: app.globalData.Standardimgurls[2],
              success(res3) {
                wx.previewImage({
                  current: '',
                  urls: [res1.tempFilePath, res2.tempFilePath, res3.tempFilePath],
                  success: function (res) {
                    wx.hideLoading()
                  }
                })
              },
              fail: (res3) => {
                wx.previewImage({
                  current: '',
                  urls: [res1.tempFilePath, res2.tempFilePath],
                  success: function (res) {
                    wx.hideLoading()
                  }
                })
              }

            })
          },
          fail: (res2) => {
            wx.previewImage({
              current: '',
              urls: [res1.tempFilePath],
              success: function (res) {
                wx.hideLoading()
              }
            })
          }
        })
      },

      fail: (res1) => {
        wx.previewImage({
          current: '',
          urls: [],
          success: function (res) {
            wx.hideLoading()
          },
          fail: function (res) {
            wx.hideLoading()
          }
        })
      }

    })
  },
  /////////////////////////////////////////////////////////////
  dataInput(e) {
    this.setData({
      datakey: e.detail.value
    })
  },
  // 设置部件型号
  partInput(e) {
    this.setData({
      partmodel: e.detail.value
    })

    const self = this
    const value = e.detail.value
    var eleId = self.data.eleId
    var partindex = self.data.partindex
    if (value) {
      wx.setStorageSync(eleId + partindex + 'memoPART', value)
    } // 监听用户输入的信息，一旦有内容输入进去，就会使用wx.setStorageSync('userText', value)设置usertext这个key的值，使用wx.getStorageSync('userText')可以得到usertext这个key的值
  },

  // 备注
  memoInput(e) {
    this.setData({
      memo: e.detail.value
    })
  },

  // 折叠面板组件
  onChange(event) {
    this.setData({
      activeNames: event.detail,
    });
  },

  // 单选框 设置有无该功能
  onChange1(event) {
    this.setData({
      radio1: event.detail,
    });
    if (this.data.radio1 == 2) {
      this.setData({
        radio2: '2'
      })
    }
  },

  // 单选框 设置工作是否正常
  onChange2(event) {
    this.setData({
      radio2: event.detail,
    });
  },

  onLoad: function () {
    console.log(app.globalData);
    app.globalData.Problemimgurls = []
    app.globalData.Standardimgurls = []
    app.globalData.Pnumberlist = []
    var that = this
    //设置默认的年份
    that.setData({
      choose_year: this.data.multiArray[0][0],
      choose_year: this.data.multiArray1[0][0],
      eleId: app.globalData.elevatorNumber,
      partindex: app.globalData.partindex,
      partmodel: app.globalData.partModel
    })
  },

  ////////////////////////////////////Deng
  goToProblem: function () {
    var that = this
    wx.showActionSheet({
      itemList: ['从相册中选择', '拍照'],
      itemColor: "#000000",
      success: function (res) {
        if (!res.cancel) {
          if (res.tapIndex == 0) {
            that.chooseWxImageProblem('album')
          } else if (res.tapIndex == 1) {
            that.chooseWxImageProblem('camera')
          }
        }
      }
    })
    that.setData({
      isTrouble: 1,
      Pphoto: that.data.Pphoto + 1
    })
  },
  chooseWxImageProblem: function (type) {
    var that = this;
    console.log("yesyesyesyes!!!!!!");

    if (type == 'album') {
      wx.chooseMedia({ //选择图片

        count: 3,
        mediaType: ['image', 'video'],
        // sourceType: ['album', 'camera'],
        maxDuration: 30,
        camera: 'back',
        sourceType: [type],
        sizeType: ['original'], //图片不能经过压缩处理

        success(res) {
          console.log("类型");
          console.log(type);
          console.log(res);

          console.log("inininininablum");
          // 处理每一张图
          for (var i = 0; i < res.tempFiles.length; i++) {
            var image = res.tempFiles[i].tempFilePath

            console.log(i);
            //获取当前时间
            let newDate = new Date();
            let year = newDate.getFullYear().toString() //年

            let month0 = (newDate.getMonth() + 1) //月
            var month = month0 < 10 ? '0' + month0.toString() : month0.toString()

            let day0 = newDate.getDate() //日
            var day = day0 < 10 ? '0' + day0.toString() : day0.toString()

            var hour0 = newDate.getHours()
            var hour = hour0 < 10 ? '0' + hour0.toString() : hour0.toString()
            var Minute0 = newDate.getMinutes()
            var minute = Minute0 < 10 ? '0' + Minute0.toString() : Minute0.toString()

            var Second0 = newDate.getSeconds()
            var second = Second0 < 10 ? '0' + Second0.toString() : Second0.toString()
            var inumber = i.toString()
            that.data.photonumStr = year + month + day + hour + minute + second + inumber
            console.log('+++++++++++++');
            console.log(that.data.photonumStr);
            console.log('+++++++++++++');
            // 上传
            if (app.globalData.Problemimgurls.length == 3) {
              app.globalData.Problemimgurls.shift()
              app.globalData.Pnumberlist.shift()
            }
            app.globalData.Problemimgurls.push(that.data.ipPortimg + '/problem/' + 'problem_' + app.globalData.elevatorNumber + '_' + that.data.partindex + '_' + that.data.photonumStr + '.jpg')
            app.globalData.Pnumberlist.push(that.data.photonumStr)

            wx.showLoading({
              title: '图片上传中...',
            })
            wx.uploadFile({
              url: that.data.ipPort + '/get_image_problem', //仅为示例，非真实的接口地址
              filePath: image,
              name: 'file',
              formData: {
                eleId: app.globalData.elevatorNumber,
                partindex: that.data.partindex,
                photoId: that.data.photonumStr
              },
              success(resload) {
                wx.hideLoading()
                // 展示图片
                let pages = getCurrentPages();
                let prevPage = pages[pages.length - 1];
                prevPage.setData({
                  previewImagePrPath: image,
                })
              },
              fail: () => {
                wx.showToast({
                  title: '上传图片失败！',
                })
              } //fail
            }) //uploadfile
          } //for
        } //success
      }) //choosemedia
    } //if(album)
    if (type == 'camera') {
      var that = this
      wx.navigateTo({
        url: '../problem/index?eleId=' + app.globalData.elevatorNumber + '&partindex=' + app.globalData.partindex + '&ipPort=' + that.data.ipPort,
      })

    } //if(camera)
  },
  goToStandard: function () {
    var that = this
    wx.showActionSheet({
      itemList: ['从相册中选择', '拍照'],
      itemColor: "#000000",
      success: function (res) {
        if (!res.cancel) {
          if (res.tapIndex == 0) {
            that.chooseWxImageStandard('album')
          } else if (res.tapIndex == 1) {
            that.chooseWxImageStandard('camera')
          }
        }
      }
    })
    that.setData({
      isTrouble: 0,
      Sphoto: that.data.Sphoto + 1
    })
  },
  chooseWxImageStandard: function (type) {
    var that = this;
    console.log("yesyesyesyes!!!!!!");
    if (type == 'album') {
      wx.chooseMedia({ //选择图片

        count: 3,
        mediaType: ['image', 'video'],
        // sourceType: ['album', 'camera'],
        maxDuration: 30,
        camera: 'back',
        sourceType: [type],
        sizeType: ['original'], //图片不能经过压缩处理

        success(res) {
          console.log("类型");
          console.log(type);
          console.log(res);

          console.log("inininininablum");
          // 处理每一张图
          for (var i = 0; i < res.tempFiles.length; i++) {
            var image = res.tempFiles[i].tempFilePath

            console.log(i);
            //获取当前时间
            let newDate = new Date();
            let year = newDate.getFullYear().toString() //年

            let month0 = (newDate.getMonth() + 1) //月
            var month = month0 < 10 ? '0' + month0.toString() : month0.toString()

            let day0 = newDate.getDate() //日
            var day = day0 < 10 ? '0' + day0.toString() : day0.toString()

            var hour0 = newDate.getHours()
            var hour = hour0 < 10 ? '0' + hour0.toString() : hour0.toString()
            var Minute0 = newDate.getMinutes()
            var minute = Minute0 < 10 ? '0' + Minute0.toString() : Minute0.toString()

            var Second0 = newDate.getSeconds()
            var second = Second0 < 10 ? '0' + Second0.toString() : Second0.toString()
            var inumber = i.toString()
            that.data.photonumStr = year + month + day + hour + minute + second + inumber
            console.log(that.data.photonumStr);
            // 上传
            if (app.globalData.Standardimgurls.length == 3) {
              app.globalData.Standardimgurls.shift()
              app.globalData.Pnumberlist.shift()
            }
            app.globalData.Standardimgurls.push(that.data.ipPortimg + '/standard/' + 'standard_' + app.globalData.elevatorNumber + '_' + that.data.partindex + '_' + that.data.photonumStr + '.jpg')
            app.globalData.Pnumberlist.push(that.data.photonumStr)
            console.log(app.globalData.Pnumberlist);
            console.log(that.data.photonumStr)
            wx.showLoading({
              title: '图片上传中...',
            })
            wx.uploadFile({
              url: that.data.ipPort + '/get_image_standard', //仅为示例，非真实的接口地址
              filePath: image,
              name: 'file',
              formData: {
                eleId: app.globalData.elevatorNumber,
                partindex: that.data.partindex,
                photoId: that.data.photonumStr
              },

              success(resload) {
                wx.hideLoading()
                console.log("upload");
                // 展示图片
                let pages = getCurrentPages();
                let prevPage = pages[pages.length - 1];
                prevPage.setData({
                  previewImageStPath: image,
                })


              },
              fail: () => {
                wx.showToast({
                  title: '上传图片失败！',
                })
              }

            }) //uploadfile

          } //for 

        } //success

      }) //choosemedia
    }
    if (type == 'camera') {
      var that = this
      wx.navigateTo({
        url: '../standard/index?eleId=' + app.globalData.elevatorNumber + '&partindex=' + that.data.partindex + '&ipPort=' + that.data.ipPort,
      })
    }
  },

  /////////////////////////////////////////////////
  //获取时间日期
  bindMultiPickerChange: function (e) {
    this.setData({
      multiIndex: e.detail.value,
      eleId: app.globalData.elevatorNumber,
      partindex: app.globalData.partindex
    })
    const index = this.data.multiIndex;
    const year = this.data.multiArray[0][index[0]];
    const month = this.data.multiArray[1][index[1]];
    const day = this.data.multiArray[2][index[2]];
    const hour = this.data.multiArray[3][index[3]];
    const minute = this.data.multiArray[4][index[4]];
    const seconds = this.data.multiArray[5][index[5]];
    this.setData({
      startTime: year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + seconds
    })
  },
  //获取时间日期
  bindMultiPickerChange1: function (e) {
    this.setData({
      multiIndex1: e.detail.value
    })
    const index = this.data.multiIndex1;
    const year = this.data.multiArray1[0][index[0]];
    const month = this.data.multiArray1[1][index[1]];
    const day = this.data.multiArray1[2][index[2]];
    const hour = this.data.multiArray1[3][index[3]];
    const minute = this.data.multiArray1[4][index[4]];
    const seconds = this.data.multiArray[5][index[5]];
    this.setData({
      endTime: year + '-' + month + '-' + day + ' ' + hour + ':' + minute + ':' + seconds
    })
  },
  //监听picker的滚动事件
  bindMultiPickerColumnChange: function (e) {
    //获取年份
    if (e.detail.column == 0) {
      let choose_year = this.data.multiArray[e.detail.column][e.detail.value];
      this.setData({
        choose_year
      })
    }
    if (e.detail.column == 1) {
      let num = parseInt(this.data.multiArray[e.detail.column][e.detail.value]);
      let temp = [];
      if (num == 1 || num == 3 || num == 5 || num == 7 || num == 8 || num == 10 || num == 12) { //判断31天的月份
        for (let i = 1; i <= 31; i++) {
          if (i < 10) {
            i = "0" + i;
          }
          temp.push("" + i);
        }
        this.setData({
          ['multiArray[2]']: temp
        });
      } else if (num == 4 || num == 6 || num == 9 || num == 11) { //判断30天的月份
        for (let i = 1; i <= 30; i++) {
          if (i < 10) {
            i = "0" + i;
          }
          temp.push("" + i);
        }
        this.setData({
          ['multiArray[2]']: temp
        });
      } else if (num == 2) { //判断2月份天数
        let year = parseInt(this.data.choose_year);
        if (((year % 400 == 0) || (year % 100 != 0)) && (year % 4 == 0)) {
          for (let i = 1; i <= 29; i++) {
            if (i < 10) {
              i = "0" + i;
            }
            temp.push("" + i);
          }
          this.setData({
            ['multiArray[2]']: temp
          });
        } else {
          for (let i = 1; i <= 28; i++) {
            if (i < 10) {
              i = "0" + i;
            }
            temp.push("" + i);
          }
          this.setData({
            ['multiArray[2]']: temp
          });
        }
      }
    }
    var data = {
      multiArray: this.data.multiArray,
      multiIndex: this.data.multiIndex
    };
    data.multiIndex[e.detail.column] = e.detail.value;
    this.setData(data);
  },
  bindMultiPickerColumnChange1: function (e) {
    //获取年份
    if (e.detail.column == 0) {
      let choose_year = this.data.multiArray1[e.detail.column][e.detail.value];
      this.setData({
        choose_year
      })
    }
    if (e.detail.column == 1) {
      let num = parseInt(this.data.multiArray1[e.detail.column][e.detail.value]);
      let temp = [];
      if (num == 1 || num == 3 || num == 5 || num == 7 || num == 8 || num == 10 || num == 12) { //判断31天的月份
        for (let i = 1; i <= 31; i++) {
          if (i < 10) {
            i = "0" + i;
          }
          temp.push("" + i);
        }
        this.setData({
          ['multiArray1[2]']: temp
        });
      } else if (num == 4 || num == 6 || num == 9 || num == 11) { //判断30天的月份
        for (let i = 1; i <= 30; i++) {
          if (i < 10) {
            i = "0" + i;
          }
          temp.push("" + i);
        }
        this.setData({
          ['multiArray1[2]']: temp
        });
      } else if (num == 2) { //判断2月份天数
        let year = parseInt(this.data.choose_year);
        if (((year % 400 == 0) || (year % 100 != 0)) && (year % 4 == 0)) {
          for (let i = 1; i <= 29; i++) {
            if (i < 10) {
              i = "0" + i;
            }
            temp.push("" + i);
          }
          this.setData({
            ['multiArray1[2]']: temp
          });
        } else {
          for (let i = 1; i <= 28; i++) {
            if (i < 10) {
              i = "0" + i;
            }
            temp.push("" + i);
          }
          this.setData({
            ['multiArray1[2]']: temp
          });
        }
      }
    }
    var data = {
      multiArray1: this.data.multiArray1,
      multiIndex1: this.data.multiIndex1
    };
    data.multiIndex1[e.detail.column] = e.detail.value;
    this.setData(data);
  },
  bindDateChange: function (e) {

    this.setData({
      date: e.detail.value
    })
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {
    const self = this
    var eleId = self.data.eleId

    var partindex = self.data.partindex

    console.log(eleId);
    console.log(partindex);
    let userText = wx.getStorageSync(eleId + partindex + 'memoPART')
    if (userText) {
      self.data.partmodel = userText
      self.setData(self.data)
    } // page载入的时候先读取一次，wx.getStorageSync('userText')里面有没有内容，有内容就填充，没有则什么也不做

    this.setData({
      fileList1: [{
        url: app.globalData.bigImg1,
        name: '图片1',
      }, ],
      fileList2: [{
        url: app.globalData.bigImg2,
        name: '图片2',
      }, ],
      fileList3: [{
        url: app.globalData.bigImg3,
        name: '图片3',
      }, ],
      fileList4: [{
        url: app.globalData.bigImg4,
        name: '图片4',
      }, ]

    })
  },

  submit: function (e) {
    var eleId = this.data.eleId
    var partindex = this.data.partindex
    var radio1 = this.data.radio1
    var radio2 = this.data.radio2
    var partmodel = this.data.partmodel
    // var memo = this.data.memo
    var Pphoto = app.globalData.Problemimgurls.length
    var Sphoto = app.globalData.Standardimgurls.length
    var flag = false //是否已完善内容
    var flag2 = true //是否已上传过一次标准图片

    if (radio1 == '2') {
      if (partmodel != '' && Pphoto != 0) {
        flag = true
      }
    } else if (radio1 == '1' && radio2 == '2') {
      if (partmodel != '' && Pphoto != 0) {
        flag = true
      }
    } else if (radio1 == '1' && radio2 == '1') {
      if (partmodel != '') {
        flag = true
      }
      if (Sphoto == 0) {
        wx.request({
          url: 'https://jimucloud.cn/basedata/pic/selectpic',
          data: {
            partId: partindex,
            eleId: eleId
          },
          header: {
            'content-type': 'application/json' // 默认值
          },
          method: 'POST',
          success: (res) => {
            if (res.data.data) { //如果返回数据不为空 即之前上传过标准图片
              console.log(res)
              console.log(res.data.data['picUrl']);
              var getPagenum = res.data.data['picUrl']
              getPagenum = getPagenum.toString()
              console.log(getPagenum);
              var refresh = getPagenum.slice(-19, -4)
              console.log(refresh);
              app.globalData.Pnumberlist.push(refresh)
              console.log(app.globalData.Pnumberlist);
            } else { //没有上传过一次标准图片
              flag2 = false
              console.log(flag2);
            }

          }
        })


      }
    }
    setTimeout(() => {
      if (flag == false && flag2 == false) {
        wx.showToast({
          title: '请完善内容',
          'icon': 'none',
          duration: 3000
        })
      } else if (flag == true && flag2 == false) {
        wx.showModal({
          title: '提示',
          content: '请上传一次标准图片',
          success(res) {
            if (res.confirm) {
              console.log('用户点击确定')
            } else if (res.cancel) {
              console.log('用户点击取消')
            }
          }
        })
      } else {
        let that = this
        console.log(app.globalData.Pnumberlist);

        wx.request({
          url: 'https://jimucloud.cn/basedata/part/insert',
          header: {
            "Content-Type": "application/json"
          },
          method: 'post',
          data: {
            maintenanceProject: app.globalData.partindex, //零件id
            partData: that.data.datakey,
            partModel: that.data.partmodel,
            isNormal: that.data.radio2,
            function: that.data.radio1,
            nextInspectionTime: that.data.startTime,
            remark: that.data.memo,
            random: app.globalData.Pnumberlist,
            maintenanceOrder: app.globalData.workNumber,
            isAi: 0,
            isTrouble: that.data.isTrouble,
          },

          success: res => {
            console.log('chenggong');
            console.log(res);
            wx.showToast({
              title: '提交成功',
            })

            // wx.redirectTo({
            //   url: '/pages/Alevel2/Alevel2',
            // })
            console.log(app.globalData);
            wx.navigateBack({
              success: function () {
                let pages = getCurrentPages();
                let beforePage = pages[pages.length - 2];
                // beforePage.changeData();
              }
            })
          },
        })


      }
      console.log(flag)
    }, 2000)
  },

})