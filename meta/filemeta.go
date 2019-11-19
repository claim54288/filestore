package meta

import (
	mydb "filestore-server/db"
	"sort"
)

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

//UpdateFileMetaDB:新增/更新文件元信息到mysql中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

//GetFileMeta:通过Sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//GetFileMetaDB:从mysql获取文件元信息
func GetFileMetaDB(filesha1 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMeta(filesha1)
	if err != nil {
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
}

//GetFileMeta:获取批量的文件元信息列表
func GetLastFileMetas(limitCnt int) []FileMeta {
	fMetaArray := make([]FileMeta, 0, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:limitCnt]
}

//RemoveFileMeta:删除元信息，这边简单删除，多线程需保证安全
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
