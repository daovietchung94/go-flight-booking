package utils

import (
	"go-training/pb"
	"time"
)

func ToPbDateTime(modelTime time.Time) *pb.DateTime {
	utctime := modelTime.UTC()
	return &pb.DateTime{
		Year:   int32(utctime.Year()),
		Month:  int32(utctime.Month()),
		Day:    int32(utctime.Day()),
		Hour:   int32(utctime.Hour()),
		Minute: int32(utctime.Minute()),
		Second: int32(utctime.Second()),
	}
}

func ToTime(pbtime *pb.DateTime) time.Time {
	return time.Date(int(pbtime.Year), time.Month(pbtime.Month), int(pbtime.Day), int(pbtime.Hour), int(pbtime.Minute), int(pbtime.Second), 0, time.UTC)
}
