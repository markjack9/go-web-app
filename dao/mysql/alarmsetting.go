package mysql

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/pkg/snowflake"
	"go-web-app/pkg/todaytime"
)

func AlarmAdd(host *models.ParamAlarmSetting) (theId int64, err error) {
	sqlStr := "insert into alarmstatistics(alarmid,hostid,alarmstatus,alarmtype,alarminfo,alarmnote,alarmstarttime) values (?,?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr,
		snowflake.IdNum(),
		host.Hostid,
		1,
		host.AlarmType,
		host.AlarmInfo,
		host.AlarmNote,
		todaytime.NowTimeFull())
	if err != nil {
		return
	}
	theId, err = ret.LastInsertId()
	if err != nil {
		return theId, err
	} else {
		fmt.Printf("插入数据的id 为 %d. \n", theId)
	}
	return

}
func AlarmSettingInit(host *models.ParamAlarmSetting) (alarmsetingdata []models.ParamAlarmSetting, err error) {

	sqlStr := `select cpuoption,memoryoption,systemdiskoption,thresholdstatus,workapiurl,workatuser,dingapiurl,dingatuser from notification;`
	if err := db.Select(&alarmsetingdata, sqlStr); err != nil {
		return alarmsetingdata, err
	}
	return
}
func AlarmUpdateNoti(host *models.ParamAlarmSetting) (n int64, err error) {
	sqlStr := `update notification set workapiurl=?,workatuser=?,dingapiurl=?,dingatuser=?`
	ret, err := db.Exec(sqlStr,
		host.WorkApiUrl,
		host.WorkAtuser,
		host.DingApiUrl,
		host.DingAtuser,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", n)
	}
	return
}
func AlarmUpdateThreshold(host *models.ParamAlarmSetting) (n int64, err error) {
	sqlStr := `update notification set cpuoption=?,memoryoption=?,systemdiskoption=?,thresholdstatus=? `
	ret, err := db.Exec(sqlStr,
		host.CpuOption,
		host.MemoryOption,
		host.SystemDiskOption,
		host.ThresholdStatus,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", n)
	}
	return
}

func AlarmEdit(host *models.ParamAlarmSetting) (n int64, err error) {
	sqlStr := `update alarmsetting set alarmtype=?,alarmstatus=?,alarmowner=? where hostid=? `
	ret, err := db.Exec(sqlStr,
		host.AlarmType,
		host.AlarmStatus,
		host.AlarmHostOnwer,
		host.Alarmid,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", n)
	}
	return
}
