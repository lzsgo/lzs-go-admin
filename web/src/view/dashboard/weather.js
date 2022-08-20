
import axios from 'axios'
import { ref } from 'vue'

const weatherInfo = ref('今天天气不多 具体信息没查询到 检查您的amapKey是否有效')
const amapKey = '024ada5cde0a8194884ff6436ec08eae'

export const useWeatherInfo = () => {
  ip()
  return weatherInfo
}

export const ip = async() => {
  // key换成你自己的 https://console.amap.com/dev/index
  if (amapKey === '') {
    return false
  }
  const res = await axios.get('https://restapi.amap.com/v3/ip?key=' + amapKey)
  if (res.data.adcode) {
    getWeather(res.data.adcode)
  }
}

const getWeather = async(code) => {
  const response = await axios.get('https://restapi.amap.com/v3/weather/weatherInfo?key=' + amapKey + '&extensions=base&city=' + code)
  if (response.data.status === '1') {
    const s = response.data.lives[0]
    weatherInfo.value = s.city + ' 天气：' + s.weather + ' 温度：' + s.temperature + '摄氏度 风向：' + s.winddirection + ' 风力：' + s.windpower + '级 空气湿度：' + s.humidity
  }
}

