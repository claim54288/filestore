package meta

import "time"

const baseFormat = "2006-01-02 15:04:05"

type ByUploadTime []FileMeta

func (a ByUploadTime) Len() int {
	return len(a)
}

func (a ByUploadTime) Less(i, j int) bool {
	t1, _ := time.Parse(baseFormat, a[i].UploadAt)
	t2, _ := time.Parse(baseFormat, a[i].UploadAt)
	return t1.UnixNano() > t2.UnixNano() //这个前面的大于后面的时候返回false表示从小到大，i>j时候返回true表示从大到小
}

func (a ByUploadTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
