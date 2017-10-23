package storage

import (
	// "time"

	// "github.com/go-xorm/xorm"
	"github.com/go-xorm/xorm/migrate"
)

var (
	migrations = []*migrate.Migration{
	// {
	// 	ID: "201608311535",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type User struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Username  string `sql:"unique_index"`
	// 			Email     string `sql:"unique_index"`
	// 			Hash      string `sql:"unique_index"`
	// 			Hashword  string
	// 			Active    bool `sql:"default:false"`
	// 			Admin     bool `sql:"default:false"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&User{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("user")
	// 	},
	// },
	// {
	// 	ID: "201608311536",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Team struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Team{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team")
	// 	},
	// },
	// {
	// 	ID: "201608311537",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type TeamUser struct {
	// 			TeamID int64 `sql:"index"`
	// 			UserID int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&TeamUser{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team_user")
	// 	},
	// },
	// {
	// 	ID: "201608311538",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_user",
	// 		).AddForeignKey(
	// 			"team_id",
	// 			"team(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311538",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_user",
	// 		).AddForeignKey(
	// 			"user_id",
	// 			"user(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311540",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Registry struct {
	// 			ID        int64  `gorm:"primary_key"`
	// 			Slug      string `sql:"unique_index"`
	// 			Name      string `sql:"unique_index"`
	// 			Host      string `sql:"unique_index"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Registry{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("registry")
	// 	},
	// },
	// {
	// 	ID: "201608311541",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Org struct {
	// 			ID         int64 `gorm:"primary_key"`
	// 			RegistryID int64 `sql:"index"`
	// 			Slug       string
	// 			Name       string
	// 			Public     bool `sql:"default:false"`
	// 			CreatedAt  time.Time
	// 			UpdatedAt  time.Time
	// 		}

	// 		return engine.CreateTables(&Org{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("org")
	// 	},
	// },
	// {
	// 	ID: "201608311542",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"org",
	// 		).AddForeignKey(
	// 			"registry_id",
	// 			"registry(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311543",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"org",
	// 		).AddUniqueIndex(
	// 			"uix_org_registry_id_slug",
	// 			"registry_id",
	// 			"slug",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"org",
	// 		).RemoveIndex(
	// 			"uix_org_registry_id_slug",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201608311544",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"org",
	// 		).AddUniqueIndex(
	// 			"uix_org_registry_id_name",
	// 			"registry_id",
	// 			"name",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"org",
	// 		).RemoveIndex(
	// 			"uix_org_registry_id_name",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201608311545",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type TeamOrg struct {
	// 			TeamID int64 `sql:"index"`
	// 			OrgID  int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&TeamOrg{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("team_org")
	// 	},
	// },
	// {
	// 	ID: "201608311546",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_org",
	// 		).AddForeignKey(
	// 			"team_id",
	// 			"team(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311547",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"team_org",
	// 		).AddForeignKey(
	// 			"org_id",
	// 			"org(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311548",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type UserOrg struct {
	// 			UserID int64 `sql:"index"`
	// 			OrgID  int64 `sql:"index"`
	// 			Perm   string
	// 		}

	// 		return engine.CreateTables(&UserOrg{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("user_org")
	// 	},
	// },
	// {
	// 	ID: "201608311549",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"user_org",
	// 		).AddForeignKey(
	// 			"user_id",
	// 			"user(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311550",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"user_org",
	// 		).AddForeignKey(
	// 			"org_id",
	// 			"org(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311551",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Repo struct {
	// 			ID        int64 `gorm:"primary_key"`
	// 			OrgID     int64 `sql:"index"`
	// 			Slug      string
	// 			Name      string
	// 			FullName  string
	// 			Public    bool `sql:"default:false"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Repo{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("repo")
	// 	},
	// },
	// {
	// 	ID: "201608311552",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"repo",
	// 		).AddForeignKey(
	// 			"org_id",
	// 			"org(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311553",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"repo",
	// 		).AddUniqueIndex(
	// 			"uix_repo_org_id_slug",
	// 			"org_id",
	// 			"slug",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"repo",
	// 		).RemoveIndex(
	// 			"uix_repo_org_id_slug",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201608311554",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"repo",
	// 		).AddUniqueIndex(
	// 			"uix_repo_org_id_name",
	// 			"org_id",
	// 			"name",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"repo",
	// 		).RemoveIndex(
	// 			"uix_repo_org_id_name",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201608311555",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		type Tag struct {
	// 			ID        int64 `gorm:"primary_key"`
	// 			RepoID    int64 `sql:"index"`
	// 			Slug      string
	// 			Name      string
	// 			FullName  string
	// 			Public    bool `sql:"default:false"`
	// 			CreatedAt time.Time
	// 			UpdatedAt time.Time
	// 		}

	// 		return engine.CreateTables(&Tag{})
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return engine.DropTables("tag")
	// 	},
	// },
	// {
	// 	ID: "201608311556",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return tx.Table(
	// 			"tag",
	// 		).AddForeignKey(
	// 			"repo_id",
	// 			"repo(id)",
	// 			"RESTRICT",
	// 			"RESTRICT",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		if engine.DriverName() == "sqlite3" {
	// 			return nil
	// 		}

	// 		return gormigrate.ErrRollbackImpossible
	// 	},
	// },
	// {
	// 	ID: "201608311557",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"tag",
	// 		).AddUniqueIndex(
	// 			"uix_tag_repo_id_slug",
	// 			"repo_id",
	// 			"slug",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"tag",
	// 		).RemoveIndex(
	// 			"uix_tag_repo_id_slug",
	// 		).Error
	// 	},
	// },
	// {
	// 	ID: "201608311558",
	// 	Migrate: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"tag",
	// 		).AddUniqueIndex(
	// 			"uix_tag_repo_id_name",
	// 			"repo_id",
	// 			"name",
	// 		).Error
	// 	},
	// 	Rollback: func(engine *xorm.Engine) error {
	// 		return tx.Table(
	// 			"tag",
	// 		).RemoveIndex(
	// 			"uix_tag_repo_id_name",
	// 		).Error
	// 	},
	// },
	}
)
