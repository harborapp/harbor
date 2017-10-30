package storage

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	migrations = []*gormigrate.Migration{
		{
			ID: "20171030135018",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        int64  `gorm:"primary_key"`
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
			ID: "20171030135032",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        int64  `gorm:"primary_key"`
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
			ID: "20171030135038",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID int64 `sql:"index"`
					UserID int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamUser{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_users").Error
			},
		},
		{
			ID: "20171030135044",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135053",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135101",
			Migrate: func(tx *gorm.DB) error {
				type Registry struct {
					ID        int64  `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					Host      string `sql:"unique_index"`
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
			ID: "20171030135118",
			Migrate: func(tx *gorm.DB) error {
				type Org struct {
					ID         int64 `gorm:"primary_key"`
					RegistryID int64 `sql:"index"`
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
			ID: "20171030135125",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135130",
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
			ID: "20171030135137",
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
			ID: "20171030135143",
			Migrate: func(tx *gorm.DB) error {
				type TeamOrg struct {
					TeamID int64 `sql:"index"`
					OrgID  int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&TeamOrg{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_orgs").Error
			},
		},
		{
			ID: "20171030135149",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135155",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135201",
			Migrate: func(tx *gorm.DB) error {
				type UserOrg struct {
					UserID int64 `sql:"index"`
					OrgID  int64 `sql:"index"`
					Perm   string
				}

				return tx.CreateTable(&UserOrg{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_orgs").Error
			},
		},
		{
			ID: "20171030135207",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135213",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135219",
			Migrate: func(tx *gorm.DB) error {
				type Repo struct {
					ID        int64 `gorm:"primary_key"`
					OrgID     int64 `sql:"index"`
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
			ID: "20171030135226",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135233",
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
			ID: "20171030135239",
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
			ID: "20171030135245",
			Migrate: func(tx *gorm.DB) error {
				type Tag struct {
					ID        int64 `gorm:"primary_key"`
					RepoID    int64 `sql:"index"`
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
			ID: "20171030135251",
			Migrate: func(tx *gorm.DB) error {
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

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
				if tx.Dialect().GetName() == "sqlite3" {
					return nil
				}

				return gormigrate.ErrRollbackImpossible
			},
		},
		{
			ID: "20171030135256",
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
			ID: "20171030135302",
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
