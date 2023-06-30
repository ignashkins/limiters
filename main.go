package main

import (
	"time"

	"github.com/Clever/leakybucket"
	"github.com/Clever/leakybucket/memory"
	"github.com/sirupsen/logrus"
)

var s = memory.New()

func main() {

	var cap uint = 3

	bucket, err := s.Create("my-bucket", cap, 5*time.Second)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(time.Until(bucket.Reset()))

	for i := 1; i <= 50; i++ {

		if bucket.Remaining() == 0 {
			sl := time.Until(bucket.Reset())
			logrus.Error("LIMIT EXCEEDED!!! Will wait ~ ", sl)
			time.Sleep(sl)
		}

		state, err := bucket.Add(1)
		if err != nil {
			logrus.Error(err)
		}

		doingSome(state, i)
	}
}

func doingSome(state leakybucket.BucketState, i int) {
	logrus.Warn(state)
	time.Sleep(500 * time.Millisecond)
}
