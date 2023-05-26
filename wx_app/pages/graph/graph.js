var app = getApp();
var lineChart = null;
Page({
  data: {
    list: [{
        age: '72',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '71',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '70',
        gzj: '',
        cddz: '72',
        jxdz: '66',
        yyygt: '',
        rzfz: '55',
        shjw: '',
        shzl: '67',
        fzfs: '441'
      },
      {
        age: '69',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '68',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '79',
        rzfz: '',
        shjw: '47',
        shzl: '',
        fzfs: '421'
      },
      {
        age: '67',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '66',
        gzj: '55',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '50',
        shjw: '',
        shzl: '',
        fzfs: '416'
      },
      {
        age: '65',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '64',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '63',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '62',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '61',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '60',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '59',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '58',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '57',
        gzj: '',
        cddz: '65',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '56',
        gzj: '',
        cddz: '',
        jxdz: '63',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '55',
        gzj: '52',
        cddz: '64',
        jxdz: '62',
        yyygt: '76',
        rzfz: '42',
        shjw: '45',
        shzl: '62',
        fzfs: '405'
      },
      {
        age: '54',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '53',
        gzj: '',
        cddz: '',
        jxdz: '51',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '330'
      },
      {
        age: '52',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '51',
        gzj: '',
        cddz: '',
        jxdz: '50',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '329'
      },
      {
        age: '50',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '49',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '48',
        gzj: '47',
        cddz: '47',
        jxdz: '49',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '328'
      },
      {
        age: '47',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '46',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '45',
        gzj: '',
        cddz: '',
        jxdz: '48',
        yyygt: '',
        rzfz: '',
        shjw: '40',
        shzl: '',
        fzfs: '323'
      },
      {
        age: '44',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '43',
        gzj: '44',
        cddz: '46',
        jxdz: '47',
        yyygt: '',
        rzfz: '30',
        shjw: '',
        shzl: '46',
        fzfs: '312'
      },
      {
        age: '42',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '67',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '267'
      },
      {
        age: '41',
        gzj: '',
        cddz: '',
        jxdz: '39',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '255'
      },
      {
        age: '40',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '39',
        gzj: '',
        cddz: '35',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '249'
      },
      {
        age: '38',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '37',
        gzj: '40',
        cddz: '',
        jxdz: '35',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '248'
      },
      {
        age: '36',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '35',
        gzj: '',
        cddz: '',
        jxdz: '34',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '244'
      },
      {
        age: '34',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '33',
        gzj: '37',
        cddz: '',
        jxdz: '',
        yyygt: '55',
        rzfz: '',
        shjw: '30',
        shzl: '34',
        fzfs: '243'
      },
      {
        age: '32',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '31',
        gzj: '36',
        cddz: '34',
        jxdz: '',
        yyygt: '',
        rzfz: '20',
        shjw: '',
        shzl: '33',
        fzfs: '234'
      },
      {
        age: '30',
        gzj: '',
        cddz: '',
        jxdz: '32',
        yyygt: '52',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '192'
      },
      {
        age: '29',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '28',
        gzj: '',
        cddz: '24',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '167'
      },
      {
        age: '27',
        gzj: '',
        cddz: '',
        jxdz: '24',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '18',
        fzfs: '163'
      },
      {
        age: '26',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '25',
        gzj: '29',
        cddz: '',
        jxdz: '23',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '160'
      },
      {
        age: '24',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '24',
        shzl: '',
        fzfs: '157'
      },
      {
        age: '23',
        gzj: '',
        cddz: '22',
        jxdz: '22',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '152'
      },
      {
        age: '22',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '21',
        gzj: '27',
        cddz: '21',
        jxdz: '21',
        yyygt: '36',
        rzfz: '10',
        shjw: '19',
        shzl: '15',
        fzfs: '149'
      },
      {
        age: '20',
        gzj: '',
        cddz: '',
        jxdz: '20',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '19',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '18',
        gzj: '',
        cddz: '19',
        jxdz: '',
        yyygt: '',
        rzfz: '9',
        shjw: '15',
        shzl: '12',
        fzfs: '123'
      },
      {
        age: '17',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: ''
      },
      {
        age: '16',
        gzj: '21',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '8',
        fzfs: '93'
      },
      {
        age: '15',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '27',
        rzfz: '5',
        shjw: '14',
        shzl: '',
        fzfs: '89'
      },
      {
        age: '14',
        gzj: '',
        cddz: '7',
        jxdz: '11',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '6',
        fzfs: '79'
      },
      {
        age: '13',
        gzj: '',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '3',
        shjw: '',
        shzl: '',
        fzfs: 75
      },
      {
        age: '12',
        gzj: '19',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '5',
        fzfs: '73'
      },
      {
        age: '11',
        gzj: '',
        cddz: '6',
        jxdz: '9',
        yyygt: '21',
        rzfz: '2',
        shjw: '',
        shzl: '',
        fzfs: '68'
      },
      {
        age: '10',
        gzj: '16',
        cddz: '',
        jxdz: '4',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '58'
      },
      {
        age: '9',
        gzj: '',
        cddz: '5',
        jxdz: '',
        yyygt: '',
        rzfz: '1',
        shjw: '1',
        shzl: '3',
        fzfs: '31'
      },
      {
        age: '8',
        gzj: '',
        cddz: '',
        jxdz: '3',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '28'
      },
      {
        age: '7',
        gzj: '',
        cddz: '1',
        jxdz: '',
        yyygt: '18',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '27'
      },
      {
        age: '6',
        gzj: '',
        cddz: '0',
        jxdz: '2',
        yyygt: '8',
        rzfz: '0',
        shjw: '4',
        shzl: '2',
        fzfs: '26'
      },
      {
        age: '5',
        gzj: '10',
        cddz: '',
        jxdz: '1',
        yyygt: '2',
        rzfz: '',
        shjw: '1',
        shzl: '1',
        fzfs: ''
      },
      {
        age: '4',
        gzj: '5',
        cddz: '',
        jxdz: '1',
        yyygt: '1',
        rzfz: '',
        shjw: '1',
        shzl: '',
        fzfs: '9'
      },
      {
        age: '3',
        gzj: '2',
        cddz: '',
        jxdz: '0',
        yyygt: '0',
        rzfz: '',
        shjw: '0',
        shzl: '0',
        fzfs: '2'
      },
      {
        age: '2',
        gzj: '1',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '1'
      },
      {
        age: '1',
        gzj: '0',
        cddz: '',
        jxdz: '',
        yyygt: '',
        rzfz: '',
        shjw: '',
        shzl: '',
        fzfs: '0'
      },
    ]
  },
  onLoad: function (e) {
    // var arr = that.data.list
    // for (var i= 0;i < length(arr);i++) {
    //   console.log('INSERT INTO \`metric_score\` VALUES (\'' +arr[i].age+ '\', \'' +arr[i].gzj+ '\', \'' +arr[i].cddz+ '\', \'' +arr[i].jxdz+ '\', \'' +arr[i].yyygt+ '\', \'' +arr[i].rzfz+ '\', \'' +arr[i].shjw+ '\', \'' +arr[i].shzl+ '\', \'' +arr[i].fzfs+ '\');')
    // }

    let age_evaluate, score, survey_id
    let titleList = ['gzj', 'cddz', 'jxdz', 'yyygt', 'rzfz', 'shjw', 'shzl', 'fzfs']
    let title
    let result = this.data.list
    wx.request({
      url: 'https://ontoweb.wust.edu.cn/lqrs/api/v1/children/plot',
      // url: 'https://www.lprs.site/api/v1/children/plot',
      method: "GET",
      data: {
        child_id: e.child_id,
      },
      success: (res) => {
        console.log(res.data.data);
        let arr = res.data.data
        let socreArr
        for (let i = 0; i < arr.length; i++) {
          socreArr = arr[i].AllScore
          for (let j = 0; j < socreArr.length; j++) {
            age_evaluate = arr[i].AllScore[j]['age_evaluate']
            score = arr[i].AllScore[j]['p_n']
            survey_id = arr[i].AllScore[j]['survey_id']
            title = titleList[survey_id-1]
            result[72 - age_evaluate][title] = result[72 - age_evaluate][title] + "\n" + `<p>${score}<sub>${i+1}</sub></p>`
            //result[72 - age_evaluate][title]
          }
        }
        this.setData({
          list: result
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
});