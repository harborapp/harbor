package data

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	migrations = []*gormigrate.Migration{
		{
			ID: "201608311535",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        int    `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Username  string `sql:"unique_index"`
					Email     string `sql:"unique_index"`
					Hash      string `sql:"unique_index"`
					Hashword  string
					Active    bool `sql:"default:false"`
					Admin     bool `sql:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "201608311536",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        int    `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Team{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("teams").Error
			},
		},
		{
			ID: "201608311537",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID int `sql:"index"`
					UserID int `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamUser{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_users").Error
			},
		},
		{
			ID: "201608311538",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"team_users",
				).AddForeignKey(
					"team_id",
					"teams(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311538",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"team_users",
				).AddForeignKey(
					"user_id",
					"users(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311540",
			Migrate: func(tx *gorm.DB) error {
				type Registry struct {
					ID        int    `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					Host      string `sql:"unique_index"`
					UseSSL    bool   `sql:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Registry{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("registries").Error
			},
		},
		{
			ID: "201608311541",
			Migrate: func(tx *gorm.DB) error {
				type Org struct {
					ID         int `gorm:"primary_key"`
					RegistryID int `sql:"index"`
					Slug       string
					Name       string
					Public     bool `sql:"default:false"`
					CreatedAt  time.Time
					UpdatedAt  time.Time
				}

				return tx.CreateTable(&Org{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("orgs").Error
			},
		},
		{
			ID: "201608311542",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"orgs",
				).AddForeignKey(
					"registry_id",
					"registries(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311543",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"orgs",
				).AddUniqueIndex(
					"uix_orgs_registry_id_slug",
					"registry_id",
					"slug",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"orgs",
				).RemoveIndex(
					"uix_orgs_registry_id_slug",
				).Error
			},
		},
		{
			ID: "201608311544",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"orgs",
				).AddUniqueIndex(
					"uix_orgs_registry_id_name",
					"registry_id",
					"name",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"orgs",
				).RemoveIndex(
					"uix_orgs_registry_id_name",
				).Error
			},
		},
		{
			ID: "201608311545",
			Migrate: func(tx *gorm.DB) error {
				type TeamOrg struct {
					TeamID int `sql:"index"`
					OrgID  int `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamOrg{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_orgs").Error
			},
		},
		{
			ID: "201608311546",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"team_orgs",
				).AddForeignKey(
					"team_id",
					"teams(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311547",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"team_orgs",
				).AddForeignKey(
					"org_id",
					"orgs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311548",
			Migrate: func(tx *gorm.DB) error {
				type UserOrg struct {
					UserID int `sql:"index"`
					OrgID  int `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&UserOrg{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_orgs").Error
			},
		},
		{
			ID: "201608311549",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"user_orgs",
				).AddForeignKey(
					"user_id",
					"users(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311550",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"user_orgs",
				).AddForeignKey(
					"org_id",
					"orgs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311551",
			Migrate: func(tx *gorm.DB) error {
				type Repo struct {
					ID        int `gorm:"primary_key"`
					OrgID     int `sql:"index"`
					Slug      string
					Name      string
					FullName  string
					Public    bool `sql:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Repo{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("repos").Error
			},
		},
		{
			ID: "201608311552",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"repos",
				).AddForeignKey(
					"org_id",
					"orgs(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311553",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"repos",
				).AddUniqueIndex(
					"uix_repos_org_id_slug",
					"org_id",
					"slug",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"repos",
				).RemoveIndex(
					"uix_repos_org_id_slug",
				).Error
			},
		},
		{
			ID: "201608311554",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"repos",
				).AddUniqueIndex(
					"uix_repos_org_id_name",
					"org_id",
					"name",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"repos",
				).RemoveIndex(
					"uix_repos_org_id_name",
				).Error
			},
		},
		{
			ID: "201608311555",
			Migrate: func(tx *gorm.DB) error {
				type Tag struct {
					ID        int `gorm:"primary_key"`
					RepoID    int `sql:"index"`
					Slug      string
					Name      string
					FullName  string
					Public    bool `sql:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Tag{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("tags").Error
			},
		},
		{
			ID: "201608311556",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"tags",
				).AddForeignKey(
					"repo_id",
					"repos(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "201608311557",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"tags",
				).AddUniqueIndex(
					"uix_tags_repo_id_slug",
					"repo_id",
					"slug",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"tags",
				).RemoveIndex(
					"uix_tags_repo_id_slug",
				).Error
			},
		},
		{
			ID: "201608311558",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"tags",
				).AddUniqueIndex(
					"uix_tags_repo_id_name",
					"repo_id",
					"name",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"tags",
				).RemoveIndex(
					"uix_tags_repo_id_name",
				).Error
			},
		},
	}
)
