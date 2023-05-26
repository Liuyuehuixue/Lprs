const formatTime = date => {
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = date.getHours()
    const minute = date.getMinutes()

  
    return [year, month, day].map(formatNumber).join('-') 
  }
  
  const formatNumber = n => {
    n = n.toString()
    return n[1] ? n : '0' + n
  }

  //提取日期
  const formatDay = date => {
    const day = date.getDate()
    return day
  }

  // 调整时间格式
  const formatSignTime = date =>{
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()

    return year + '年' + month + '月' + day + '日'
  }

  // 调整时间格式
  const formatSigntime = date =>{
    const hour = formatNumber(date.getHours())
    const minute = formatNumber(date.getMinutes())

    return hour + ":" +minute
  }
  
  // 向外暴露
  module.exports = {
    formatTime: formatTime,
    formatSign: formatSignTime,
    formatmine: formatSigntime,
    formatDay: formatDay
  }
  
  
  