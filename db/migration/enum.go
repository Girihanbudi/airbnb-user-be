package migration

type Migration int

const (
	MigrationUp Migration = iota
	MigrationDown
)

func (m Migration) String() string {
	return []string{"up", "down"}[m]
}
