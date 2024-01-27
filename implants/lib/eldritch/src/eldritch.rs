/// Tome for eldritch to execute.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Tome {
    #[prost(string, tag = "1")]
    pub eldritch: ::prost::alloc::string::String,
    #[prost(map = "string, string", tag = "2")]
    pub parameters: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::string::String,
    >,
    #[prost(string, repeated, tag = "3")]
    pub file_names: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
}
/// Process running on the host system.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Process {
    #[prost(uint64, tag = "1")]
    pub pid: u64,
    #[prost(uint64, tag = "2")]
    pub ppid: u64,
    #[prost(string, tag = "3")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub principal: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub path: ::prost::alloc::string::String,
    #[prost(string, tag = "6")]
    pub cmd: ::prost::alloc::string::String,
    #[prost(string, tag = "7")]
    pub env: ::prost::alloc::string::String,
    #[prost(string, tag = "8")]
    pub cwd: ::prost::alloc::string::String,
    #[prost(enumeration = "process::Status", tag = "9")]
    pub status: i32,
}
/// Nested message and enum types in `Process`.
pub mod process {
    #[derive(
        Clone,
        Copy,
        Debug,
        PartialEq,
        Eq,
        Hash,
        PartialOrd,
        Ord,
        ::prost::Enumeration
    )]
    #[repr(i32)]
    pub enum Status {
        Unspecified = 0,
        Unknown = 1,
        Idle = 2,
        Run = 3,
        Sleep = 4,
        Stop = 5,
        Zombie = 6,
        Tracing = 7,
        Dead = 8,
        WakeKill = 9,
        Waking = 10,
        Parked = 11,
        LockBlocked = 12,
        UninteruptibleDiskSleep = 13,
    }
    impl Status {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Status::Unspecified => "STATUS_UNSPECIFIED",
                Status::Unknown => "STATUS_UNKNOWN",
                Status::Idle => "STATUS_IDLE",
                Status::Run => "STATUS_RUN",
                Status::Sleep => "STATUS_SLEEP",
                Status::Stop => "STATUS_STOP",
                Status::Zombie => "STATUS_ZOMBIE",
                Status::Tracing => "STATUS_TRACING",
                Status::Dead => "STATUS_DEAD",
                Status::WakeKill => "STATUS_WAKE_KILL",
                Status::Waking => "STATUS_WAKING",
                Status::Parked => "STATUS_PARKED",
                Status::LockBlocked => "STATUS_LOCK_BLOCKED",
                Status::UninteruptibleDiskSleep => "STATUS_UNINTERUPTIBLE_DISK_SLEEP",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
            match value {
                "STATUS_UNSPECIFIED" => Some(Self::Unspecified),
                "STATUS_UNKNOWN" => Some(Self::Unknown),
                "STATUS_IDLE" => Some(Self::Idle),
                "STATUS_RUN" => Some(Self::Run),
                "STATUS_SLEEP" => Some(Self::Sleep),
                "STATUS_STOP" => Some(Self::Stop),
                "STATUS_ZOMBIE" => Some(Self::Zombie),
                "STATUS_TRACING" => Some(Self::Tracing),
                "STATUS_DEAD" => Some(Self::Dead),
                "STATUS_WAKE_KILL" => Some(Self::WakeKill),
                "STATUS_WAKING" => Some(Self::Waking),
                "STATUS_PARKED" => Some(Self::Parked),
                "STATUS_LOCK_BLOCKED" => Some(Self::LockBlocked),
                "STATUS_UNINTERUPTIBLE_DISK_SLEEP" => Some(Self::UninteruptibleDiskSleep),
                _ => None,
            }
        }
    }
}
/// ProcessList of running processes on the host system.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ProcessList {
    #[prost(message, repeated, tag = "1")]
    pub list: ::prost::alloc::vec::Vec<Process>,
}
/// File on the host system.
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct File {
    #[prost(string, tag = "1")]
    pub path: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub owner: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub group: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub permissions: ::prost::alloc::string::String,
    #[prost(int64, tag = "5")]
    pub size: i64,
    #[prost(string, tag = "6")]
    pub sha3_256_hash: ::prost::alloc::string::String,
    #[prost(bytes = "vec", tag = "7")]
    pub chunk: ::prost::alloc::vec::Vec<u8>,
}
