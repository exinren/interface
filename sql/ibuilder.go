package sql

type IBuilder interface {
	GenSelect(*DbData) string
	GenPage(*DbData) string
}