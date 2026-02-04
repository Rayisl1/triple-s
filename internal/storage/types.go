package storage

type BucketMeta struct {
	Name         string
	CreationDate string
}
type ObjectMeta struct {
	Name         string
	Size         int64
	ContentType  string
	CreationDate string
}
