package meta

//FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1 string //文件Sha1值
	FileName string //文件名
	FileSize int64  //文件大小
	Location string //本地路径
	UploadAt string //上传时间
}

var fileMetas map[string]FileMeta

//首次运行初始化
func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta:通过Sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
