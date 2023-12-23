package seniverse

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type V4WeatherDailyResponse struct {
	Status     string    `json:"status"`      // 状态
	ApiVersion string    `json:"api_version"` // api版本
	ApiStatus  string    `json:"api_status"`  // api状态
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    float64   `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime float64   `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Alert struct {
			Status  string `json:"status"`
			Content []struct {
				Pubtimestamp  int       `json:"pubtimestamp"` // 发布时间，单位是 Unix 时间戳
				AlertID       string    `json:"alertId"`      // 预警 ID
				Status        string    `json:"status"`       // 预警信息的状态
				Adcode        string    `json:"adcode"`       // 区域代码
				Location      string    `json:"location"`     // 位置
				Province      string    `json:"province"`     // 省
				City          string    `json:"city"`         // 市
				County        string    `json:"county"`       // 县
				Code          string    `json:"code"`         // 预警代码
				Source        string    `json:"source"`       // 发布单位
				Title         string    `json:"title"`        // 标题
				Description   string    `json:"description"`  // 描述
				RegionID      string    `json:"regionId"`
				Latlon        []float64 `json:"latlon"`
				RequestStatus string    `json:"request_status"`
			} `json:"content"`
			Adcodes []struct {
				Adcode int    `json:"adcode"`
				Name   string `json:"name"`
			} `json:"adcodes"` // 行政区划层级信息
		} `json:"alert"` // 预警数据
		WeatherDaily struct {
			Status string `json:"status"`
			Astro  []struct {
				Date    string `json:"date"`
				Sunrise struct {
					Time string `json:"time"`
				} `json:"sunrise"`
				Sunset struct {
					Time string `json:"time"`
				} `json:"sunset"`
			} `json:"astro"` // 日出日落时间，当地时区的时刻，tzshift 不作用在这个变量)
			Precipitation08H20H []struct {
				Date        string  `json:"date"`
				Max         float64 `json:"max"`
				Min         float64 `json:"min"`
				Avg         float64 `json:"avg"`
				Probability float64 `json:"probability"`
			} `json:"precipitation_08h_20h"` // 白天降水数据
			Precipitation20H32H []struct {
				Date        string  `json:"date"`
				Max         float64 `json:"max"`
				Min         float64 `json:"min"`
				Avg         float64 `json:"avg"`
				Probability float64 `json:"probability"`
			} `json:"precipitation_20h_32h"` // 夜晚降水数据
			Precipitation []struct {
				Date        string  `json:"date"`
				Max         float64 `json:"max"`
				Min         float64 `json:"min"`
				Avg         float64 `json:"avg"`
				Probability float64 `json:"probability"`
			} `json:"precipitation"` // 降水数据
			Temperature []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature"` // 全天地表 2 米气温
			Temperature08H20H []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature_08h_20h"` // 白天地表 2 米气温
			Temperature20H32H []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature_20h_32h"` // 夜晚地表 2 米气温
			Wind []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"` // 全天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"` // 全天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"` // 全天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind"`
			Wind08H20H []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"` // 白天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"` // 白天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"` // 白天地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind_08h_20h"`
			Wind20H32H []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"` // 夜晚地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"` // 夜晚地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"` // 夜晚地表 10 米风速
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind_20h_32h"`
			Humidity []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"humidity"` // 地表 2 米相对湿度(%)
			Cloudrate []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"cloudrate"` // 云量(0.0-1.0)
			Pressure []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"pressure"` // 地面气压
			Visibility []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"visibility"` // 地表水平能见度
			Dswrf []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"dswrf"` // 向下短波辐射通量(W/M2)
			AirQuality struct {
				Aqi []struct {
					Date string `json:"date"`
					Max  struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"max"`
					Avg struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"avg"`
					Min struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"min"`
				} `json:"aqi"` // 国标 AQI
				Pm25 []struct {
					Date string  `json:"date"`
					Max  float64 `json:"max"`
					Avg  float64 `json:"avg"`
					Min  float64 `json:"min"`
				} `json:"pm25"` // PM2.5 浓度(μg/m3)
			} `json:"air_quality"`
			Skycon []struct {
				Date  string `json:"date"`
				Value string `json:"value"` // 全天主要 天气现象
			} `json:"skycon"`
			Skycon08H20H []struct {
				Date  string `json:"date"`
				Value string `json:"value"` // 白天主要 天气现象
			} `json:"skycon_08h_20h"`
			Skycon20H32H []struct {
				Date  string `json:"date"`
				Value string `json:"value"` // 夜晚主要 天气现象
			} `json:"skycon_20h_32h"`
			LifeIndex struct {
				Ultraviolet []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 紫外线指数自然语言
				} `json:"ultraviolet"`
				CarWashing []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 洗车指数自然语言
				} `json:"carWashing"`
				Dressing []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 穿衣指数自然语言
				} `json:"dressing"`
				Comfort []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 舒适度指数自然语言
				} `json:"comfort"`
				ColdRisk []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"` // 感冒指数自然语言
				} `json:"coldRisk"`
			} `json:"life_index"`
		} `json:"WeatherDaily"` // 天级别预报
		Primary float64 `json:"primary"`
	} `json:"result"`
}

type V4WeatherDailyResult struct {
	Result V4WeatherDailyResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求·
}

func newV4WeatherDailyResult(result V4WeatherDailyResponse, body []byte, http gorequest.Response) *V4WeatherDailyResult {
	return &V4WeatherDailyResult{Result: result, Body: body, Http: http}
}

// WeatherDaily 天气网格预报（中国/15天/逐日）
// https://seniverse.yuque.com/hyper_data/api_v4/weather_daily
func (c *V4Client) WeatherDaily(ctx context.Context, locations string, notMustParams ...gorequest.Params) (*V4WeatherDailyResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("locations", locations)
	// 请求
	request, err := c.request(ctx, apiUrlV4+"?fields=weather_daily", params, http.MethodGet)
	if err != nil {
		return newV4WeatherDailyResult(V4WeatherDailyResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response V4WeatherDailyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newV4WeatherDailyResult(response, request.ResponseBody, request), err
}
