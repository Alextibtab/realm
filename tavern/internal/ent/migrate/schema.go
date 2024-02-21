// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BeaconsColumns holds the columns for the "beacons" table.
	BeaconsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "principal", Type: field.TypeString, Nullable: true},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "agent_identifier", Type: field.TypeString, Nullable: true},
		{Name: "last_seen_at", Type: field.TypeTime, Nullable: true},
		{Name: "interval", Type: field.TypeUint64, Nullable: true},
		{Name: "beacon_host", Type: field.TypeInt},
	}
	// BeaconsTable holds the schema information for the "beacons" table.
	BeaconsTable = &schema.Table{
		Name:       "beacons",
		Columns:    BeaconsColumns,
		PrimaryKey: []*schema.Column{BeaconsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "beacons_hosts_host",
				Columns:    []*schema.Column{BeaconsColumns[9]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "size", Type: field.TypeInt, Default: 0},
		{Name: "hash", Type: field.TypeString, Size: 100},
		{Name: "content", Type: field.TypeBytes, SchemaType: map[string]string{"mysql": "LONGBLOB"}},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
	}
	// HostsColumns holds the columns for the "hosts" table.
	HostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "identifier", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "primary_ip", Type: field.TypeString, Nullable: true},
		{Name: "platform", Type: field.TypeEnum, Enums: []string{"PLATFORM_BSD", "PLATFORM_LINUX", "PLATFORM_MACOS", "PLATFORM_UNSPECIFIED", "PLATFORM_WINDOWS"}},
		{Name: "last_seen_at", Type: field.TypeTime, Nullable: true},
	}
	// HostsTable holds the schema information for the "hosts" table.
	HostsTable = &schema.Table{
		Name:       "hosts",
		Columns:    HostsColumns,
		PrimaryKey: []*schema.Column{HostsColumns[0]},
	}
	// HostCredentialsColumns holds the columns for the "host_credentials" table.
	HostCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "principal", Type: field.TypeString},
		{Name: "secret", Type: field.TypeString, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"KIND_PASSWORD", "KIND_SSH_KEY", "KIND_UNSPECIFIED"}},
		{Name: "host_credential_host", Type: field.TypeInt},
		{Name: "task_reported_credentials", Type: field.TypeInt},
	}
	// HostCredentialsTable holds the schema information for the "host_credentials" table.
	HostCredentialsTable = &schema.Table{
		Name:       "host_credentials",
		Columns:    HostCredentialsColumns,
		PrimaryKey: []*schema.Column{HostCredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "host_credentials_hosts_host",
				Columns:    []*schema.Column{HostCredentialsColumns[6]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "host_credentials_tasks_reported_credentials",
				Columns:    []*schema.Column{HostCredentialsColumns[7]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// HostFilesColumns holds the columns for the "host_files" table.
	HostFilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "path", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString, Nullable: true},
		{Name: "group", Type: field.TypeString, Nullable: true},
		{Name: "permissions", Type: field.TypeString, Nullable: true},
		{Name: "size", Type: field.TypeUint64, Default: 0},
		{Name: "hash", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "content", Type: field.TypeBytes, Nullable: true},
		{Name: "host_files", Type: field.TypeInt, Nullable: true},
		{Name: "host_file_host", Type: field.TypeInt},
		{Name: "task_reported_files", Type: field.TypeInt},
	}
	// HostFilesTable holds the schema information for the "host_files" table.
	HostFilesTable = &schema.Table{
		Name:       "host_files",
		Columns:    HostFilesColumns,
		PrimaryKey: []*schema.Column{HostFilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "host_files_hosts_files",
				Columns:    []*schema.Column{HostFilesColumns[10]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "host_files_hosts_host",
				Columns:    []*schema.Column{HostFilesColumns[11]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "host_files_tasks_reported_files",
				Columns:    []*schema.Column{HostFilesColumns[12]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// HostProcessesColumns holds the columns for the "host_processes" table.
	HostProcessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "pid", Type: field.TypeUint64},
		{Name: "ppid", Type: field.TypeUint64},
		{Name: "name", Type: field.TypeString},
		{Name: "principal", Type: field.TypeString},
		{Name: "path", Type: field.TypeString, Nullable: true},
		{Name: "cmd", Type: field.TypeString, Nullable: true},
		{Name: "env", Type: field.TypeString, Nullable: true},
		{Name: "cwd", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"STATUS_DEAD", "STATUS_IDLE", "STATUS_LOCK_BLOCKED", "STATUS_PARKED", "STATUS_RUN", "STATUS_SLEEP", "STATUS_STOP", "STATUS_TRACING", "STATUS_UNINTERUPTIBLE_DISK_SLEEP", "STATUS_UNKNOWN", "STATUS_UNSPECIFIED", "STATUS_WAKE_KILL", "STATUS_WAKING", "STATUS_ZOMBIE"}},
		{Name: "host_processes", Type: field.TypeInt, Nullable: true},
		{Name: "host_process_host", Type: field.TypeInt},
		{Name: "task_reported_processes", Type: field.TypeInt},
	}
	// HostProcessesTable holds the schema information for the "host_processes" table.
	HostProcessesTable = &schema.Table{
		Name:       "host_processes",
		Columns:    HostProcessesColumns,
		PrimaryKey: []*schema.Column{HostProcessesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "host_processes_hosts_processes",
				Columns:    []*schema.Column{HostProcessesColumns[12]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "host_processes_hosts_host",
				Columns:    []*schema.Column{HostProcessesColumns[13]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "host_processes_tasks_reported_processes",
				Columns:    []*schema.Column{HostProcessesColumns[14]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// QuestsColumns holds the columns for the "quests" table.
	QuestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "parameters", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "quest_tome", Type: field.TypeInt},
		{Name: "quest_bundle", Type: field.TypeInt, Nullable: true},
		{Name: "quest_creator", Type: field.TypeInt, Nullable: true},
	}
	// QuestsTable holds the schema information for the "quests" table.
	QuestsTable = &schema.Table{
		Name:       "quests",
		Columns:    QuestsColumns,
		PrimaryKey: []*schema.Column{QuestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "quests_tomes_tome",
				Columns:    []*schema.Column{QuestsColumns[5]},
				RefColumns: []*schema.Column{TomesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "quests_files_bundle",
				Columns:    []*schema.Column{QuestsColumns[6]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "quests_users_creator",
				Columns:    []*schema.Column{QuestsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "public_key", Type: field.TypeString, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "private_key", Type: field.TypeString, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "repository_owner", Type: field.TypeInt, Nullable: true},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repositories_users_owner",
				Columns:    []*schema.Column{RepositoriesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"group", "service"}},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "claimed_at", Type: field.TypeTime, Nullable: true},
		{Name: "exec_started_at", Type: field.TypeTime, Nullable: true},
		{Name: "exec_finished_at", Type: field.TypeTime, Nullable: true},
		{Name: "output", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "output_size", Type: field.TypeInt, Default: 0},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "quest_tasks", Type: field.TypeInt},
		{Name: "task_beacon", Type: field.TypeInt},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_quests_tasks",
				Columns:    []*schema.Column{TasksColumns[9]},
				RefColumns: []*schema.Column{QuestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tasks_beacons_beacon",
				Columns:    []*schema.Column{TasksColumns[10]},
				RefColumns: []*schema.Column{BeaconsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TomesColumns holds the columns for the "tomes" table.
	TomesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "support_model", Type: field.TypeEnum, Enums: []string{"UNSPECIFIED", "FIRST_PARTY", "COMMUNITY"}, Default: "UNSPECIFIED"},
		{Name: "tactic", Type: field.TypeEnum, Enums: []string{"UNSPECIFIED", "RECON", "RESOURCE_DEVELOPMENT", "INITIAL_ACCESS", "EXECUTION", "PERSISTENCE", "PRIVILEGE_ESCALATION", "DEFENSE_EVASION", "CREDENTIAL_ACCESS", "DISCOVERY", "LATERAL_MOVEMENT", "COLLECTION", "COMMAND_AND_CONTROL", "EXFILTRATION", "IMPACT"}, Default: "UNSPECIFIED"},
		{Name: "param_defs", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "hash", Type: field.TypeString, Size: 100},
		{Name: "eldritch", Type: field.TypeString, SchemaType: map[string]string{"mysql": "LONGTEXT"}},
		{Name: "tome_uploader", Type: field.TypeInt, Nullable: true},
		{Name: "tome_repository", Type: field.TypeInt, Nullable: true},
	}
	// TomesTable holds the schema information for the "tomes" table.
	TomesTable = &schema.Table{
		Name:       "tomes",
		Columns:    TomesColumns,
		PrimaryKey: []*schema.Column{TomesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tomes_users_uploader",
				Columns:    []*schema.Column{TomesColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tomes_repositories_repository",
				Columns:    []*schema.Column{TomesColumns[12]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 25},
		{Name: "oauth_id", Type: field.TypeString, Unique: true},
		{Name: "photo_url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "MEDIUMTEXT"}},
		{Name: "session_token", Type: field.TypeString, Size: 200},
		{Name: "access_token", Type: field.TypeString, Size: 200},
		{Name: "is_activated", Type: field.TypeBool, Default: false},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// HostTagsColumns holds the columns for the "host_tags" table.
	HostTagsColumns = []*schema.Column{
		{Name: "host_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// HostTagsTable holds the schema information for the "host_tags" table.
	HostTagsTable = &schema.Table{
		Name:       "host_tags",
		Columns:    HostTagsColumns,
		PrimaryKey: []*schema.Column{HostTagsColumns[0], HostTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "host_tags_host_id",
				Columns:    []*schema.Column{HostTagsColumns[0]},
				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "host_tags_tag_id",
				Columns:    []*schema.Column{HostTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TomeFilesColumns holds the columns for the "tome_files" table.
	TomeFilesColumns = []*schema.Column{
		{Name: "tome_id", Type: field.TypeInt},
		{Name: "file_id", Type: field.TypeInt},
	}
	// TomeFilesTable holds the schema information for the "tome_files" table.
	TomeFilesTable = &schema.Table{
		Name:       "tome_files",
		Columns:    TomeFilesColumns,
		PrimaryKey: []*schema.Column{TomeFilesColumns[0], TomeFilesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tome_files_tome_id",
				Columns:    []*schema.Column{TomeFilesColumns[0]},
				RefColumns: []*schema.Column{TomesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tome_files_file_id",
				Columns:    []*schema.Column{TomeFilesColumns[1]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BeaconsTable,
		FilesTable,
		HostsTable,
		HostCredentialsTable,
		HostFilesTable,
		HostProcessesTable,
		QuestsTable,
		RepositoriesTable,
		TagsTable,
		TasksTable,
		TomesTable,
		UsersTable,
		HostTagsTable,
		TomeFilesTable,
	}
)

func init() {
	BeaconsTable.ForeignKeys[0].RefTable = HostsTable
	HostCredentialsTable.ForeignKeys[0].RefTable = HostsTable
	HostCredentialsTable.ForeignKeys[1].RefTable = TasksTable
	HostFilesTable.ForeignKeys[0].RefTable = HostsTable
	HostFilesTable.ForeignKeys[1].RefTable = HostsTable
	HostFilesTable.ForeignKeys[2].RefTable = TasksTable
	HostProcessesTable.ForeignKeys[0].RefTable = HostsTable
	HostProcessesTable.ForeignKeys[1].RefTable = HostsTable
	HostProcessesTable.ForeignKeys[2].RefTable = TasksTable
	QuestsTable.ForeignKeys[0].RefTable = TomesTable
	QuestsTable.ForeignKeys[1].RefTable = FilesTable
	QuestsTable.ForeignKeys[2].RefTable = UsersTable
	RepositoriesTable.ForeignKeys[0].RefTable = UsersTable
	RepositoriesTable.Annotation = &entsql.Annotation{
		Table: "repositories",
	}
	TasksTable.ForeignKeys[0].RefTable = QuestsTable
	TasksTable.ForeignKeys[1].RefTable = BeaconsTable
	TomesTable.ForeignKeys[0].RefTable = UsersTable
	TomesTable.ForeignKeys[1].RefTable = RepositoriesTable
	HostTagsTable.ForeignKeys[0].RefTable = HostsTable
	HostTagsTable.ForeignKeys[1].RefTable = TagsTable
	TomeFilesTable.ForeignKeys[0].RefTable = TomesTable
	TomeFilesTable.ForeignKeys[1].RefTable = FilesTable
}
