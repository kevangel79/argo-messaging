package stores

import "time"

// Store encapsulates the generic store interface
type Store interface {
	Initialize()
	QuerySubsByTopic(projectUUID, topic string) ([]QSub, error)
	QueryTopicsByACL(projectUUID, user string) ([]QTopic, error)
	QuerySubsByACL(projectUUID, user string) ([]QSub, error)
	QuerySubs(projectUUID string, userUUID string, name string, pageToken string, pageSize int32) ([]QSub, int32, string, error)
	QueryTopics(projectUUID string, userUUID string, name string, pageToken string, pageSize int32) ([]QTopic, int32, string, error)
	QueryDailyTopicMsgCount(projectUUID string, name string, date time.Time) ([]QDailyTopicMsgCount, error)
	UpdateTopicLatestPublish(projectUUID string, name string, date time.Time) error
	UpdateTopicPublishRate(projectUUID string, name string, rate float64) error
	UpdateSubLatestConsume(projectUUID string, name string, date time.Time) error
	UpdateSubConsumeRate(projectUUID string, name string, rate float64) error
	RemoveTopic(projectUUID string, name string) error
	RemoveSub(projectUUID string, name string) error
	PaginatedQueryUsers(pageToken string, pageSize int32, projectUUID string) ([]QUser, int32, string, error)
	QueryUsers(projectUUID string, uuid string, name string) ([]QUser, error)
	UpdateUser(uuid string, projects []QProjectRoles, name string, email string, serviceRoles []string, modifiedOn time.Time) error
	AppendToUserProjects(userUUID string, projectUUID string, pRoles ...string) error
	UpdateUserToken(uuid string, token string) error
	RemoveUser(uuid string) error
	QueryProjects(uuid string, name string) ([]QProject, error)
	UpdateProject(projectUUID string, name string, description string, modifiedOn time.Time) error
	RemoveProject(uuid string) error
	RemoveProjectTopics(projectUUID string) error
	RemoveProjectSubs(projectUUID string) error
	QueryDailyProjectMsgCount(projectUUID string) ([]QDailyProjectMsgCount, error)
	QueryTotalMessagesPerProject(projectUUIDs []string, startDate time.Time, endDate time.Time) ([]QProjectMessageCount, error)
	InsertUser(uuid string, projects []QProjectRoles, name string, token string, email string, serviceRoles []string, createdOn time.Time, modifiedOn time.Time, createdBy string) error
	InsertProject(uuid string, name string, createdOn time.Time, modifiedOn time.Time, createdBy string, description string) error
	InsertOpMetric(hostname string, cpu float64, mem float64) error
	InsertTopic(projectUUID string, name string) error
	IncrementTopicMsgNum(projectUUID string, name string, num int64) error
	IncrementDailyTopicMsgCount(projectUUID string, topicName string, num int64, date time.Time) error
	IncrementTopicBytes(projectUUID string, name string, totalBytes int64) error
	IncrementSubBytes(projectUUID string, name string, totalBytes int64) error
	IncrementSubMsgNum(projectUUID string, name string, num int64) error
	InsertSub(projectUUID string, name string, topic string, offest int64, maxMessages int64, ack int, push string, rPolicy string, rPeriod int, vhash string, verified bool) error
	HasProject(name string) bool
	HasUsers(projectUUID string, users []string) (bool, []string)
	QueryOneSub(projectUUID string, name string) (QSub, error)
	QueryPushSubs() []QSub
	HasResourceRoles(resource string, roles []string) bool
	GetOpMetrics() []QopMetric
	GetUserRoles(projectUUID string, token string) ([]string, string)
	GetUserFromToken(token string) (QUser, error)
	UpdateSubOffset(projectUUID string, name string, offset int64)
	UpdateSubPull(projectUUID string, name string, offset int64, ts string) error
	UpdateSubOffsetAck(projectUUID string, name string, offset int64, ts string) error
	ModSubPush(projectUUID string, name string, push string, maxMessages int64, rPolicy string, rPeriod int, vhash string, verified bool) error
	QueryACL(projectUUID string, resource string, name string) (QAcl, error)
	ExistsInACL(projectUUID string, resource string, resourceName string, userUUID string) error
	ModACL(projectUUID string, resource string, name string, acl []string) error
	AppendToACL(projectUUID string, resource string, name string, acl []string) error
	RemoveFromACL(projectUUID string, resource string, name string, acl []string) error
	ModAck(projectUUID string, name string, ack int) error
	GetAllRoles() []string
	Clone() Store
	Close()
}
