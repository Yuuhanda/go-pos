package model

// PointType defines the type of point transaction
type PointType string

const (
	PointTypeEarned   PointType = "EARNED"
	PointTypeRedeemed PointType = "REDEEMED"
)

// MemberPoint represents the member_point table in the database
type MemberPoint struct {
	ID       int       `json:"id_point" db:"id_point(32)"`
	MemberID int       `json:"id_member" db:"id_member"`
	Type     PointType `json:"type" db:"type"`
	Points   int       `json:"points" db:"point(32)s"`
	
	// Optional relation field (not in database)
	Member   *Member   `json:"member,omitempty" db:"-"`
}
